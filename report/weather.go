package report

import (
	"io/ioutil"
	"net/http"
)

// https://wttr.in/Shanghai?lang=zh-cn&Tq0
func GetWeatherReport() string {
	req, err := http.NewRequest("GET", "https://wttr.in/Shanghai?lang=zh-cn&mTq0", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "plain/text")
	req.Header.Add("User-Agent", "curl")
	req.Proto = ""

	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
