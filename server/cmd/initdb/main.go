package main

import (
	"fmt"
	"strings"

	"assert200.com/vincivia/internal/db"
	_ "github.com/lib/pq"
)

func main() {
	sqlContent := `
	
	set client_encoding = 'UTF8';
	
	DROP TABLE share;

		create table "share" (
				id serial primary key,
				symbol    text unique,
				name      text,
				last_sale  numeric,
				net_change numeric,
				pct_change numeric,
				market_cap numeric,
				country   text,
				ipo_year   numeric,
				industry  text,
				sector    text,
				recorded_at date
		);
	`
	requests := strings.Split(string(sqlContent), ";")

	for _, request := range requests {
		if len(request) == 0 {
			continue
		}

		fmt.Printf("\n\nEXECUTING SQL: %s", request)
		_, err := db.Get().Exec(request)
		checkErr(err)
		fmt.Printf("\nEXECUTION COMPLETE.\n")
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
