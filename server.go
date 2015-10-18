package prism

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"

	"golang.org/x/net/context"

	"github.com/tmc/prism/config"
	"github.com/tmc/prism/monitor"
	"github.com/tmc/prism/splitter"
)

// Server is the primary top level type in prism. It is responsible for reverse-proxying to the
// upstream and mirroring requests to the sinks.
type Server struct {
	config *config.Config

	upstreamProxy *httputil.ReverseProxy
	monitor       *monitor.Monitor
	splitters     []*splitter.Splitter
}

// NewServer constructs a Server from a given config struct.
func NewServer(config *config.Config) (*Server, error) {
	if config == nil {
		return nil, fmt.Errorf("prism: got nil config")
	}
	s := &Server{
		config:        config,
		upstreamProxy: &httputil.ReverseProxy{},
	}
	for _, splitterConf := range config.Splitters {
		splitter, err := splitter.NewSplitter(splitterConf)
		if err != nil {
			return nil, err
		}
		s.splitters = append(s.splitters, splitter)
	}
	s.monitor = monitor.NewMonitor(s)
	return s, nil
}

// Start starts listeners and blocks until cancellation or error.
func (s *Server) Start(ctx context.Context) error {
	log.Println("Starting monitor on", s.config.MonitorAddr)
	go func() {
		http.ListenAndServe(s.config.MonitorAddr, s.monitor)
	}()

	var wg sync.WaitGroup
	errs := make(chan error)
	done := make(chan struct{})

	for _, s := range s.splitters {
		wg.Add(1)
		go func(s *splitter.Splitter) {
			defer wg.Done()
			errs <- s.Start(ctx)
		}(s)
	}

	var err error
	select {
	case <-ctx.Done():
		log.Println("got cancellation signal.")
		err = ctx.Err()
		if err == context.Canceled {
			err = nil
		}
	case err = <-errs:
	case <-done:
		log.Println("all splitters exited.")
	}
	return err
}

// GetConfig returns the configuration with which the Server was created.
func (s *Server) GetConfig() config.Config {
	return *s.config
}
