package withTime

import (
	"context"
	"time"
)

func WithTimeout(timeoutLength int) (context.Context, context.CancelFunc) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeoutLength))

	return ctxTimeout, cancel
}