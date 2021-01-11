package entity

// Exchange type
type Exchange string

const (
	// US Exchange include many including Nasdaq and NYSE
	US Exchange = "us"
	// ASX Exchange
	ASX Exchange = "asx"
)

// Share Entity
type Share struct {
	Symbol    string
	Name      string
	Exchange  Exchange
	LastSale  float64
	NetChange float64
	PctChange float64
	MarketCap float64
	Country   string
	IPOYear   int
	Industry  string
	Sector    string
}
