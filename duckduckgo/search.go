package duckduckgo

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/googlesearch/client"
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
	fmt.Println("Did you want to Run the Software")
	fmt.Print("Yes or No: ") //Print function is used to display output in same line
	var first string
	fmt.Scanln(&first)
	if first == "yes" || first == "Yes" || first == "YES" {
		fmt.Println("Crawling Started")
		resp, err := DuckDuckScrape(query)
		if err != nil {
			log.Println("Having Error to visiting url: ", err)
		}
		for _, duck := range resp {
			log.Println("[+]", duck.ResultTitle)
			log.Println("[+]", duck.ResultURL)
			log.Println("[+]", duck.ResultDesc)
			log.Println()
			log.Println("[+] ===============================")
			log.Println()
		}
		fmt.Println("Crawling Finished...")
	} else {
		os.Exit(1)
	}
}
