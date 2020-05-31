package stackoverflow

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

type StackoverflowQuestionSearch struct {
	ResultRank     int
	ResultURL      string
	ResultTitle    string
	ResultDesc     string
	ResultLanguage string
	ResultAsked    string
	ResultUser     string
	ResultAnswer   string
	ResultVotes    string
}

var StackOverflowDomains = "https://stackoverflow.com/search?"

// https://stackoverflow.com/search?tab=active&pagesize=50&q=golang
func buildStackUrl(searchTerm string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	tabs := "active"    // relevance || votes || newest
	searchlimit := "50" // 10 || 30
	return fmt.Sprintf("%stab=%s&pagesize=%s&q=%s", StackOverflowDomains, tabs, searchlimit, searchTerm)
}

func StackOverFlowPagination(searchTerm, paginationpath string) (int, error) {
	stackurl := buildStackUrl(searchTerm)
	fmt.Println(stackurl)
	res, err := client.Request(stackurl)
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
	pagination := doc.Find(paginationpath).Text()
	// fmt.Println(pagination)
	number, _ := strconv.Atoi(pagination)
	return number, nil
}

func ResultParser(response *http.Response, class1, class2, class3, class4, class5, class6, class7, class8, class9 string) ([]StackoverflowQuestionSearch, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}

	results := []StackoverflowQuestionSearch{}
	sel := doc.Find(class1)
	rank := 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find(class2)
		link, _ := linkTag.Attr("href")
		titleTag := item.Find(class3)
		descTag := item.Find(class4)
		languageTag := item.Find(class5)
		askedTag := item.Find(class6)
		userTag := item.Find(class7)
		user, _ := userTag.Attr("href")
		answerTag := item.Find(class8)
		voteTag := item.Find(class9)

		language := languageTag.Text()
		// language = strings.Replace(language, "\n", " ", -1)
		// language = strings.Replace(language, "\t", "", -1)

		ask := askedTag.Text()
		ask = strings.Trim(ask, " ")
		ask = strings.Replace(ask, "\t", "", -1)
		ask = strings.Replace(ask, "\n", "", -1)

		answer := answerTag.Text()
		answer = strings.Trim(answer, " ")
		answer = strings.Replace(answer, "\t", "", -1)
		answer = strings.Replace(answer, "\n", "", -1)

		vote := voteTag.Text()
		vote = strings.Trim(vote, " ")
		vote = strings.Replace(vote, "\t", "", -1)
		vote = strings.Replace(vote, "\n", "", -1)

		desc := descTag.Text()
		desc = strings.Replace(desc, "\n", "", -1)
		desc = strings.Replace(desc, "\t", "", -1)

		title := titleTag.Text()
		link = strings.Trim(link, " ")

		if link != "" && link != "#" {
			result := StackoverflowQuestionSearch{
				rank,
				link,
				title,
				desc,
				language,
				ask,
				user,
				answer,
				vote,
			}
			results = append(results, result)
			rank += 1
		}
	}
	return results, err
}

func StackResults(searchTerm string) ([]StackoverflowQuestionSearch, error) {
	// pagination, err := GithubPagination(searchTerm,"#mainbar > div.s-pagination.pager.fl > a:nth-child(7)")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// if pagination == 0 {
	// 	log.Println("pagination ", 0)
	// 	return
	// }
	// fmt.Printf("%T", pagination)

	githuburl := buildStackUrl(searchTerm)
	fmt.Println(githuburl)
	res, err := client.Request(githuburl)
	if err != nil {
		return nil, err
	}
	scrapes, err := ResultParser(res, "div.question-summary", "div.result-link > h3 > a", "div.result-link > h3 > a", "div.excerpt", "div.tags", " div.started.fr > span", "div.started.fr > a", "div.statscontainer > div.stats > div.status.answered-accepted", "div.statscontainer > div.stats > div.vote > div.votes")
	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func StackOverflowSearch(query string) {
	resp, err := StackResults(query)
	if err != nil {
		log.Println("Having Error to visiting url: ", err)
	}
	for _, stack := range resp {
		fmt.Println()
		fmt.Println()
		fmt.Println("", color.Info(stack.ResultTitle))
		fmt.Printf("https://stackoverflow.com/%s\n", stack.ResultURL)
		fmt.Println("---------------------------------------------------")
		fmt.Println("", stack.ResultDesc)
		fmt.Println("", stack.ResultLanguage)
		fmt.Println("", stack.ResultAsked)
		fmt.Println("", stack.ResultUser)
		fmt.Println("", stack.ResultAnswer)
		fmt.Println("", stack.ResultVotes)
		fmt.Println()
		fmt.Println()
	}

}
