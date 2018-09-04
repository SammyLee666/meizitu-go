package parser

import (
	"regexp"
	"bytes"
)

var (
	ItemReg = regexp.MustCompile(`<a href="(http://www.meizitu.com/a/\d+.html)"  target='_blank'>(.+?)</a>`)
	NextReg = regexp.MustCompile(`<li><a href='(.+?)'>下一页</a></li>`)
)

func List(body []byte) ParserResult {
	next, ok := extractString(NextReg, body)
	result := ParserResult{}

	images, ok := extractString(ItemReg, body)
	for _, v := range images {
		result.Items = append(result.Items, string(v[2]))
		result.Request = append(result.Request, Request{
			Url:        string(v[1]),
			ParserFunc: Image,
		})
	}


	if ok {
		baseNextUrl := JoinBaseUrl(string(next[0][1]))
		result.Request = append(result.Request, Request{
			Url:        baseNextUrl,
			ParserFunc: List,
		})
	}

	return result
}

func extractString(reg *regexp.Regexp, body []byte) ([][][]byte, bool) {
	match := reg.FindAllSubmatch(body, -1)
	if len(match) == 0 {
		return nil, false
	}
	return match, true
}

func JoinBaseUrl(url string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://www.meizitu.com/a/")
	buffer.WriteString(url)
	return buffer.String()
}
