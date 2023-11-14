package throttledtransport

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Transport struct {
	Base    http.RoundTripper
	Limiter *rate.Limiter
}

// client := &http.Client{
// 	Transport: throttled.New(http.DefaultTransport, rate.NewLimiter(rate.Limit(1), 1))
// }

func New(base http.RoundTripper, limiter *rate.Limiter) *Transport {
	return &Transport{
		Base:    base,
		Limiter: limiter,
	}
}

// RoundTrip implementation with rate limiting
func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	res := t.Limiter.Reserve()

	select {
	case <-time.After(res.Delay()):
		return t.Base.RoundTrip(r)
	case <-r.Context().Done():
		res.Cancel()
		return nil, r.Context().Err()
	}
}
