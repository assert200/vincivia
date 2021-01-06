package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"assert200.com/vincivia/internal/dto"
)

// url="https://query1.finance.yahoo.com/v7/finance/quote?symbols=GOOG,WORK"

// Do It
func Do() {
	var u url.URL
	u.Scheme = "https"
	u.Host = "query1.finance.yahoo.com"
	u.Path = "v7/finance/quote"

	v := url.Values{}
	v.Add("symbols", "GOOG,WORK")

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

	var quoteResponse dto.QuoteResponse
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("%V", quoteResponse)
}
