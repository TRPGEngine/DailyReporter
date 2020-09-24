package main

import (
	"fmt"
	"github.com/TRPGEngine/DailyReporter/report"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	apikey := os.Getenv("WAKATIME_APIKEY")
	wakaReport := report.GetWakatimeCodingReport(apikey)

	username := os.Getenv("USERNAME")
	githubReport := report.GetGithubActiveReport(username)

	reportText := wakaReport + "---------\n" + githubReport

	fmt.Print(reportText)
}
