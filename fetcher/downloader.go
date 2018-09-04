package fetcher

import (
	"net/http"
	"log"
	"regexp"
	"os"
	"io"
)

func ImgDownloader(url string, dir string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("NewRequest ERR: ", err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do ERR: ", err)
	}
	if resp.StatusCode != 200 {
		log.Printf("Get Url: %s ERR! ,ERR_CODE:%d\n", url, resp.StatusCode)
	}
	defer resp.Body.Close()
	name := getName(url)
	autoMakeDirAndPutImg(name, dir, resp.Body)
}

func getName(url string) (string) {
	nameReg := regexp.MustCompile(`uploads/.+?/.+?/.+?/(.+?\..+$)`)
	nameMatch := nameReg.FindStringSubmatch(url)
	return nameMatch[1]
}

func autoMakeDirAndPutImg(name string, dir string, img io.Reader) {
	legalDir := legalDir(dir)

	ok, err := PathExists(legalDir)
	if err != nil {
		log.Println("PathExists ERR: ", err)
	}

	if !ok {
		os.Mkdir(legalDir, 0777)
	}

	file, err := os.Create(legalDir + "/" + name)
	if err != nil {
		log.Println("os.Create ERR: ", err)
	}

	io.Copy(file, img)

}

func legalDir(dir string) string {
	reg := regexp.MustCompile(` \| 妹子图`)
	return reg.ReplaceAllString(dir, "")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
