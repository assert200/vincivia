package db

import (
	"errors"
	"time"

	"assert200.com/vincivia/internal/entity"
)

// AddRecord to DB
func AddRecord(e entity.Share, recordedAt time.Time) (int, error) {
	var newRecordID int
	sql := `SELECT * FROM add_record($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`

	err := Get().QueryRow(sql, e.Symbol, e.Name, e.Exchange, e.Country, e.IPOYear, e.Industry, e.Sector, e.LastSale, e.NetChange, e.PctChange, e.MarketCap, recordedAt).Scan(&newRecordID)
	if err != nil {
		return 0, errors.New("Error adding share to DB: " + err.Error())
	}
	return newRecordID, nil
}
