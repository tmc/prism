
# httputils
    import "github.com/tmc/prism/httputils"






## func CopyHeader
``` go
func CopyHeader(h http.Header) http.Header
```
CopyHeader copies an http.Header.


## func CopyRequest
``` go
func CopyRequest(r *http.Request) (*http.Request, error)
```
CopyRequest copies an http.Request in a shallow fashion. All fields are shared except for Body
which is read fully from the request and Header which is copied.


## func FreezeBody
``` go
func FreezeBody(r *http.Request) (io.ReadCloser, error)
```
FreezeBody reads an http.Request's Body and replaces it iwth a bytes.Buffer.



## type RequestResponse
``` go
type RequestResponse struct {
    Request   *http.Request
    Response  Response
    StartedAt time.Time
    EndedAt   time.Time
    // contains filtered or unexported fields
}
```
RequestResponse is a Request-Response pair









### func NewRequestResponse
``` go
func NewRequestResponse(req *http.Request, resp Response) *RequestResponse
```
NewRequestResponse creates a new RequestResponse. Sets StartedAt to time.Now()




### func (\*RequestResponse) Done
``` go
func (r *RequestResponse) Done() <-chan struct{}
```
Done returns a channel that blocks until the RequestResponse is marked as done.



### func (\*RequestResponse) MarkDone
``` go
func (r *RequestResponse) MarkDone()
```
MarkDone marks a RequestResponse as completed as in the Response has been generated.



## type Response
``` go
type Response interface {
    ResponseBody() io.ReadCloser
    StatusCode() int
    Header() http.Header
}
```
Response is an HTTP response.









### func NewRawResponse
``` go
func NewRawResponse(r *http.Response) (Response, error)
```
NewRawResponse consumes an http.Response by reading and closing it's body and replacing it with a byte.Buffer.




## type TeeResponseWriter
``` go
type TeeResponseWriter struct {
    http.ResponseWriter
    Body    *bytes.Buffer
    Code    int
    Flushed bool
}
```
TeeResponseWriter is an http.ResponseWriter that saves the reponse and status code for later inspection.









### func NewTeeResponseWriter
``` go
func NewTeeResponseWriter(w http.ResponseWriter) *TeeResponseWriter
```
NewTeeResponseWriter creates a TeeResponseWriter from an http.ResponseWriter




### func (\*TeeResponseWriter) Flush
``` go
func (w *TeeResponseWriter) Flush()
```
Flush flushes the underlying ResponseWriter if it is an io.Flusher and saves that it has been flushed.



### func (\*TeeResponseWriter) ResponseBody
``` go
func (w *TeeResponseWriter) ResponseBody() io.ReadCloser
```
ResponseBody returns the saved response body.



### func (\*TeeResponseWriter) StatusCode
``` go
func (w *TeeResponseWriter) StatusCode() int
```
StatusCode returns the saved status code.



### func (\*TeeResponseWriter) Write
``` go
func (w *TeeResponseWriter) Write(b []byte) (int, error)
```
Write writes the provided byte slice to the underlying ResponseWriter as well as the the internal buffer.



### func (\*TeeResponseWriter) WriteHeader
``` go
func (w *TeeResponseWriter) WriteHeader(code int)
```
WriteHeader writes the provided status code to the underlying ResponseWriter and saves it






