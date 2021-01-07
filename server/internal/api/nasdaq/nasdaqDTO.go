package nasdaq

// StocksResponse Type
type StocksResponse struct {
	Data Data `json:"data"`
}

// Data Type
type Data struct {
	Rows []Row `json:"rows"`
}

// Row Type
type Row struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	LastSale  string `json:"lastsale"`
	NetChange string `json:"netchange"`
	PctChange string `json:"pctchange"`
	MarketCap string `json:"marketCap"`
	Country   string `json:"country"`
	IPOYear   string `json:"ipoyear"`
	Industry  string `json:"industry"`
	Sector    string `json:"sector"`
	URL       string `json:"url"`
}

// Security Type
type Security struct {
	Symbol    string
	Name      string
	LastSale  float64
	NetChange float64
	PctChange float64
	MarketCap float64
	Country   string
	IPOYear   int
	Industry  string
	Sector    string
	URL       string
}
