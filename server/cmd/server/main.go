package main

import (
	"log"

	"assert200.com/vincivia/internal/api/nasdaq"
)

func main() {
	log.Println("** LETS DO IT ** ")
	// yahoo.GetQuotes([]string{"GOOG", "WORK", "BTC-USD"})
	stocks := nasdaq.GetStocks()

	for _, stock := range stocks {
		log.Println("%V", stock)
	}
}
