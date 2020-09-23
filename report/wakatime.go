package report

import (
	"encoding/base64"
	"github.com/aquilax/go-wakatime"
	"net/http"
	"strconv"
	"time"
)

type MyTransport struct {
	http.Transport
	apikey string
}

func (t *MyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(t.apikey)))

	resp, err := http.DefaultTransport.RoundTrip(req)
	return resp, err
}

func getChineseTimeText(total wakatime.SummaryGrandTotal) string {
	str := ""
	if total.Hours > 0 {
		str += strconv.Itoa(total.Hours) + "时 "
	}

	str += strconv.Itoa(total.Minutes) + "分"

	return str
}

// 获取wakatime的汇报信息
func GetWakatimeCodingReport(apikey string) string {
	client := wakatime.New(&MyTransport{apikey: apikey})
	summary, err := client.Summaries(wakatime.CurrentUser, time.Now().AddDate(0, 0, -1).UTC(), time.Now().UTC(), nil, nil)
	if err != nil {
		panic(err)
	}

	yesterday := summary.Data[0]

	reportText := yesterday.Range.Date + "汇总 总计编码时间: " + getChineseTimeText(yesterday.GrandTotal) + "\n"

	for _, p := range yesterday.Projects {
		reportText += p.Name + ": " + getChineseTimeText(p.SummaryGrandTotal) + "\n"
	}

	return reportText
}
