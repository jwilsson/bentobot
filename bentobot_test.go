package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/slack-go/slack"
)

func TestHandleRequest(t *testing.T) {
	expected := ":bento:?"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var message slack.WebhookMessage

		json.NewDecoder(r.Body).Decode(&message)

		if message.Text != expected {
			t.Fatalf("Expected %s, got %s", expected, message.Text)
		}
	}))

	os.Setenv("SLACK_WEBHOOK_URL", ts.URL)

	defer os.Unsetenv("SLACK_WEBHOOK_URL")
	defer ts.Close()

	handleRequest()
}
