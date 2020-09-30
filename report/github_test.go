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

func TestGetCommitTitle(t *testing.T) {
  case1 := getCommitTitle("title")
  if case1 != "title" {
    t.Fail()
  }

  case2 := getCommitTitle("title\n\nconnnect")
  if case2 != "title" {
    t.Fail()
  }
}
