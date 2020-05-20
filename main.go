package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/googlesearch/duckduckgo"
	"github.com/googlesearch/google"
)

func main() {
	query := flag.String("q", "", "Please provide the query you want to search")
	Engine := flag.String("E", "", "Provide the Search Engine")
	countryCode := flag.String("s", "com", "Please provide the country code")
	// Domain := flag.String("D", "", "Provide the direct domain")
	flag.Parse()
	if *Engine == "google" {
		if *query != "" {
			google.GoogleSearch(*query, *countryCode)
		} else {
			fmt.Println("Provide the Query")
			os.Exit(2)
		}
	} else if *Engine == "duckduckgo" {
		duckduckgo.DuckDuckSearch(*query)

	} else {
		resp, err := http.Get("https://duckduckgo.com/html/?q=golang+developer")
		if err != nil {
			log.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(body))
		fmt.Println("Provide the Search Engine")
		os.Exit(1)

	}

	time.Sleep(time.Second * 30)
}
