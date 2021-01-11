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
}
