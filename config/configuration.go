package config

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Config struct {
	// Configuration of the splitters
	Splitters []SplitterConfig
	// address to set up HTTP monitoring interface
	MonitorAddr string
}

// SplitterConfig describes an instance of a splitter.
type SplitterConfig struct {
	Label      string
	ListenAddr string // listen Addr
	Source     string // upstream source

	WaitForResponse    bool // if true, the requests are not relayed to sinks until the upstream response has been read.
	Sinks              []Sink
	SinkRequestTimeout time.Duration `json:",omitempty"` // timeout on request to sinks
	RequestBufferSize  int           `json:",omitempty"` // size of request queue
}

type MatcherConfig struct {
	Type    string
	Content string
}

type Sink struct {
	Name string
	Addr string
}

func (c Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("err encoding %#v, %s", c, err)
	}
	return string(b)

}

func NewConfig(conf io.Reader) (*Config, error) {
	c := &Config{}
	return c, json.NewDecoder(conf).Decode(&c)
}
