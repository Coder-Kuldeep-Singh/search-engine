package duckduckgo

import (
	"fmt"
	"log"
	"strings"

	"github.com/googlesearch/client"
	"github.com/googlesearch/color"
	"github.com/googlesearch/result"
)

var DuckDuckDomains = "https://duckduckgo.com/html/?q="

func buildDuckDuckGoUrl(searchTerm string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	// SearchLimit := "100"
	return fmt.Sprintf("%s%s", DuckDuckDomains, searchTerm)
}

func DuckDuckScrape(searchTerm string) ([]result.Result, error) {
	DuckDuckUrl := buildDuckDuckGoUrl(searchTerm)
	fmt.Println(DuckDuckUrl)
	res, err := client.Request(DuckDuckUrl)
	if err != nil {
		return nil, err
	}
	scrapes, err := result.ResultParser(res, "div.links_main", "h2 > a", "h2 > a", "a.result__snippet")

	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func DuckDuckSearch(query string) {
	resp, err := DuckDuckScrape(query)
	if err != nil {
		log.Println("Having Error to visiting url: ", err)
	}
	for _, duck := range resp {
		fmt.Println()
		fmt.Println()
		// fmt.Println("[+]               ", duck.ResultRank)
		fmt.Println("             ", color.Info(duck.ResultTitle))
		fmt.Println("             ", duck.ResultURL)
		fmt.Println("             ---------------------------------------------------")
		fmt.Println("             ", duck.ResultDesc)
		fmt.Println()
		fmt.Println()
	}

}
