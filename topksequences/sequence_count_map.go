package topksequences

import (
	"github.com/gammazero/deque"
	"strings"
)

// sequenceCountMap stores the counts of sequences in a map
type sequenceCountMap map[string]int

func newSequenceCountMap() sequenceCountMap {
	return map[string]int{}
}

// incCount will increment the count of the given sequence
func (s sequenceCountMap) incCount(sequence string) {
	if count, present := s[sequence]; present {
		s[sequence] = count + 1
	} else {
		s[sequence] = 1
	}
}

// sequenceCountMapKey builds key of the map
// Key is built using the values present in the deque window
func sequenceCountMapKey(deque *deque.Deque[string]) string {
	var sb strings.Builder

	for i := 0; i < deque.Len(); i++ {
		sb.WriteString(deque.At(i))

		if i < deque.Len()-1 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

// mergeMaps merges the given maps
func mergeMaps(map1, map2 sequenceCountMap) sequenceCountMap {
	for key, count1 := range map1 {
		if count2, present := map2[key]; present {
			map1[key] = count2 + count1
			delete(map2, key)
		}
	}

	for key, count := range map2 {
		map1[key] = count
	}

	return map1
}
