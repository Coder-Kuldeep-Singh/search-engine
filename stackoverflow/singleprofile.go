package stackoverflow

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/googlesearch/client"
)

type UsersProfile struct {
	Rank int
	Name string
}

// https://stackoverflow.com/users/5928186/shubham-khatri?tab=profile
func UserProfileUrl(url string) string {
	tab := "profile" // profile || activity
	return fmt.Sprintf("%s?tab=%s", url, tab)
}

func MainProfiles(response *http.Response, mainClass, username string) ([]UsersProfile, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}

	results := []UsersProfile{}
	sel := doc.Find(mainClass)
	rank := 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		usernameTag := item.Find(username)

		name := usernameTag.Text()
		result := UsersProfile{
			rank,
			name,
		}
		results = append(results, result)
		rank += 1
	}
	return results, err
}

func MemberProfile(url string) {
	mainuser := UserProfileUrl(url)
	fmt.Println(mainuser)

	res, err := client.Request(mainuser)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	time.Sleep(3 * time.Second)
	scrapes, err := MainProfiles(res, "div.main-content", "div.grid--cell")
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range scrapes {
		fmt.Println("-----------------------user main profile------------------------------------")
		fmt.Println(user.Name)
		fmt.Println("-----------------------user main profile------------------------------------")

	}
	time.Sleep(3 * time.Second)
}
