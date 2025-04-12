package parser

import (
	"golang.org/x/net/html"
	"log"
	"strings"

	"github.com/gookit/goutil/arrutil"
)

func GetImagesFromHtml(page string) ([]string, error) {
	var res []string

	doc, err := html.Parse(strings.NewReader(page))
	if err != nil {
		log.Printf("Ошибка при парсинге HTML: %v", err)
		return res, err
	}

	res = arrutil.StringsUnique(findImgTags(doc))

	return res, err
}

func findImgTags(n *html.Node) []string {
	var srcList []string

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				srcList = append(srcList, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		srcList = append(srcList, findImgTags(c)...)
	}

	return srcList
}
