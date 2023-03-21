package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func cleanString(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	return input
}

func ReadAndCleanString() string {
	result, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	result = cleanString(result)
	return result
}
