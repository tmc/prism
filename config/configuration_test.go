package config_test

import (
	"fmt"

	"github.com/tmc/prism/config"
)

func ExampleConfig() {
	c := config.Config{
		Splitters: []config.SplitterConfig{
			{
				Label:      "main",
				ListenAddr: "localhost:7000",
				Source:     "localhost:8000",
				Sinks: []config.Sink{
					{"version a", "localhost:8001"},
					{"version b", "localhost:8002"},
				},
				// Doesn't send requests to sinks until the response from
				// source arrives.
				WaitForResponse: true,
			},
		},
		MonitorAddr: ":7070",
	}
	fmt.Println(c)
	// output:
	// {"Splitters":[{"Label":"main","ListenAddr":"localhost:7000","Source":"localhost:8000","WaitForResponse":true,"Sinks":[{"Name":"version a","Addr":"localhost:8001"},{"Name":"version b","Addr":"localhost:8002"}]}],"MonitorAddr":":7070"}
}
