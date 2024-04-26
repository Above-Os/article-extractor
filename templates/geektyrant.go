package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) GeektyrantScrapContent(document *goquery.Document) string {

	contents := ""

	/*document.Find("figure.iframe-embed").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})*/
	document.Find("div.sqs-block-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
