package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Body struct {
	Text string `json:"text"`
}

func handleRequest(ctx context.Context) (string, error) {
	jsonBody, _ := json.Marshal(Body{
		Text: ":bento:?",
	})

	_, err := http.Post(
		os.Getenv("SLACK_WEBHOOK_URL"),
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	return "", err
}

func main() {
	lambda.Start(handleRequest)
}
