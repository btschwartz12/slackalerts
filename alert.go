package slackalerts

import (
	"context"
	"fmt"
	"time"
)

func SendAlert(ctx context.Context, slackWebhookUrl, title string, blocks []Block) error {
	if _, ok := ctx.Deadline(); !ok {
		defaultCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		ctx = defaultCtx
	}

	err := sendToSlack(slackWebhookUrl, title, blocks)
	if err != nil {
		return fmt.Errorf("failed to send alert: %w", err)
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("failed to send message to Slack: %w", ctx.Err())
	default:
		return nil
	}
}
