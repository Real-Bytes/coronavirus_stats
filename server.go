package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	cs "github.com/ryanjb1/coronavirus_stats/coronavirusstats"
)

const (
	LookupURL string = "https://www.worldometers.info/coronavirus/"
)

func PrintOverallData() {
	fmt.Println("\nCoronavirus data:")
	for _, overallItem := range cs.OverallData {
		fmt.Printf("%s - %d\n", overallItem.Title, overallItem.Data)
	}
}

func PrintCoronaData() {
	fmt.Println("\nCountry\t| Total Cases\t| NewCases\t| TotalDeaths\t| NewDeaths\t| TotalRecovered\t| ActiveCases\t| Serious\t| TotalCasesOneMil\t|")
	for _, item := range cs.CoronaData {
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

type CoronaData struct {
	PageTitle string
	Data []cs.CoronaVirusStat
}

func startServer() {
	r := mux.NewRouter()

	tmpl := template.Must(template.ParseFiles("index.html"))

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := CoronaData{
			PageTitle: "Corona Stats",
			Data: cs.CoronaData,
		}
		tmpl.Execute(writer, data)
	})

	http.ListenAndServe(":8888", r)
}

func main() {
	docs, err := cs.GetURLInfo(LookupURL)
	if err != nil {
		log.Println(err)
	}

	cs.GetOverallData(docs)
	cs.GetCountriesData(docs)

	fmt.Printf("Corona stats live on http://localhost:8888")

	startServer()

}
