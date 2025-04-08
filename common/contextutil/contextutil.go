package contextutil

import (
	"context"
	"time"
)

func WithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
