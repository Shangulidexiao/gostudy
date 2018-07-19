package ch4


var issuseUrl = "https://api.github.com/search/issues"

type IssuseResult struct {
	TotalCount int `json:"total_count"`
	Item *Issuse 
}

type Issuse struct {
	HTMLURL string `json:"url"`
	RepositoryUrl string `json:"repository_url"`

}

