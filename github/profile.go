package github

import (
	"fmt"
	"strings"
)

func Profile(profile_link string) {
	repo_to_profile := strings.Split(profile_link, "/")
	protocol := ReplaceSpace(repo_to_profile[0])
	domain := ReplaceSpace(repo_to_profile[2])
	username := ReplaceSpace(repo_to_profile[4])
	fmt.Printf("%s//%s/%s\n", protocol, domain, username)
}

func ReplaceSpace(value string) string {
	return strings.Replace(value, " ", "", -1)
}
