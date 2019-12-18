package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nlopes/slack"
)

func handleRequest(ctx context.Context) (string, error) {
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
