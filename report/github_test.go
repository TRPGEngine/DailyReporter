package report

import (
	"testing"
)

func TestGetGithubActiveReport(t *testing.T) {
	text := GetGithubActiveReport("moonrailgun")

	if text == "" {
		t.Fail()
	}
}
