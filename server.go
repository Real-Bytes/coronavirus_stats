package main

import (
	"fmt"
	"log"

	cs "github.com/ryanjb1/coronavirus_stats/coronavirusstats"
)

func PrintOverallData() {
	fmt.Println("\nCoronavirus data:")
	for _, overall_item := range cs.overallData {
		fmt.Printf("%s - %d\n", overall_item.Title, overall_item.Data)
	}
}

func PrintCoronaData() {
	fmt.Println("\nCountry\t| Total Cases\t| NewCases\t| TotalDeaths\t| NewDeaths\t| TotalRecovered\t| ActiveCases\t| Serious\t| TotalCasesOneMil\t|")
	for _, item := range cs.coronaData {
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
	docs, err := cs.GetURLInfo(cs.lookupURL)
	if err != nil {
		log.Println(err)
	}

	cs.GetOverallData(docs)
	cs.GetCountriesData(docs)

	PrintOverallData()
	PrintCoronaData()
}
