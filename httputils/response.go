package httputils

import (
	"io"
	"net/http"
)

// Response is an HTTP response.
type Response interface {
	ResponseBody() io.ReadCloser
	StatusCode() int
	Header() http.Header
}

var _ Response = (*rawResponse)(nil)

type rawResponse struct {
	r *http.Response
}

// NewRawResponse consumes an http.Response by reading and closing it's body and replacing it with a byte.Buffer.
func NewRawResponse(r *http.Response) (Response, error) {
	var err error
	if r.Body != nil {
		_, r.Body, err = drainBody(r.Body)
	}
	return &rawResponse{r}, err
}

func (r *rawResponse) Header() http.Header {
	return r.r.Header
}

func (r *rawResponse) ResponseBody() io.ReadCloser {
	return r.r.Body
}

func (r *rawResponse) StatusCode() int {
	return r.r.StatusCode
}
