package httputils

import (
	"net/http"
	"time"
)

// RequestResponse is a Request-Response pair
type RequestResponse struct {
	Request   *http.Request
	Response  Response
	StartedAt time.Time
	EndedAt   time.Time
	done      chan struct{}
}

// Creates a new RequestResponse. Sets StartedAt to time.Now()
func NewRequestResponse(req *http.Request, resp Response) *RequestResponse {
	return &RequestResponse{
		Request:   req,
		Response:  resp,
		StartedAt: time.Now(),
		done:      make(chan struct{}),
	}
}

func (r *RequestResponse) MarkDone() {
	r.EndedAt = time.Now()
	close(r.done)
}

func (r *RequestResponse) Done() <-chan struct{} {
	return r.done
}
