package monitor

import (
	"encoding/json"
	"net/http"

	"github.com/tmc/prism/config"
)

type GetConfiger interface {
	GetConfig() config.Config
}

type Monitor struct {
	server GetConfiger
}

func NewMonitor(server GetConfiger) *Monitor {
	return &Monitor{server}
}

func (m *Monitor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(m.server.GetConfig())
}
