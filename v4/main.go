package main

import (
	"go-mircoservice-learn/v4/cmd"
	"log"
)

func main() {

	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
