package topksequences

import (
	"os"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func sanitizeText(str string) string {
	return strings.TrimSpace(
		strings.TrimSpace(strings.ToLower(nonAlphanumericRegex.ReplaceAllString(str, ""))),
	)
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
