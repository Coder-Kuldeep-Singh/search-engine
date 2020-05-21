package bing

import (
	"fmt"
	"log"
	"strings"

	"github.com/googlesearch/client"
	"github.com/googlesearch/color"
	"github.com/googlesearch/result"
)

var DuckDuckDomains = "https://www.bing.com/search?q="

func buildBingGoUrl(searchTerm string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	// SearchLimit := "100"
	return fmt.Sprintf("%s%s", DuckDuckDomains, searchTerm)
}

func BingScrape(searchTerm string) ([]result.Result, error) {
	DuckDuckUrl := buildBingGoUrl(searchTerm)
	fmt.Println(DuckDuckUrl)
	res, err := client.Request(DuckDuckUrl)
	if err != nil {
		return nil, err
	}
	scrapes, err := result.ResultParser(res, "li.b_algo", "h2 > a", "h2 > a", "div.b_caption > p")

	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func BingSearch(query string) {
	resp, err := BingScrape(query)
	if err != nil {
		log.Println("Having Error to visiting url: ", err)
	}
	for _, bing := range resp {
		fmt.Println()
		fmt.Println()
		// fmt.Println("[+]               ", bing.ResultRank)
		fmt.Println("             ", color.Info(bing.ResultTitle))
		fmt.Println("             ", bing.ResultURL)
		fmt.Println("             ---------------------------------------------------")
		fmt.Println("             ", bing.ResultDesc)
		fmt.Println()
		fmt.Println()
	}

}
