package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"golang.org/x/net/context"

	"github.com/tmc/prism"
	"github.com/tmc/prism/config"
)

var (
	conf = flag.String("conf", "", "path to configuration file")
)

func main() {
	flag.Parse()
	if *conf == "" {
		log.Println("required parameter -conf was not supplied.")
		os.Exit(1)
	}

	f, err := os.Open(*conf)
	if err != nil {
		log.Println("error opening configuration file:", err)
		os.Exit(1)
	}
	c, err := config.NewConfig(f)
	if err != nil {
		log.Println("error parsing configuration file:", err)
		os.Exit(1)
	}

	s, err := prism.NewServer(c)
	if err != nil {
		log.Println("error creating prism server:", err)
		os.Exit(1)
	}

	ctx, cancelFn := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		// Block until a signal is received.
		s := <-c
		log.Println("Got signal:", s)
		cancelFn()
	}()
	if err := s.Start(ctx); err != nil {
		log.Println("error running prism server:", err)
		os.Exit(1)
	}
}
