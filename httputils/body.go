package httputils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// CopyRequest copies an http.Request in a shallow fashion. All fields are shared except for Body
// which is read fully from the request and Header which is copied..
func CopyRequest(r *http.Request) (*http.Request, error) {
	body := r.Body
	var err error

	//TODO(tmc): avoid copy in the bytes.Buffer case
	if body, err = FreezeBody(r); err != nil {
		return nil, err
	}

	u, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, err
	}

	return &http.Request{
		Method:     r.Method,
		URL:        u,
		Proto:      r.Proto,
		ProtoMajor: r.ProtoMajor,
		ProtoMinor: r.ProtoMinor,
		Header:     r.Header,
		Body:       body,
		Host:       r.Host,
	}, nil
}

func CopyHeader(h http.Header) http.Header {
	result := http.Header{}
	for h, vals := range h {
		for _, v := range vals {
			result.Add(h, v)
		}
	}
	return result
}

func FreezeBody(r *http.Request) (io.ReadCloser, error) {
	var body io.ReadCloser
	var err error
	body, r.Body, err = drainBody(r.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

// drainBody reads the given buffer and returns two ReadClosers that point to the read data.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, nil, err
	}
	if err = b.Close(); err != nil {
		return nil, nil, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
