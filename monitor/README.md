
# monitor
    import "github.com/tmc/prism/monitor"







## type GetConfiger
``` go
type GetConfiger interface {
    GetConfig() config.Config
}
```
GetConfiger is the interface that the monitor will consume to show configuration information.











## type Monitor
``` go
type Monitor struct {
    // contains filtered or unexported fields
}
```
Monitor is the type that provides runtime information about the running prism Server.









### func NewMonitor
``` go
func NewMonitor(server GetConfiger) *Monitor
```
NewMonitor creates a Monitor given the necessary components.




### func (\*Monitor) ServeHTTP
``` go
func (m *Monitor) ServeHTTP(rw http.ResponseWriter, r *http.Request)
```
ServeHTTP satisfies the net/http.Handler interface






