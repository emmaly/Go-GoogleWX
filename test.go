package main

import (
	"googlewx"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Requires QUERY argument.\n")
		os.Exit(1)
	}

	query := os.Args[1]

	wx, err := googlewx.Get(query)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", wx)
}
