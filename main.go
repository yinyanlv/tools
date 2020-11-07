package main

import (
	"fmt"
	"log"
	"tools/cmd"
	"tools/internal/time"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute() err: %v", err)
	}

	fmt.Println(time.GetNowTime())
}
