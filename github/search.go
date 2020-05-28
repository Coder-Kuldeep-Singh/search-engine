package github

// https://github.com/search?p=1&q=rust&type=Repositories
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/googlesearch/client"
	"github.com/googlesearch/color"
)

type GithubResult struct {
	ResultRank     int
	ResultURL      string
	ResultTitle    string
	ResultDesc     string
	ResultStar     string
	ResultLanguage string
}

var GithubDomains = "https://github.com/search?p="

func buildGithubUrl(searchTerm string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	return fmt.Sprintf("%s1&q=%s&type=Repositories", GithubDomains, searchTerm)
}

func GithubPagination(searchTerm string) (int, error) {
	githuburl := buildGithubUrl(searchTerm)
	fmt.Println(githuburl)
	res, err := client.Request(githuburl)
	if err != nil {
		// log.Println(err)
		return 0, err
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		// log.Println(err)
		return 0, err
	}
	// Find the review items
	pagination := doc.Find("#js-pjax-container > div > div.col-12.col-md-9.float-left.px-2.pt-3.pt-md-0.codesearch-results > div > div.paginate-container.codesearch-pagination-container > div > a:nth-child(9)").Text()
	number, _ := strconv.Atoi(pagination)
	return number, nil
	//
}

func ResultParser(response *http.Response, class1, class2, class3, class4, class5, class6 string) ([]GithubResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}

	results := []GithubResult{}
	sel := doc.Find(class1)
	rank := 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find(class2)
		link, _ := linkTag.Attr("href")
		titleTag := item.Find(class3)
		descTag := item.Find(class4)
		starTag := item.Find(class5)
		languageTag := item.Find(class6)

		star := starTag.Text()
		star = strings.Replace(star, " ", "", -1)
		star = strings.Replace(star, "\n", "", -1)
		star = strings.Replace(star, "\t", "", -1)

		language := languageTag.Text()
		language = strings.Trim(language, " ")
		language = strings.Replace(language, "\t", "", -1)
		language = strings.Replace(language, "\n", "", -1)

		desc := descTag.Text()
		desc = strings.Replace(desc, "\n", "", -1)
		desc = strings.Replace(desc, "\t", "", -1)

		title := titleTag.Text()
		link = strings.Trim(link, " ")

		if link != "" && link != "#" {
			result := GithubResult{
				rank,
				link,
				title,
				desc,
				star,
				language,
			}
			results = append(results, result)
			rank += 1
		}
	}
	return results, err
}

func GithubResults(searchTerm string) ([]GithubResult, error) {
	// pagination, err := GithubPagination(searchTerm)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// if pagination == 0 {
	// 	log.Println("pagination ", 0)
	// 	return
	// }
	// fmt.Printf("%T", pagination)

	githuburl := buildGithubUrl(searchTerm)
	fmt.Println(githuburl)
	res, err := client.Request(githuburl)
	if err != nil {
		// log.Println(err)
		return nil, err
	}
	scrapes, err := ResultParser(res, "li.repo-list-item > div.mt-n1", "div.text-normal > a", "div.text-normal > a", "p.mb-1", "div > div.mr-3 > a.muted-link", "div > div.mr-3 >  span")

	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func GithubSearch(query string) {
	resp, err := GithubResults(query)
	if err != nil {
		log.Println("Having Error to visiting url: ", err)
	}
	for _, git := range resp {
		fmt.Println()
		fmt.Println()
		// fmt.Println("[+]               ", duck.ResultRank)
		fmt.Println("             ", color.Info(git.ResultTitle))
		fmt.Printf("              https://github.com/%s\n", git.ResultURL)
		fmt.Println("             ---------------------------------------------------")
		fmt.Println("     ", git.ResultDesc)
		fmt.Println("             ", git.ResultStar)
		fmt.Println("         ", git.ResultLanguage)
		fmt.Println()
		fmt.Println()
	}

}
