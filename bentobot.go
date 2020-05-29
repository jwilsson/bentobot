package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

func handleRequest() (string, error) {
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
