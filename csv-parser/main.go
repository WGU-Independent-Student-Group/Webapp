package main

import (
	"fmt"
	"group-gis/internal/data"
	"log"
)

func main() {
	records, err := data.Load("internal/data/mrds.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %d records\n", len(records))
}
