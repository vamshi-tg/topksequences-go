# 100 most 3 Word Sequences in a text

A go program that outputs a list of the 100 most common three word sequences in the text, along with a count of how many 
times each occurred in the text.

## Running the application

go v1.18 or greater must be installed. To check run `go version`

```shell
$ go version
go version go1.18.3 linux/amd64
```

## Running the program with `go run`
* Clone the repository and cd into the repository.
* Pass files as arguments to the below command.

```shell
go run main.go
````

```shell
# Example
go run main.go ./topksequences/test_resources/files/file_2.txt

# Output
320 - of the same
126 - the same species
125 - conditions of life
116 - in the same
108 - of natural selection
103 - from each other
98 - species of the
89 - on the other
81 - the other hand
78 - the case of
76 - the theory of
74 - of the world
74 - parts of the
73 - some of the
70 - through natural selection
69 - with respect to
67 - in the case
65 - it may be
65 - the inhabitants of
65 - the species of
64 - of the species
62 - that of the
61 - forms of life
61 - the same genus
60 - individuals of the
58 - as far as
56 - the number of
55 - part of the
55 - those of the
53 - the principle of
52 - the nature of
52 - to each other
51 - in this case
51 - on the same
50 - at the same
50 - more or less
50 - nature of the
50 - to the same
49 - as in the
49 - in regard to
47 - a state of
47 - and in the
47 - in which the
47 - nearly the same
47 - one of the
47 - the individuals of
46 - each other in
46 - inhabitants of the
46 - it has been
46 - state of nature
46 - the amount of
45 - are descended from
45 - from a common
45 - the united states
45 - we can understand
44 - might have been
44 - will have been
43 - by natural selection
43 - the conditions of
42 - and of the
42 - and on the
42 - animals and plants
42 - in a state
42 - the same manner
42 - to believe that
42 - to have been
42 - which have been
41 - respect to the
41 - the same time
41 - we have seen
41 - would have been
40 - as well as
40 - in some degree
40 - it would be
40 - members of the
40 - on the theory
40 - there is no
40 - varieties of the
39 - belonging to the
39 - each other and
39 - of the most
38 - could be given
38 - it is that
38 - that it is
38 - theory of natural
37 - in order to
36 - it is not
36 - of life and
36 - species belonging to
36 - the present day
36 - the process of
36 - the sterility of
35 - reason to believe
35 - the power of
34 - in relation to
33 - and this is
33 - at the present
33 - believe that the
33 - for instance the
33 - from the same
```

We can pass multiple files too:
```shell
go run main.go file1.txt file2.txt file3.txt
```

Accept input from stdin:
```shell
cat file1.txt | go run main.go
```

## Alternatively, the above commands can be run with the pre-built package
* CD into repo
* Run the below command

```shell
# Passing files as argument
./go-topksequences ./topksequences/test_resources/files/file_2.txt

or 

# Passing the input from Stdin
cat ./topksequences/test_resources/files/file_2.txt | ./go-topksequences
```

## Running Tests
```shell
go test -v ./topksequences
```