package main

import (
	"fmt"
	"github.com/TRPGEngine/DailyReporter/report"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func GetEnvConfig(key string) string {
	return os.Getenv(key)
}

func main() {
	apikey := GetEnvConfig("WAKATIME_APIKEY")
	wakaReport := report.GetWakatimeCodingReport(apikey)

	username := GetEnvConfig("USERNAME")
	githubReport := report.GetGithubActiveReport(username)

	fmt.Print(wakaReport)
	fmt.Print(githubReport)
}
