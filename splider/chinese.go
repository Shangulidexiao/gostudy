package splider

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func Exec() {
	htmlNode, err := download("http://daymenu.cn")

	if err != nil {
		log.Fatal(err)
	}

	outline(nil, htmlNode)
}
func findLinks(links []string, htmlNode *html.Node) []string {
	if url := link(htmlNode); url != "" {
		links = append(links, url)
	}
	for c := htmlNode.FirstChild; c != nil; c = c.NextSibling {
		links = findLinks(links, c)
	}
	return links
}

func link(htmlNode *html.Node) (url string) {
	if strings.ToLower(htmlNode.Data) != "img" || htmlNode.Type != html.ElementNode {
		return
	}
	for _, a := range htmlNode.Attr {
		if strings.ToLower(a.Key) == "src" {
			url = a.Val
			break
		}
	}
	return
}

func download(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return &html.Node{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return &html.Node{}, fmt.Errorf("访问 %s:%s", url, response.Status)
	}
	htmlNode, err := html.Parse(response.Body)
	if err != nil {
		return &html.Node{}, err
	}
	return htmlNode, nil
}

func outline(stack []string, htmlNode *html.Node) {
	if htmlNode.Type == html.ElementNode {
		stack = append(stack, htmlNode.Data)
		fmt.Println(stack)
	}

	for c := htmlNode.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
