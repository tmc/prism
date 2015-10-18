
# config
    import "github.com/tmc/prism/config"







## type Config
``` go
type Config struct {
    // Configuration of the splitters
    Splitters []SplitterConfig
    // address to set up HTTP monitoring interface
    MonitorAddr string
}
```
Config is the top level configuration description for a prism Server instance.
Each Server has one monitor address (a debug and stats collection endpoint) and
a set of Splitters.









### func NewConfig
``` go
func NewConfig(conf io.Reader) (*Config, error)
```
NewConfig constructs a Config instance from a io.Reader.




### func (Config) String
``` go
func (c Config) String() string
```
String provides a string representation of a Config.



## type Sink
``` go
type Sink struct {
    Name string
    Addr string
}
```
Sink is a downstream http server that gets copies of the incoming requests.











## type SplitterConfig
``` go
type SplitterConfig struct {
    Label      string // arbitrary label for humans
    ListenAddr string // listen Addr
    Source     string // upstream source http address

    WaitForResponse    bool // if true, the requests are not relayed to sinks until the upstream response has been read.
    Sinks              []Sink
    SinkRequestTimeout int `json:",omitempty"` // timeout on request to sinks in seconds
    RequestBufferSize  int `json:",omitempty"` // size of request queue
}
```
SplitterConfig describes an instance of a splitter.














