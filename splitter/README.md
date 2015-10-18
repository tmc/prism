
# splitter
    import "github.com/tmc/prism/splitter"







## type Splitter
``` go
type Splitter struct {
    Config config.SplitterConfig
    Client *http.Client
    // contains filtered or unexported fields
}
```
Splitter is the type that manages an upstream and a number of downstream sinks.









### func NewSplitter
``` go
func NewSplitter(config config.SplitterConfig) (*Splitter, error)
```
NewSplitter constructs a splitter described by a config.SplitterConfig




### func (\*Splitter) ServeHTTP
``` go
func (s *Splitter) ServeHTTP(rw http.ResponseWriter, r *http.Request)
```
ServeHTTP satisfies the net/http.Handler interface



### func (\*Splitter) Start
``` go
func (s *Splitter) Start(ctx context.Context) error
```
Start starts an http listener on the address in SplitterConfig.ListenAddress






