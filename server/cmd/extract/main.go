package main

import (
	"fmt"
	"log"

	"assert200.com/vincivia/internal/api/asx"
	"assert200.com/vincivia/internal/api/nasdaq"
	"assert200.com/vincivia/internal/api/yahoo"
	"assert200.com/vincivia/internal/date"
	"assert200.com/vincivia/internal/db"
)

func main() {

	recordedAt := date.CurrDay()

	log.Println("** Loading all US shares from Nasdaq API ** ")
	shares := nasdaq.GetShares()
	for _, share := range shares {
		_, err := db.AddRecord(share, recordedAt)
		if err != nil {
			log.Printf("ERROR: Adding record: %+v\nWith error: %v\n", share, err)
		} else {
			fmt.Print(".")
		}
	}
	log.Println("** Nasdaq Done ** ")

	log.Println("** Loading all shares from ASX API ** ")
	asxShares := asx.GetShares()

	for i := range asxShares {
		asxShares[i].Symbol = asxShares[i].Symbol + ".AX"
	}

	// Split the asxShares into batches of 20 items.
	batch := 20

	for i := 0; i < len(asxShares); i += batch {
		j := i + batch
		if j > len(asxShares) {
			j = len(asxShares)
		}

		asxSharesChunk := yahoo.GetQuotes(asxShares[i:j])

		for _, share := range asxSharesChunk {
			_, err := db.AddRecord(share, recordedAt)
			if err != nil {
				log.Printf("ERROR: Adding record: %+v\nWith error: %v\n", share, err)
			} else {
				fmt.Print(".")
			}
		}
	}
}
