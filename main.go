package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type ParsedDiff struct {
	CreatedRules []string
	DeletedRules []string
	UpdatedRules []string
}

func main() {
	testReader, err := os.Open("diff5.txt")
	if err != nil {
		os.Exit(1)
	}
	defer testReader.Close()
	text, err := io.ReadAll(testReader)
	resp, err := ParseFileDiff(string(text))
	fmt.Printf("%v", resp)

}

func ParseFileDiff(diff string) (ParsedDiff, error) {
	parsedDiff := ParsedDiff{}
	hunks := strings.Split(diff, "diff --git")
	//remove first empty element
	hunks = hunks[1:]
	for _, hunk := range hunks {
		//create file
		if strings.Contains(hunk, "--- /dev/null") {
			parsedDiff.CreatedRules = append(parsedDiff.CreatedRules, hunk)
		} else if strings.Contains(hunk, "+++ /dev/null") { //delete file
			parsedDiff.DeletedRules = append(parsedDiff.DeletedRules, hunk)
		} else {
			parsedDiff.UpdatedRules = append(parsedDiff.UpdatedRules, hunk)
		}
	}
	return parsedDiff, nil
}
