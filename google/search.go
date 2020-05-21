package google

import (
	"fmt"
	"log"
	"strings"

	"github.com/googlesearch/client"
	"github.com/googlesearch/color"
	"github.com/googlesearch/result"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"uk":  "https://www.google.co.uk/search?q=",
	"ru":  "https://www.google.ru/search?q=",
	"fr":  "https://www.google.fr/search?q=",
}

func buildGoogleUrl(searchTerm string, countryCode string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		return fmt.Sprintf("%s%s&num=100&hl=%s", googleBase, searchTerm, "en")
	} else {
		return fmt.Sprintf("%s%s&num=100&hl=%s", googleDomains["com"], searchTerm, "en")
	}
}
func GoogleScrape(searchTerm string, countryCode string) ([]result.Result, error) {
	googleUrl := buildGoogleUrl(searchTerm, countryCode)
	res, err := client.Request(googleUrl)
	if err != nil {
		return nil, err
	}
	scrapes, err := result.ResultParser(res, "div.g", "a", "h3", "span.st")
	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func GoogleSearch(query, countryCode string) {
	resp, err := GoogleScrape(query, countryCode)
	if err != nil {
		log.Println("Having Error to visiting url: ", err)
	}
	for _, google := range resp {
		fmt.Println()
		fmt.Println()
		// fmt.Println("[+]               ", google.ResultRank)
		fmt.Println("            ", color.Info(google.ResultTitle))
		fmt.Println("            ", google.ResultURL)
		fmt.Println("            ---------------------------------------------------")
		fmt.Println("            ", google.ResultDesc)
		fmt.Println()
		fmt.Println()
	}
}
