package main

import (
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/jwilsson/go-bot-utils"
	"github.com/slack-go/slack"
)

func handleRequest() (string, error) {
	if !utils.ShouldRun(time.Now(), os.Getenv("RUN_AT_TIME"), os.Getenv("TARGET_TIMEZONE")) {
		return "", nil
	}

	err := slack.PostWebhook(
		os.Getenv("SLACK_WEBHOOK_URL"),
		&slack.WebhookMessage{
			Text: ":bento:?",
		},
	)

	return "", err
}

func main() {
	lambda.Start(handleRequest)
}
