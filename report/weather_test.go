package report

import "testing"

func TestGetWeatherReport(t *testing.T) {
	weatherReport := GetWeatherReport()

	println(weatherReport)
}
