package topksequences

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name           string
		inputFilePaths []string
		wantFilePath   string
	}{
		{
			name:           "file_1",
			inputFilePaths: []string{"test_resources/files/file_1.txt"}, //TODO: [refactor] pass only file names
			wantFilePath:   "test_resources/results/file_1.txt",
		},
		{
			name:           "file_2",
			inputFilePaths: []string{"test_resources/files/file_2.txt"},
			wantFilePath:   "test_resources/results/file_2.txt",
		},
		{
			name:           "file_1 and file_2",
			inputFilePaths: []string{"test_resources/files/file_1.txt", "test_resources/files/file_2.txt"},
			wantFilePath:   "test_resources/results/file_1_2.txt",
		},
		{
			name:           "empty file",
			inputFilePaths: []string{"test_resources/files/empty_file.txt"},
			wantFilePath:   "test_resources/results/empty_file.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			oldOut := out
			defer func() {
				os.Args = oldArgs
				out = oldOut
			}()

			defaultMockArgs := []string{""}
			mockArgs := append(defaultMockArgs, tt.inputFilePaths...)
			os.Args = mockArgs

			var output string
			buf := bytes.NewBufferString(output)
			out = buf

			Execute()

			want, err := os.ReadFile(tt.wantFilePath)
			if err != nil {
				t.Errorf("Failed to read test resource result file")
				return
			}

			if buf.String() != string(want) {
				t.Errorf("Execute() got != want")

				fmt.Printf("got\n %s\n want\n%s", buf.String(), string(want))
			}
		})
	}
}
