package yahoo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"assert200.com/vincivia/internal/entity"
)

// url="https://query1.finance.yahoo.com/v7/finance/quote?symbols=GOOG,WORK"

// GetQuotes from Yahoo API
func GetQuotes(shares []entity.Share) []entity.Share {
	var u url.URL
	u.Scheme = "https"
	u.Host = "query1.finance.yahoo.com"
	u.Path = "v7/finance/quote"

	symbolsString := ""
	for i, share := range shares {
		symbolsString += share.Symbol
		if i < len(shares)-1 {
			symbolsString += ","
		}
	}

	v := url.Values{}
	v.Add("symbols", symbolsString)

	u.RawQuery = v.Encode()

	URLUnescaped, _ := url.PathUnescape(u.String())

	req, err := http.NewRequest("GET", URLUnescaped, nil)
	if err != nil {
		log.Fatal("Error building request: ", err)
	}

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

	var quoteResponse QuoteResponse
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		log.Fatal("ERROR: Unmarshal response: ", err)
	}

	for i, share := range shares {
		shareResponse, err := FindResultBySymbol(quoteResponse.Response.Results, share.Symbol)
		if err != nil {
			log.Printf("WARNING: %v not found in yahoo results", share.Symbol)
			continue
		}

		shares[i].LastSale = shareResponse.RegularMarketPrice
		shares[i].MarketCap = shareResponse.MarketCap
		shares[i].NetChange = shareResponse.RegularMarketChange
		shares[i].PctChange = shareResponse.RegularMarketChangePercent
	}

	return shares
}
