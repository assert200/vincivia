package nasdaq

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// This gets everything (including NYSE Shares)
// https://api.nasdaq.com/api/screener/stocks?download=true
// https://api.nasdaq.com/api/screener/stocks?download=true&exchange=nyse

// GetShares from Nasdaq API
func GetShares() []Share {
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

	var shares []Share

	for _, row := range stocksResponse.Data.Rows {
		var share Share

		share.Symbol = strings.TrimSpace(row.Symbol)
		share.Name = strings.TrimSpace(row.Name)
		share.LastSale = quickParseFloat64(strings.TrimLeft(strings.TrimSpace(row.LastSale), "$"))
		share.NetChange = quickParseFloat64(strings.TrimSpace(row.NetChange))
		share.PctChange = quickParseFloat64(strings.TrimRight(strings.TrimSpace(row.PctChange), "%"))
		share.MarketCap = quickParseFloat64(strings.TrimSpace(row.MarketCap))
		share.Country = strings.TrimSpace(row.Country)
		share.IPOYear, _ = strconv.Atoi(strings.TrimSpace(row.IPOYear))
		share.Industry = strings.TrimSpace(row.Industry)
		share.Sector = strings.TrimSpace(row.Sector)
		shares = append(shares, share)
	}

	return shares
}

func quickParseFloat64(rawText string) float64 {
	cleanText := strings.ToLower(strings.TrimSpace(rawText))

	if cleanText == "" {
		return 0.0
	}

	if cleanText == "na" {
		return 0.0
	}

	value, err := strconv.ParseFloat(cleanText, 64)
	if err != nil {
		log.Println("WARNING: Was unable to quickParseFloat64: ", cleanText)
	}
	return value
}
