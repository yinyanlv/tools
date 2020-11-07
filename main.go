package main

import (
	"log"
	"tools/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute() err: %v", err)
	}
	// timer.TestFormatTime()
}
