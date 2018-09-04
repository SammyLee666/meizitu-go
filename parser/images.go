package parser

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"Img/fetcher"
)

func Image(body []byte) ParserResult {
	images, name := parserImages(body)
	for _, image := range images {
		fetcher.ImgDownloader(image, name)
	}
	return ParserResult{}
}

func parserImages(body []byte) (result []string, name string) {
	reader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
	}
	doc.Find(".postContent p img").Each(func(i int, selection *goquery.Selection) {
		val, _ := selection.Attr("src")
		result = append(result, val)
	})

	name = doc.Find("title").Text()

	return
}
