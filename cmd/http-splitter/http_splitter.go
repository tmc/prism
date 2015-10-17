package main

import (
	"flag"
	"log"
	"os"

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

	splitter, err := splitter.NewServer(c)
	if err != nil {
		log.Println("error creating splitter:", err)
		os.Exit(1)
	}

	if err := splitter.Start(); err != nil {
		log.Println("error running splitter:", err)
		os.Exit(1)
	}
}
