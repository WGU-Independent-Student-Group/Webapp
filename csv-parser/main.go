package main

import (
	"fmt"
	"group-gis/internal/data"
	"log"
)

func main() {
	metadata, siteIdent, siteChar, geoLoc, err := data.Load("internal/data/mrds.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %d records\n", len(metadata))

	if len(metadata) > 0 {
		fmt.Printf("First Metadata entry: %+v\n", metadata[0])
	}
	if len(siteIdent) > 0 {
		fmt.Printf("First SiteIdentification entry: %+v\n", siteIdent[0])
	}
	if len(siteChar) > 0 {
		fmt.Printf("First SiteCharacteristics entry: %+v\n", siteChar[0])
	}
	if len(geoLoc) > 0 {
		fmt.Printf("First GeographicLocation entry: %+v\n", geoLoc[0])
	}
}
