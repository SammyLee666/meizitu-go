package fetcher

import (
	"net/http"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"log"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Printf("Get Url: %s ERR! ,ERR_CODE:%d\n", url, resp.StatusCode)
		return nil, err
	}
	defer resp.Body.Close()
	body := ToUTF8(resp.Body)
	return body, nil

}

func ToUTF8(r io.Reader) []byte {
	reader, err := charset.NewReader(r, "")
	if err != nil {
		return nil
	}
	bytes, err := ioutil.ReadAll(reader)
	if err != nil{
		return nil
	}
	return bytes
}
