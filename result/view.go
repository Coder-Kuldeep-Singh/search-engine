package result

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

func ResultParser(response *http.Response, class1, class2, class3, class4 string) ([]Result, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}

	results := []Result{}
	sel := doc.Find(class1)
	rank := 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find(class2)
		link, _ := linkTag.Attr("href")
		titleTag := item.Find(class3)
		descTag := item.Find(class4)
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" {
			result := Result{
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
