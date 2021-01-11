package asx

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"assert200.com/vincivia/internal/entity"
)

const asxAllListedCompanies = "http://www.asx.com.au/asx/research/ASXListedCompanies.csv"

// GetShares from ASX
func GetShares() []entity.Share {
	response, err := http.Get(asxAllListedCompanies)
	if err != nil {
		log.Fatal("ERROR: couldn't open the source csv file from the asx.")
	}
	defer response.Body.Close()

	csvContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ERROR: couldn't read the source csv body")
	}

	//Need to remove top two lines to make it a correct CSV
	rows := strings.Split(string(csvContent), "\n")
	csvContentWithoutTop2LinesString := strings.Join(rows[3:], "\n")
	csvContentWithoutTop2LinesBuffer := bytes.NewBufferString(csvContentWithoutTop2LinesString)

	reader := csv.NewReader(csvContentWithoutTop2LinesBuffer)
	myData, err := reader.ReadAll()

	if err != nil {
		log.Fatalf("ERROR: Couldn't parse the csv data: %s", err)
	}

	var shares []entity.Share

	for _, element := range myData {

		var share entity.Share
		share.Name = element[0]
		share.Symbol = element[1]
		share.Industry = element[2]
		share.Exchange = entity.ASX

		shares = append(shares, share)
	}

	return shares
}
