package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Gettranslate(s string) string {
	params := url.Values{}
	URL, _ := url.Parse("http://fanyi.youdao.com/translate")
	params.Set("doctype", "json")
	params.Set("type", "ZH_CN")
	params.Set("i", s)
	URL.RawQuery = params.Encode()
	urlPath := URL.String()
	resp, error := http.Get(urlPath)
	if error != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
