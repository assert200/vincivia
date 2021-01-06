package dto

// QuoteResponse form DTO
type QuoteResponse struct {
	Results Results `json:"quoteResponse"`
}

// Results form DTO
type Results struct {
	Results []Result `json:"result"`
}

// Result form DTO
type Result struct {
	Symbol              string  `json:"symbol"`
	ShortName           string  `json:"shortName"`
	RegularMarketPrice  float64 `json:"regularMarketPrice"`
	RegularMarketVolume int     `json:"regularMarketVolume"`
	// RegularMarketTime   time.Time `json:"regularMarketTime"`
}
