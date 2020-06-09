package stackoverflow

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/googlesearch/client"
)

type Users struct {
	Rank       int
	Link       string
	Name       string
	Reputation string
	Language   string
	Image      string
}

// https://stackoverflow.com/users?page=2&tab=reputation&filter=all
var StackUserUrl = "https://stackoverflow.com/users?page="

func StackAllProfilesUrl(pages int) string {
	tab := "reputation" // reputation || voters || editors
	filters := "all"    // week ||  month || quarter || year || all
	return fmt.Sprintf("%s%d&tab=%s&filter=%s", StackUserUrl, pages, tab, filters)
}

func StackPaginationUrl() string {
	url := "https://stackoverflow.com/users?tab="
	tab := "reputation" // reputation || voters || editors
	filters := "all"    // week ||  month || quarter || year || all
	return fmt.Sprintf("%s%s&filter=%s", url, tab, filters)
}

func StackNewProfileUrl() string {
	tab := "newusers"
	return fmt.Sprintf("%s%s", StackUserUrl, tab)
}

func StackModeratorsUrl() string {
	tab := "moderators"
	return fmt.Sprintf("%s%s", StackUserUrl, tab)
}

func StackOverFlowPaginationUsers(paginationpath string) (int, error) {
	users := StackPaginationUrl()
	fmt.Println(users)
	res, err := client.Request(users)
	if err != nil {
		// log.Println(err)
		return 0, err

	}
	defer res.Body.Close()
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

func Profiles(response *http.Response, mainClass, profilelink, username, userlocation, reputation, languages, userimage string) ([]Users, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}

	results := []Users{}
	sel := doc.Find(mainClass)
	rank := 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		profileTag := item.Find(profilelink)
		link, _ := profileTag.Attr("href")
		usernameTag := item.Find(username)
		userlocationTag := item.Find(userlocation)
		reputationTag := item.Find(reputation)
		languageTag := item.Find(languages)
		userimageTag := item.Find(userimage)
		user, _ := userimageTag.Attr("src")

		language := languageTag.Text()

		locations := userlocationTag.Text()
		locations = strings.Replace(locations, "\n", "", -1)
		locations = strings.Replace(locations, "\t", "", -1)

		reputations := reputationTag.Text()
		user = strings.Trim(user, " ")

		name := usernameTag.Text()
		link = strings.Trim(link, " ")

		if link != "" && link != "#" {
			result := Users{
				rank,
				link,
				name,
				reputations,
				language,
				user,
			}
			results = append(results, result)
			rank += 1
		}
	}
	return results, err
}

func ProfileResults() {
	pagination, err := StackOverFlowPaginationUsers("#user-browser > div.s-pagination.pager.fr > a:nth-child(7)")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("-----------------------------New Line--------------------------------------")
	// CreateFolder()
	// createFile, err := GenerateFile()
	if err != nil {
		log.Printf("Error to create file %v\n", err)
	}
	// fmt.Println(pagination)
	for i := 1; i <= pagination; i++ {
		userpage := StackAllProfilesUrl(i)
		fmt.Println(userpage)
		res, err := client.Request(userpage)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		time.Sleep(3 * time.Second)
		scrapes, err := Profiles(res, "div.user-info", "div.user-details > a", "div.user-details > a", "div.user-details > span.user-location", "div.user-details > div.-flair > span.reputation-score", "div.user-tags > a", "div.user-gravatar48 > a > div.gravatar-wrapper-48 > img")
		if err != nil {
			fmt.Println(err)
		}
		for _, results := range scrapes {
			fmt.Println("-----------------------------New Line--------------------------------------")
			result := results.Image + "\n"
			fmt.Println(result)
			// _, err = AppendData(createFile, result)
			// if err != nil {
			// 	log.Printf("Error to append data into file %v\n", err)
			// }
			url := fmt.Sprintf("https://stackoverflow.com%s", results.Link)
			mainurl := UserProfileUrl(url)
			fmt.Println(mainurl)

			fmt.Println("-----------------------------New Line--------------------------------------")
		}
		time.Sleep(3 * time.Second)

	}
}

//GenerateFile function generating file
func GenerateFile() (*os.File, error) {
	create, err := os.Create("./images/paths.txt")
	if err != nil {
		return nil, err
	}
	return create, nil
}

//AppendData function appends data into files we created in previous function
func AppendData(file *os.File, data string) (string, error) {
	size, err := file.WriteString(data)
	if err != nil {
		return "", err
	}
	return string(size), nil
}

//CreateFolder function creates the folder and also checks
//If there is any that name folder exists or not
func CreateFolder() {
	_, err := os.Stat("images")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("images", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
