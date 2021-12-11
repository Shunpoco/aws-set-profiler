package parser

import (
	"bufio"
	"os"
	"strings"
)

func ParseConfig(path string) (items []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Trim(text, "[]")
		words := strings.Split(text, " ")
		if len(words) == 2 && words[0] == "profile" {
			items = append(items, words[1])
		}
	}

	return
}
