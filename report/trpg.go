package report

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TRPGEngineReportResult struct {
	Count struct {
		Chatlog  int32 `json:"chatlog"`
		Login    int32 `json:"login"`
		Register int32 `json:"register"`
	} `json:"count"`
	Result bool `json:"result"`
}

func GetTRPGEngineReport() string {
	yesterday := getYesterday()
	dateStr := yesterday.Format("2006-01-02")

	req, err := http.NewRequest("GET", "https://trpgapi.moonrailgun.com/report/stats/daily?date="+dateStr, nil)
	if err != nil {
		panic(err)
	}

	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	res := new(TRPGEngineReportResult)

	err = json.Unmarshal(bytes, res)
	if err != nil {
		panic(err)
	}

	text := fmt.Sprintf("TRPG Engine %s:\n 登录次数: %d\n 注册次数: %d\n 消息数量: %d\n", dateStr, res.Count.Login, res.Count.Register, res.Count.Chatlog)

	return text
}
