package httputils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

var _ Response = (*TeeResponseWriter)(nil)

// TeeResponseWriter is an http.ResponseWriter that saves the reponse and status code for later inspection.
type TeeResponseWriter struct {
	http.ResponseWriter
	Body    *bytes.Buffer
	Code    int
	Flushed bool
}

// NewTeeResponseWriter creates a TeeResponseWriter from an http.ResponseWriter
func NewTeeResponseWriter(w http.ResponseWriter) *TeeResponseWriter {
	return &TeeResponseWriter{
		ResponseWriter: w,
		Body:           new(bytes.Buffer),
	}
}

// Write writes the provided byte slice to the underlying ResponseWriter as well as the the internal buffer.
func (w *TeeResponseWriter) Write(b []byte) (int, error) {
	if n, err := w.ResponseWriter.Write(b); err != nil {
		return n, err
	}
	return w.Body.Write(b)
}

// WriteHeader writes the provided status code to the underlying ResponseWriter and saves it
func (w *TeeResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.Code = code
}

// Flush flushes the underlying ResponseWriter if it is an io.Flusher and saves that it has been flushed.
func (w *TeeResponseWriter) Flush() {
	if flusher, ok := w.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
	w.Flushed = true
}

// Response interface adherence

// ResponseBody returns the saved response body.
func (w *TeeResponseWriter) ResponseBody() io.ReadCloser {
	return ioutil.NopCloser(w.Body)
}

// StatusCode returns the saved status code.
func (w *TeeResponseWriter) StatusCode() int {
	return w.Code
}
