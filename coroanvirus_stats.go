package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
    lookupUrl string = "https://www.worldometers.info/coronavirus/"
)

var (
    corona_data = []CoronaVirusStat{}
    overall_data = []OverallStat{}
)

type CoronaVirusStat struct {
    Country string
    TotalCases int
    NewCases int
    TotalDeaths int
    NewDeaths int
    TotalRecovered int
    ActiveCases int
    Serious int
    TotalCasesOneMil int
}

type OverallStat struct {
    Title string
    Data int
}

func convertStrInt(convert string) int {
    convertedString := strings.TrimSpace(convert)
    data_string := strings.ReplaceAll(convertedString, ",", "")
    data_string = strings.ReplaceAll(data_string, "+", "")
    number, err := strconv.Atoi(data_string)
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
            Data: convertStrInt(sel.Find("span").Text()),
        }
        overall_data = append(overall_data, data)
    })
}

func GetCountriesData(doc *goquery.Document) {

    table := doc.Find("#main_table_countries_today")
    tbody := table.Find("tbody")
    rows := tbody.Find("tr")

    for i := range rows.Nodes {
        country_row := rows.Eq(i)

        columns := country_row.Find("td")

        data := CoronaVirusStat{
            Country: columns.Eq(0).Text(),
            TotalCases: convertStrInt(columns.Eq(1).Text()),
            NewCases: convertStrInt(columns.Eq(2).Text()),
            TotalDeaths: convertStrInt(columns.Eq(3).Text()),
            NewDeaths: convertStrInt(columns.Eq(4).Text()),
            TotalRecovered: convertStrInt(columns.Eq(5).Text()),
            ActiveCases: convertStrInt(columns.Eq(6).Text()),
            Serious: convertStrInt(columns.Eq(7).Text()),
            TotalCasesOneMil: convertStrInt(columns.Eq(8).Text()),
        }
        corona_data = append(corona_data, data)
    }
}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
func GetUrlInfo(url string) (*goquery.Document, error) {

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

func PrintOverallData() {
    fmt.Println("\nCoronavirus data:")
    for _, overall_item := range overall_data {
        fmt.Printf("%s - %d\n", overall_item.Title, overall_item.Data)
    }
}

func PrintCoronaData() {
    fmt.Println("\nCountry\t| Total Cases\t| NewCases\t| TotalDeaths\t| NewDeaths\t| TotalRecovered\t| ActiveCases\t| Serious\t| TotalCasesOneMil\t|\n")
    for _, item := range corona_data {
        fmt.Printf("%s\t| %d\t| %d\t| %d\t| %d\t| %d\t| %d\t| %d\t| %d\t|\n",
            item.Country,
            item.TotalCases,
            item.NewCases,
            item.TotalDeaths,
            item.NewDeaths,
            item.TotalRecovered,
            item.ActiveCases,
            item.Serious,
            item.TotalCasesOneMil,
        )
    }
}

func main() {
	docs, err := GetUrlInfo(lookupUrl)
	if err != nil {
		log.Println(err)
	}

	GetOverallData(docs)
	GetCountriesData(docs)

    PrintOverallData()
    PrintCoronaData()
}
