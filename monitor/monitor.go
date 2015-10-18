package monitor

import (
	"encoding/json"
	"net/http"

	"github.com/tmc/prism/config"
)

// GetConfiger is the interface that the monitor will consume to show configuration information.
type GetConfiger interface {
	GetConfig() config.Config
}

// Monitor is the type that provides runtime information about the running prism Server.
type Monitor struct {
	server GetConfiger
}

// NewMonitor creates a Monitor given the necessary components.
func NewMonitor(server GetConfiger) *Monitor {
	return &Monitor{server}
}

// ServeHTTP satisfies the net/http.Handler interface
func (m *Monitor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(m.server.GetConfig())
}
