package main

import (
	"aws-set-profiler/src/parser"
	"aws-set-profiler/src/prompt"
	"aws-set-profiler/src/writer"
	"flag"
	"fmt"
	"os"
)

var configDir string

func init() {
	homeDir, _ := os.UserHomeDir()
	flag.StringVar(&configDir, "directory", fmt.Sprintf("%s/.aws", homeDir), "The directory of aws config file.")
}

func main() {
	flag.Parse()

	configFile := fmt.Sprintf("%s/config", configDir)
	profileFile := fmt.Sprintf("%s/profile", configDir)

	items, err := parser.ParseConfig(configFile)
	if err != nil {
		fmt.Printf("Parser failed: %v\n", err)
		return
	}
	profile, err := prompt.GetProfile(items)
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	writer.WriteProfile(profileFile, profile)
}
