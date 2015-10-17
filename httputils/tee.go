package httputils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

var _ Response = (*TeeResponseWriter)(nil)

type TeeResponseWriter struct {
	http.ResponseWriter
	Body    *bytes.Buffer
	Code    int
	Flushed bool
}

func NewTeeResponseWriter(w http.ResponseWriter) *TeeResponseWriter {
	return &TeeResponseWriter{
		ResponseWriter: w,
		Body:           new(bytes.Buffer),
	}
}

func (w *TeeResponseWriter) Write(b []byte) (int, error) {
	if n, err := w.ResponseWriter.Write(b); err != nil {
		return n, err
	}
	return w.Body.Write(b)
}

func (w *TeeResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.Code = code
}

func (w *TeeResponseWriter) Flush() {
	w.Flushed = true
}

// Response interface adherence

func (w *TeeResponseWriter) IsFlushed() bool {
	return w.Flushed
}

func (w *TeeResponseWriter) ResponseBody() io.ReadCloser {
	return ioutil.NopCloser(w.Body)
}

func (w *TeeResponseWriter) StatusCode() int {
	return w.Code
}
