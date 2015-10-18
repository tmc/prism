package config

import (
	"encoding/json"
	"fmt"
	"io"
)

// Config is the top level configuration description for a prism Server instance.
// Each Server has one monitor address (a debug and stats collection endpoint) and
// a set of Splitters.
type Config struct {
	// Configuration of the splitters
	Splitters []SplitterConfig
	// address to set up HTTP monitoring interface
	MonitorAddr string
}

// SplitterConfig describes an instance of a splitter.
type SplitterConfig struct {
	Label      string // arbitrary label for humans
	ListenAddr string // listen Addr
	Source     string // upstream source http address

	WaitForResponse    bool // if true, the requests are not relayed to sinks until the upstream response has been read.
	Sinks              []Sink
	SinkRequestTimeout int `json:",omitempty"` // timeout on request to sinks in seconds
	RequestBufferSize  int `json:",omitempty"` // size of request queue
}

// Sink is a downstream http server that gets copies of the incoming requests.
type Sink struct {
	Name string
	Addr string
}

// String provides a string representation of a Config.
func (c Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("err encoding %#v, %s", c, err)
	}
	return string(b)

}

// NewConfig constructs a Config instance from a io.Reader.
func NewConfig(conf io.Reader) (*Config, error) {
	c := &Config{}
	return c, json.NewDecoder(conf).Decode(&c)
}
