package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/googlesearch/bing"
	"github.com/googlesearch/color"
	"github.com/googlesearch/duckduckgo"
	"github.com/googlesearch/google"
)

func main() {
	query := flag.String("q", "", "Please provide the query you want to search")
	Engine := flag.String("E", "", "Provide the Search Engine")
	countryCode := flag.String("s", "com", "Please provide the country code")
	// Domain := flag.String("D", "", "Provide the direct domain")
	flag.Parse()
	if *query != "" {
		fmt.Println()
		fmt.Println("-----------------------------------Crawling Started-----------------------------------------")
		if *Engine == "google" {
			google.GoogleSearch(*query, *countryCode)
		} else if *Engine == "duckduckgo" {
			duckduckgo.DuckDuckSearch(*query)
		} else if *Engine == "bing" {
			bing.BingSearch(*query)
		} else {
			fmt.Println("Provide the Query")
			os.Exit(2)
		}
		fmt.Println("-----------------------------------Crawling Finished-----------------------------------------")
		fmt.Println()
		time.Sleep(time.Second * 10)
	} else {
		fmt.Println(color.Warn("Provide the Name of Search Engine you want's to use"))
		os.Exit(1)
	}

}
