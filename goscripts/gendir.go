package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var rawString string
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("failed reading from stdin %s", err.Error())
		os.Exit(1)
	}
	rawString = text

	dirName := generateName(rawString)
	initializeDir(dirName)
}

func initializeDir(s string) {

}

func generateName(s string) string {
	before, after, found := strings.Cut(s, ".")
	if !found {
		fmt.Errorf("separator '.' not found in the string")
		os.Exit(1)
	}
	numLen := len(before)
	zeroPad := strings.Repeat("0", 4-numLen)
	before = zeroPad + before

	after = strings.TrimSpace(after)
	after = strings.ReplaceAll(after, " ", "_")

	s = before + "__" + after

	fmt.Println(s)
	return s
}
