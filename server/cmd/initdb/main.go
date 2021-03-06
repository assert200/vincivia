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

	create table share (
		id serial primary key,
		symbol    text,
		name      text,
		exchange  text,
		country   text,
		ipo_year  numeric,
		industry  text,
		sector    text,
		unique (symbol, exchange)

		);

	create table record (
			id serial primary key,
			share_id int not null references share(id),
			last_sale  numeric,
			net_change numeric,
			pct_change numeric,
			market_cap numeric,
			recorded_at date not null,
			unique (share_id, recorded_at)
	);		
`

	/*
	   -- Run this command manually to create the stored procedure:

	   CREATE OR REPLACE FUNCTION add_record(p_symbol text, p_name text, p_exchange text, p_country text, p_ipo_year numeric, p_industry text, p_sector text, p_last_sale numeric, p_net_change numeric, p_pct_change numeric, p_market_cap numeric, p_recorded_at date) RETURNS integer AS $$
	   		declare
	   			share_id share.id%type;
	   			record_id record.id%type;
	   		begin
	   			SELECT id
	   			FROM share
	   			INTO share_id
	   			WHERE symbol=p_symbol and exchange=p_exchange;
	   			IF NOT FOUND THEN
	   				RAISE NOTICE 'The share with symbol % could not be found', p_symbol;
	   			insert into share (symbol, name, exchange, country, ipo_year, industry, sector) VALUES (p_symbol, p_name, p_exchange, p_country, p_ipo_year, p_industry, p_sector) RETURNING ID INTO share_id;
	   			ELSE
	   				RAISE NOTICE 'The share id is %', share_id;
	   			END IF;

	   			INSERT INTO record (share_id, last_sale, net_change, pct_change, market_cap, recorded_at) VALUES (share_id, p_last_sale, p_net_change, p_pct_change, p_market_cap, p_recorded_at) returning id into record_id;
	   			RETURN record_id;
	   		end
	   	$$ LANGUAGE plpgsql;

	*/

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
