package yahoo

import (
	"fmt"
	"log"
	"strings"

	"github.com/googlesearch/client"
	"github.com/googlesearch/color"
	"github.com/googlesearch/result"
)

var YahooDomains = "https://in.search.yahoo.com/search;_ylt=AwrPhSwDZsZe9ksARk.7HAx.;_ylu=X3oDMTEzZzNnZms4BGNvbG8Dc2czBHBvcwMxBHZ0aWQDBHNlYwNwYWdpbmF0aW9u?p="

func buildYahooUrl(searchTerm string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	// SearchLimit := "100"
	return fmt.Sprintf("%s%s&pz=50&fr2=sb-top-in.search&b=1&pz=100&pstart=1", YahooDomains, searchTerm)
}

func YahooScrape(searchTerm string) ([]result.Result, error) {
	YahooUrl := buildYahooUrl(searchTerm)
	fmt.Println(YahooUrl)
	res, err := client.Request(YahooUrl)
	if err != nil {
		return nil, err
	}
	scrapes, err := result.ResultParser(res, "li > div.algo-sr", "div.compTitle > h3 > a", "div.compTitle > h3 > a", "div.compText > p")

	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func YahooSearch(query string) {
	resp, err := YahooScrape(query)
	if err != nil {
		log.Println("Having Error to visiting url: ", err)
	}
	for _, yahoo := range resp {
		fmt.Println()
		fmt.Println()
		fmt.Println("[+]               ", yahoo.ResultRank)
		fmt.Println("             ", color.Info(yahoo.ResultTitle))
		fmt.Println("             ", yahoo.ResultURL)
		fmt.Println("             ---------------------------------------------------")
		fmt.Println("             ", yahoo.ResultDesc)
		fmt.Println()
		fmt.Println()
	}

}
