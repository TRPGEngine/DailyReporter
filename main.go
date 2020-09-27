package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TRPGEngine/DailyReporter/report"
	_ "github.com/joho/godotenv/autoload"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func sendReport(targetUrl string, reportText string) {
	data := make(map[string]string)
	data["msg"] = reportText
	bytesData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(targetUrl, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	botUrl := os.Getenv("BOT_URL")

	if botUrl == "" {
		panic("缺少机器人Url")
	}

	apikey := os.Getenv("WAKATIME_APIKEY")
	wakaReport := report.GetWakatimeCodingReport(apikey)

	username := os.Getenv("USERNAME")
	githubReport := report.GetGithubActiveReport(username)

	weatherReport := report.GetWeatherReport()

	reportText := strings.Join([]string{wakaReport, githubReport, weatherReport}, "---------\n")

	sendReport(botUrl, reportText)
}
