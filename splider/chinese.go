package splider

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"

	"golang.org/x/net/html"
	"golang.org/x/text/transform"
)

const domain = "http://xh.5156edu.com"
const initUrl = domain + "/pinyi.html"

func ExecSplider() {
	response, err := http.Get(initUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	doc, err := html.Parse(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(visit(nil, doc))

}
func visit(links []map[string]string, n *html.Node) []map[string]string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" && strings.HasPrefix(attr.Val, "html2") {
				links = append(links, map[string]string{gbktoutf8(n.FirstChild.Data): attr.Val})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func chinese(url string) {

}
func insertDb(pin, han string) bool {

	return true
}

func exists(han string) bool {
	return true
}

func gbktoutf8(gbk string) string {
	utf8, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(gbk)), simplifiedchinese.GBK.NewEncoder()))
	return string(utf8)
}
