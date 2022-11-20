package topksequences

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/gammazero/deque"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	k = 100
	n = 3 // words sequence size
)

var (
	maxWorkers                         = 10        // Max workers that process individual files
	out               io.Writer        = os.Stdout // Holds the redirect location of the results. Default is StdOut
	gSequenceCountMap sequenceCountMap             // global map to store the frequency of 'n' word sequences from all the given files/text.
)

/*
	Execute will return a list of the 'k' most common 'n' word sequences from the given files/text
	By default, returns 100 most common 3 word sequence in the given text
*/
func Execute() {
	gSequenceCountMap = newSequenceCountMap()

	mergingMapsDone := make(chan bool)
	sequenceCountMapStream := make(chan sequenceCountMap)

	go mergeSequenceCountMaps(sequenceCountMapStream, mergingMapsDone)

	if isInputFromPipe() {
		processText(os.Stdin, sequenceCountMapStream)
	} else {
		processFilesFromArgs(sequenceCountMapStream)
	}
	close(sequenceCountMapStream)

	<-mergingMapsDone
	topKSequences := getTopKSequences(gSequenceCountMap, k)
	printSequenceCounts(topKSequences)
}

/*
	processFileFromArgs process all the files that are given the args.

	It maintains a worker pool of size 'numWorkers'. Each file is processed concurrently
 	by a worker from the worker pool.
*/
func processFilesFromArgs(sequenceCountMapStream chan<- sequenceCountMap) {
	validateArgs()

	filePaths := getFilePathsFromArgs()

	// based on the num of files dynamically set num of workers
	numWorkers := maxWorkers
	if len(filePaths) < numWorkers {
		numWorkers = len(filePaths)
	}

	filePathStream := make(chan string, numWorkers)

	var wg sync.WaitGroup
	// start worker routines for processing individual files from the filePathStream
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(filePathStream, sequenceCountMapStream)
		}()
	}

	// push files to process to the filePathStream
	for _, filePath := range filePaths {
		filePathStream <- filePath
	}
	close(filePathStream)

	wg.Wait()
}

func validateArgs() {
	filePaths := getFilePathsFromArgs()

	if len(filePaths) < 1 {
		log.Fatal("Missing parameter, provide a file name")
	}

	// check if all the files exists
	allFilesExists := true
	for _, filePath := range filePaths {
		exists, err := fileExists(filePath)

		if err != nil {
			fmt.Println("Unknown error occurred while during the validation:" + err.Error())
			return
		}
		if !exists {
			fmt.Println("File not found at path:", filePath)
			allFilesExists = false
		}
	}

	if !allFilesExists {
		log.Fatal("Above files not found. Cannot proceed !!")
	}
}

func worker(files <-chan string, sequenceCountMapStream chan<- sequenceCountMap) {
	for file := range files {
		processFile(file, sequenceCountMapStream)
	}
}

func processFile(filePath string, sequenceCountMapStream chan<- sequenceCountMap) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file", err.Error())
	}
	defer file.Close()

	processText(file, sequenceCountMapStream)
}

/*
	processText scans the text, builds the fSequenceCountMap and pushes it to a channel.

	Params:
	r - source of the text to be processed
	sequenceCountMapStream - channel to which the resultant fSequenceCountMap is pushed

	Result:
	fSequenceCountMap - Holds the count of all the 'n' word sequences present int the text

	It uses sliding window technique to scan the buffered text. It maintains a double
	ended queue for the sliding window, which will store a 'n' word sequence while traversing
    the text.
*/
func processText(r io.Reader, sequenceCountMapStream chan<- sequenceCountMap) {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	scanner.Split(bufio.ScanWords)

	var windowDeque deque.Deque[string]

	// populate the window
	for i := 0; i < n && scanner.Scan(); {
		word := sanitizeText(scanner.Text())
		if word != "" {
			windowDeque.PushBack(word)
			i++
		}
	}

	// sequence count map for the current text
	fSequenceCountMap := newSequenceCountMap()

	// abort when enough text is not present
	if windowDeque.Len() < n {
		sequenceCountMapStream <- fSequenceCountMap
		return
	}

	// keep sliding the window until the text is completed
	for scanner.Scan() {
		nextWord := sanitizeText(scanner.Text())
		if nextWord == "" {
			continue
		}

		key := sequenceCountMapKey(&windowDeque)
		fSequenceCountMap.incCount(key)

		// update the sliding window with the new word
		windowDeque.PopFront()
		windowDeque.PushBack(nextWord)
	}

	if windowDeque.Len() == n {
		// add any remaining items from the window
		key := sequenceCountMapKey(&windowDeque)
		fSequenceCountMap.incCount(key)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failure while scanning text.\n Error: %s", err.Error())
		return
	}

	sequenceCountMapStream <- fSequenceCountMap
}

func getFilePathsFromArgs() []string {
	return os.Args[1:]
}

/*
	getTopKSequences returns the top k n-word sequences from the global gSequenceCountMap.

	It uses a min-heap to return the k n-word sequences. It efficiently does this in O(n log(k))
	time.
*/
func getTopKSequences(sequenceCountMap sequenceCountMap, k int) []*sequenceCount {
	topKSequences := newSequenceCountMinHeap()

	for sequence, count := range sequenceCountMap {
		if strings.TrimSpace(sequence) == "" {
			continue
		}

		heap.Push(topKSequences, &sequenceCount{sequence, count})

		if topKSequences.Len() > k {
			heap.Pop(topKSequences)
		}
	}

	topKSequencesLen := topKSequences.Len()
	result := make([]*sequenceCount, topKSequencesLen)

	i := topKSequencesLen - 1
	for !topKSequences.Empty() {
		sc := topKSequences.Top()
		heap.Pop(topKSequences)
		result[i] = sc
		i--
	}

	return result
}

/*
	mergeSequenceCountMaps receives the sequence counts maps of different files from a channel and
    merges it into the global gSequenceCountMap.

	Once the merging is completed it sends a signal on the 'done' channel.
*/
func mergeSequenceCountMaps(sequenceCountMapStream <-chan sequenceCountMap, done chan<- bool) {
	for val := range sequenceCountMapStream {
		gSequenceCountMap = mergeMaps(gSequenceCountMap, val)
	}
	done <- true
}

func printSequenceCounts(sequenceCounts []*sequenceCount) {
	for _, sc := range sequenceCounts {
		fmt.Fprintf(out, "%d - %s\n", sc.Value, sc.Key)
	}
}
