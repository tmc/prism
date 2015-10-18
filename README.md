# prism
    import "github.com/tmc/prism"

[![GoDoc](https://godoc.org/github.com/tmc/prism?status.svg)](http://godoc.org/github.com/tmc/prism)

Package prism is a package that helps you write traffic splitting tools.

An HTTP splitter sits in line on an http request/response path and replays some
portion of traffic to additional backends.

A splitter is defined by one upstream "source" and will repeat some portion of
traffic to a number of "sinks".

Example configuration:


	{
	  "Splitters": [
	    {
	      "Label": "main",
	      "Source": "<a href="http://localhost:8000">http://localhost:8000</a>",
	      "ListenAddr": "localhost:7000",
	      "SinkRequestTimeout": 30,
	      "WaitForResponse": true,
	      "Sinks": [
	        {
	          "Name": "version a",
	          "Addr": "localhost:8001"
	        },
	        {
	          "Name": "version b",
	          "Addr": "localhost:8002"
	        }
	      ]
	    }
	  ],
	  "MonitorAddr": ":7070"
	}







## type Server
``` go
type Server struct {
    // contains filtered or unexported fields
}
```
Server is the primary top level type in prism. It is responsible for reverse-proxying to the
upstream and mirroring requests to the sinks.









### func NewServer
``` go
func NewServer(config *config.Config) (*Server, error)
```
NewServer constructs a Server from a given config struct.




### func (\*Server) GetConfig
``` go
func (s *Server) GetConfig() config.Config
```
GetConfig returns the configuration with which the Server was created.



### func (\*Server) Start
``` go
func (s *Server) Start(ctx context.Context) error
```
Start starts listeners and blocks until cancellation or error.



