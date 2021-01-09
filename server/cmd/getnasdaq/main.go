package main

import (
	"fmt"
	"log"

	"assert200.com/vincivia/internal/api/nasdaq"
	"assert200.com/vincivia/internal/date"
	"assert200.com/vincivia/internal/db"
)

func main() {
	log.Println("** Loading Nasdaq ** ")

	// yahoo.GetQuotes([]string{"GOOG", "WORK", "BTC-USD"})
	shares := nasdaq.GetShares()

	recordedAt := date.CurrDay()
	for _, share := range shares {

		_, err := db.AddShare(share, recordedAt)
		if err != nil {
			log.Printf("ERROR: Adding share: %+v\nWith error: %v\n", share, err)
		} else {
			fmt.Print(".")
		}
	}

	log.Println("** Done ** ")
}
