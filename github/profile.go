package github

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Profile(profile_link string, file *os.File) {
	repo_to_profile := strings.Split(profile_link, "/")
	protocol := ReplaceSpace(repo_to_profile[0])
	domain := ReplaceSpace(repo_to_profile[2])
	username := ReplaceSpace(repo_to_profile[4])
	url := fmt.Sprintf("%s//%s/%s\n", protocol, domain, username)
	_, err := AppendData(file, url)
	if err != nil {
		log.Printf("Error to append data into file %v\n", err)
	}
	fmt.Println(url)
}

func ReplaceSpace(value string) string {
	return strings.Replace(value, " ", "", -1)
}

//GenerateFile function generating file
func GenerateFile() (*os.File, error) {
	create, err := os.Create("./github-profiles/profiles.txt")
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
	_, err := os.Stat("github-profiles")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("github-profiles", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
