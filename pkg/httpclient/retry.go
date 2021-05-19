package httpclient

import (
	"context"
	"net/http"
	"time"
)

const (
	// DefaultRetryTimes If the request fails, retry at most 3 times
	DefaultRetryTimes = 3
	// DefaultRetryDelay waits 100 milliseconds before retrying
	DefaultRetryDelay = time.Millisecond * 100
)

// Verify parse the body and verify that it is correct
type RetryVerify func(body []byte) (shouldRetry bool)

func shouldRetry(ctx context.Context, httpCode int) bool {
	select {
	case <-ctx.Done():
		return false
	default:
	}

	switch httpCode {
	case
		_StatusReadRespErr,
		_StatusDoReqErr,

		http.StatusRequestTimeout,
		http.StatusLocked,
		http.StatusTooEarly,
		http.StatusTooManyRequests,

		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:

		return true

	default:
		return false
	}
}
