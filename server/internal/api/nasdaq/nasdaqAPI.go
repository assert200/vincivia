package nasdaq

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// This gets everything (including NYSTE Stockss)
// https://api.nasdaq.com/api/screener/stocks?download=true
// https://api.nasdaq.com/api/screener/stocks?download=true&exchange=nyse

// GetStocks from Nasdaq API
func GetStocks() []Security {
	var u url.URL
	u.Scheme = "https"
	u.Host = "api.nasdaq.com"
	u.Path = "api/screener/stocks"

	v := url.Values{}
	v.Add("download", "true")

	u.RawQuery = v.Encode()

	URLUnescaped, _ := url.PathUnescape(u.String())

	req, err := http.NewRequest("GET", URLUnescaped, nil)
	if err != nil {
		log.Fatal("Error building request: ", err)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

	client := http.DefaultClient
	r, err := client.Do(req)

	if err != nil {
		log.Fatal("ERROR: Sending request: ", err)
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Fatal("ERROR: Reading response: ", err)
	}

	var stocksResponse StocksResponse
	err = json.Unmarshal(body, &stocksResponse)
	if err != nil {
		log.Fatal("ERROR: Unmarshal response: ", err)
	}

	var securities []Security

	for _, row := range stocksResponse.Data.Rows {
		var security Security

		security.Symbol = row.Symbol
		security.Name = row.Name
		security.LastSale, _ = strconv.ParseFloat(row.LastSale, 64)
		security.NetChange, _ = strconv.ParseFloat(row.NetChange, 64)
		security.PctChange, _ = strconv.ParseFloat(row.PctChange, 64)
		security.MarketCap, _ = strconv.ParseFloat(row.MarketCap, 64)
		security.Country = row.Country
		security.IPOYear, _ = strconv.Atoi(row.IPOYear)
		security.Industry = row.Industry
		security.Sector = row.Sector
		security.URL = row.URL
		securities = append(securities, security)
	}

	return securities
}
