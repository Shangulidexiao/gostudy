package ch4

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var issuseUrl = "https://api.github.com/search/issues?q=vue"

type IssuseResult struct {
	TotalCount int       `json:"total_count"`
	Item       []*Issuse `json:"items"`
}

type Issuse struct {
	Number        int
	HTMLURL       string `json:"html_url"`
	RepositoryUrl string `json:"repository_url"`
	Title         string
	CreatedAt     time.Time `json:"created_at"`
	//Body          string
	User *User
}

type User struct {
	HTMLURL string `json:"html_url"`
	Login   string `json:"login"`
}

func GetIssues() IssuseResult {
	var result IssuseResult

	resp, err := http.Get(issuseUrl)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	return result
}

func ExecIssuse() {
	var result IssuseResult = GetIssues()
	tpl, err := ioutil.ReadFile(`E:\go\src\gostudy\ch4\issuse.html`)
	if err != nil {

	}

	t, err := template.New(`a`).Parse(string(tpl))
	t.Execute(os.Stdout, result)
}
