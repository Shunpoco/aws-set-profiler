package main

import (
	"aws-set-profiler/src/parser"
	"aws-set-profiler/src/prompt"
	"aws-set-profiler/src/writer"
	"fmt"
)

func main() {
	items, err := parser.ParseConfig("./test_config")
	if err != nil {
		fmt.Printf("Parser faile: %v\n", err)
		return
	}
	profile, err := prompt.GetProfile(items)
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	writer.WriteProfile("config", profile)
}
