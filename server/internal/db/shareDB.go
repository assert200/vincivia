package db

import (
	"errors"
	"time"

	"assert200.com/vincivia/internal/api/nasdaq"
)

// AddShare to DB
func AddShare(d nasdaq.Share, recordedAt time.Time) (int, error) {
	var newShareID int
	sql := "insert into share (symbol, name, last_sale, net_change, pct_change, market_cap, country, ipo_year, industry, sector, recorded_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning id"
	err := Get().QueryRow(sql, d.Symbol, d.Name, d.LastSale, d.NetChange, d.PctChange, d.MarketCap, d.Country, d.IPOYear, d.Industry, d.Sector, recordedAt).Scan(&newShareID)
	if err != nil {
		return 0, errors.New("Error adding share to DB: " + err.Error())
	}
	return newShareID, nil
}
