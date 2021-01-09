package db

import (
	"errors"
	"time"

	"assert200.com/vincivia/internal/api/nasdaq"
)

// AddShare to DB
func AddShare(d nasdaq.Share, recordedAt time.Time) (int, error) {
	var newShareID int
	sql := `SELECT * FROM add_record($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

	err := Get().QueryRow(sql, d.Symbol, d.Name, d.Country, d.IPOYear, d.Industry, d.Sector, d.LastSale, d.NetChange, d.PctChange, d.MarketCap, recordedAt).Scan(&newShareID)
	if err != nil {
		return 0, errors.New("Error adding share to DB: " + err.Error())
	}
	return newShareID, nil
}
