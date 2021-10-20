package repos

import (
	"context"
	"time"
)

func getContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	return ctx, cancel
}