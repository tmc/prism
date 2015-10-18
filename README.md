# prism
    import "github.com/tmc/prism"

Documentation: http://godoc.org/github.com/tmc/prism


Package prism is a package that helps you write traffic splitting tools.

An HTTP splitter sits in line on an http request/response path and replays some
portion of traffic to additional backends.

A splitter is defined by one upstream "source" and will repeat some portion of
traffic to a number of "sinks".

Example configuration:


	{
	  "MonitorAddr": ":7070",
	  "Splitters": [
	    {
	      "Sinks": [
	        {
	          "Addr": "localhost:8001",
	          "Name": "version a"
	        },
	        {
	          "Addr": "localhost:8002",
	          "Name": "version b"
	        }
	      ],
	      "WaitForResponse": true,
	      "Source": "localhost:8000",
	      "Addr": "localhost:7000",
	      "Label": "main"
	    }
	  ]
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




