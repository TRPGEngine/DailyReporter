package report

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
	"time"
)

func getEventTypeChineseName(eventType string) string {
	switch eventType {
	case "PushEvent":
		return "提交"
	case "PullRequestEvent":
		return "请求合并"
	case "CreateEvent":
		return "创建"
	case "ForkEvent":
		return "Fork"
	default:
		return eventType
	}
}

func getEventPayloadText(event github.Event) string {
	return ""
}

func GetGithubActiveReport(username string) string {
	client := github.NewClient(nil)

	events, _, err := client.Activity.ListEventsPerformedByUser(context.Background(), username, true, &github.ListOptions{Page: 1, PerPage: 10})
	if err != nil {
		panic(err)
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	yesterday := today.AddDate(0, 0, -1)
	text := ""

	for _, event := range events {
		createdAt := event.GetCreatedAt()

		if createdAt.Unix() < today.Unix() && createdAt.Unix() > yesterday.Unix() {
			payload, err := event.ParsePayload()
			if err != nil {
				panic(err)
			}

			fmt.Printf("%+v", payload)

			eventType := getEventTypeChineseName(event.GetType())
			eventRepoName := event.GetRepo().GetName()
			text += fmt.Sprintf("%s %s\n", eventType, eventRepoName)
		}

	}

	return text
}
