package main

import (
	"fmt"
	"os"
)

func main() {
	csv := "mrds.csv"

	if _, err := os.Stat(csv); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("mrds file not found")
			os.Exit(1)
		}

		fmt.Printf("error accessing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("csv exists!")
}
