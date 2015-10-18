package splitter

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"golang.org/x/net/context"

	"github.com/tmc/prism/config"
	"github.com/tmc/prism/httputils"
)

// defaultBufferLen is the channel capacity of ongoing request/response pairs
// this value is used if Config.RequestBufferSize is zero.
const defaultBufferLen = 1

// Splitter is the type that manages an upstream and a number of downstream sinks.
type Splitter struct {
	Config config.SplitterConfig
	Client *http.Client

	upstreamProxy *httputil.ReverseProxy
	upstream      *url.URL

	requests chan *httputils.RequestResponse
}

// NewSplitter constructs a splitter described by a config.SplitterConfig
func NewSplitter(config config.SplitterConfig) (*Splitter, error) {
	bufferLen := config.RequestBufferSize
	if bufferLen == 0 {
		bufferLen = defaultBufferLen
	}
	s := &Splitter{
		Config:   config,
		Client:   http.DefaultClient,
		requests: make(chan *httputils.RequestResponse, bufferLen),
	}
	s.upstreamProxy = &httputil.ReverseProxy{
		Director: s.requestDirector,
	}
	var err error
	s.upstream, err = url.Parse(s.Config.Source)

	return s, err
}

// Start starts an http listener on the address in SplitterConfig.ListenAddress
func (s *Splitter) Start(ctx context.Context) error {
	fmt.Println("Starting", s.Config.Label, "on", s.Config.ListenAddr)
	go s.handleRequests(ctx)
	return http.ListenAndServe(s.Config.ListenAddr, s)
}

// ServeHTTP satisfies the net/http.Handler interface
func (s *Splitter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	teeRW := httputils.NewTeeResponseWriter(rw)

	r2, err := httputils.CopyRequest(r)

	// on request copy error attempt and early exit (which may be invalid, unfortunately)
	if err != nil {
		log.Println("splitter: error copying request:", err)
		s.upstreamProxy.ServeHTTP(teeRW, r)
		return
	}

	reqresp := httputils.NewRequestResponse(r2, teeRW)
	defer reqresp.MarkDone()

	select {
	case s.requests <- reqresp:
	default:
		log.Println("splitter: dropped request on the floor")
	}
	s.upstreamProxy.ServeHTTP(teeRW, r)
}

func (s *Splitter) requestDirector(r *http.Request) {
	r.URL.Scheme = s.upstream.Scheme
	r.URL.Host = s.upstream.Host
}

func (s *Splitter) handleRequests(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("satisfying cancellation request")
			return
		case r := <-s.requests:
			go s.handleRequest(ctx, r)
		}
	}
}

func (s *Splitter) handleRequest(ctx context.Context, reqresp *httputils.RequestResponse) {
	log.Println("got request to split:", reqresp)
	if s.Config.WaitForResponse {
		log.Println("waiting for response")
		select {
		case <-ctx.Done():
			return
		case <-reqresp.Done():
		}
		log.Println("done waiting for response")
	}

	sinkResponses := make(chan *httputils.RequestResponse, len(s.Config.Sinks))
	responsesDone := make(chan struct{})
	wg := sync.WaitGroup{}

	ctx, _ = context.WithTimeout(ctx, s.Config.SinkRequestTimeout)
	for _, sink := range s.Config.Sinks {
		wg.Add(1)
		req, err := httputils.CopyRequest(reqresp.Request)
		if err != nil {
			log.Println("failed to copy request:", err)
		}
		go func(sink config.Sink, req *http.Request) {
			defer wg.Done()
			sinkResponse, err := s.performSinkRequest(ctx, req, sink)
			if err != nil {
				log.Println("error performing request to", sink, "-", err)
				return
			}
			sinkResponses <- sinkResponse
		}(sink, req)
	}
	go func() {
		wg.Wait()
		close(responsesDone)
		close(sinkResponses)
	}()

	responses := []*httputils.RequestResponse{}
	select {
	case <-ctx.Done():
		log.Println("handleRequest: got cancellation signal")
	case <-responsesDone:
	}
	for response := range sinkResponses {
		responses = append(responses, response)
		log.Println("status:", response.Response.StatusCode())
	}

	// TODO(tmc): compare responses, record information
	log.Println(len(responses), "shadow requests performed")
}

func (s *Splitter) performSinkRequest(ctx context.Context, req *http.Request, sink config.Sink) (*httputils.RequestResponse, error) {
	result := httputils.NewRequestResponse(req, nil)
	// swap in sink address
	req.URL.Host = sink.Addr
	// assume "http" if no scheme is present
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	result.Response, err = httputils.NewRawResponse(resp)
	return result, err
}
