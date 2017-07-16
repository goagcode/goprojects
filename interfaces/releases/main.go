package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ReleaseInfo struct {
	Id      uint   `json:"id"`
	TagName string `json:"tag_name"`
}

func getLatestReleaseTag(repo string) (string, error) {
	apiUrl, err := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	response, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	releases := []ReleaseInfo{}
	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}
	tag := releases[0].TagName
	return tag, nil
}

func getReleaseTagMessage(repo string) (string, error) {
	tag, err := getLatestReleaseTag(repo)
	if err != nil {
		return "", fmt.Errorf("Error querying Github API: %s", err)
	}
	return fmt.Sprintf("The latest release is %s", tag), nil
}

func main() {
	msg, err := getReleaseTagMessage("/docker/machine")
	if err != nil {
		fmt.Println(os.Stderr, msg)
	}
	fmt.Println(msg)
}
