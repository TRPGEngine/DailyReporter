package main

import (
	"fmt"
	"github.com/TRPGEngine/DailyReporter/report"
	"os"
)

func GetEnvConfig(key string) string {
	return os.Getenv(key)
}

func main() {
	apikey := GetEnvConfig("WAKATIME_APIKEY")

	wakaReport := report.GetWakatimeCodingReport(apikey)

	fmt.Print(wakaReport)
}
