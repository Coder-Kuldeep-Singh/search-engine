package duckduckgo

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DuckDuckResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var DuckDuckDomains = "https://duckduckgo.com/html/?q="

func buildDuckDuckGoUrl(searchTerm string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	// SearchLimit := "100"
	return fmt.Sprintf("%s%s", DuckDuckDomains, searchTerm)
}

func DuckDuckRequest(searchURL string) (*http.Response, error) {

	baseClient := &http.Client{}

	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

	res, err := baseClient.Do(req)

	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func DuckDuckResultParser(response *http.Response) ([]DuckDuckResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	results := []DuckDuckResult{}
	sel := doc.Find("div.links_main")
	// fmt.Println(sel)
	rank := 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("h2 > a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h2 > a")
		descTag := item.Find("a")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" {
			result := DuckDuckResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank += 1
		}
	}
	return results, err
}

func DuckDuckScrape(searchTerm string) ([]DuckDuckResult, error) {
	DuckDuckUrl := buildDuckDuckGoUrl(searchTerm)
	fmt.Println(DuckDuckUrl)
	res, err := DuckDuckRequest(DuckDuckUrl)
	if err != nil {
		return nil, err
	}
	scrapes, err := DuckDuckResultParser(res)

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
