package splitter

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"sync"

	"golang.org/x/net/context"

	"github.com/tmc/prism/config"
	"github.com/tmc/prism/monitor"
	"github.com/tmc/prism/splitter"
)

type Server struct {
	config *config.Config

	upstreamProxy *httputil.ReverseProxy
	monitor       *monitor.Monitor
	splitters     []*splitter.Splitter
}

func NewServer(config *config.Config) (*Server, error) {
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

func (s *Server) Start() error {
	log.Println("Starting monitor on", s.config.MonitorAddr)
	go func() {
		http.ListenAndServe(s.config.MonitorAddr, s.monitor)
	}()

	ctx, cancelFn := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	errs := make(chan error)
	for _, s := range s.splitters {
		wg.Add(1)
		go func(s *splitter.Splitter) {
			defer wg.Done()
			errs <- s.Start(ctx)
		}(s)
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		// Block until a signal is received.
		s := <-c
		log.Println("Got signal:", s)
		cancelFn()
	}()
	defer cancelFn()
	<-ctx.Done()
	err := ctx.Err()
	if err == context.Canceled {
		return nil
	}
	return err
}

func (s *Server) GetConfig() config.Config {
	return *s.config
}

//todo remove

func (s *Server) splitterStats() map[string]interface{} {
	return map[string]interface{}{
		"implemented": false,
	}
}
