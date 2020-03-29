package coronavirusstats

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	CoronaData  []CoronaVirusStat
	OverallData []OverallStat
)

type CoronaVirusStat struct {
	Country          string
	TotalCases       int
	NewCases         int
	TotalDeaths      int
	NewDeaths        int
	TotalRecovered   int
	ActiveCases      int
	Serious          int
	TotalCasesOneMil int
}

type OverallStat struct {
	Title string
	Data  int
}

func convertStrInt(convert string) int {
	convertedString := strings.TrimSpace(convert)
	dataString := strings.ReplaceAll(convertedString, ",", "")
	dataString = strings.ReplaceAll(dataString, "+", "")
	number, err := strconv.Atoi(dataString)
	if err != nil {
		// check error
	}
	return number
}

func GetOverallData(doc *goquery.Document) {
	// Save each of coronavirus main info

	// Get Overall Information
	doc.Find("#maincounter-wrap").Each(func(index int, sel *goquery.Selection) {
		data := OverallStat{
			Title: sel.Find("h1").Text(),
			Data:  convertStrInt(sel.Find("span").Text()),
		}
		OverallData = append(OverallData, data)
	})
}

func GetCountriesData(doc *goquery.Document) {

	table := doc.Find("#main_table_countries_today")
	tbody := table.Find("tbody")
	rows := tbody.Find("tr")

	for i := range rows.Nodes {
		countryRow := rows.Eq(i)

		columns := countryRow.Find("td")

		data := CoronaVirusStat{
			Country:          columns.Eq(0).Text(),
			TotalCases:       convertStrInt(columns.Eq(1).Text()),
			NewCases:         convertStrInt(columns.Eq(2).Text()),
			TotalDeaths:      convertStrInt(columns.Eq(3).Text()),
			NewDeaths:        convertStrInt(columns.Eq(4).Text()),
			TotalRecovered:   convertStrInt(columns.Eq(5).Text()),
			ActiveCases:      convertStrInt(columns.Eq(6).Text()),
			Serious:          convertStrInt(columns.Eq(7).Text()),
			TotalCasesOneMil: convertStrInt(columns.Eq(8).Text()),
		}
		CoronaData = append(CoronaData, data)
	}
}

// GetURLInfo gets the latest blog title headings from the url
// given and returns them as a list.
func GetURLInfo(url string) (*goquery.Document, error) {

	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc, nil
}
