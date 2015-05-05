# har
--
    import "github.com/sebcat/har"

HAR 1.2 implementation See
[Spec](http://www.softwareishard.com/blog/har-12-spec/)

## Usage

#### type Browser

```go
type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}
```


#### type Cache

```go
type Cache struct {
	BeforeRequest *CacheRequest `json:"beforeRequest,omitempty"`
	AfterRequest  *CacheRequest `json:"afterRequest,omitempty"`
	Comment       string        `json:"comment,omitempty"`
}
```


#### type CacheRequest

```go
type CacheRequest struct {
	Expires    string `json:"expires,omitempty"`
	LastAccess string `json:"lastAccess"`
	ETag       string `json:"eTag"`
	HitCount   int    `json:"hitCount"`
	Comment    string `json:"comment,omitempty"`
}
```


#### type Content

```go
type Content struct {
	Size        int    `json:"size"`
	Compression int    `json:"compression,omitempty"`
	MIMEType    string `json:"mimeType"`
	Text        string `json:"text,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Comment     string `json:"comment,omitempty"`
}
```


#### type Cookie

```go
type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Expires  string `json:"expires,omitempty"`
	HTTPOnly bool   `json:"httpOnly,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	Comment  string `json:"comment,omitempty"`
}
```


#### type Creator

```go
type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"commment,omitempty"`
}
```


#### type Entry

```go
type Entry struct {
	Pageref         string   `json:"pageref,omitempty"`
	StartedDateTime string   `json:"startedDateTime"`
	Time            float64  `json:"time"`
	Request         Request  `json:"request"`
	Response        Response `json:"response"`
	Cache           Cache    `json:"cache"`
	Timings         Timings  `json:"timings"`
	ServerIPAddress string   `json:"serverIPAddress,omitempty"`
	Connection      string   `json:"connection,omitempty"`
	Comment         string   `json:"comment,omitempty"`
}
```


#### type HAR

```go
type HAR struct {
	Log Log `json:"log"`
}
```


#### func  Load

```go
func Load(r io.Reader) (*HAR, error)
```
Load a HAR from a reader

#### func  LoadFile

```go
func LoadFile(path string) (*HAR, error)
```
Load a HAR from a file

#### type Header

```go
type Header struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}
```


#### type Log

```go
type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Browser Browser `json:"browser,omitempty"`
	Pages   []Page  `json:"pages,omitempty"`
	Entries []Entry `json:"entries"`
	Comment string  `json:"comment,omitempty"`
}
```


#### type Page

```go
type Page struct {
	StartedDateTime string      `json:"startedDateTime"`
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	PageTimings     PageTimings `json:"pageTimings"`
	Comment         string      `json:"comment,omitempty"`
}
```


#### type PageTimings

```go
type PageTimings struct {
	OnContentLoad float64 `json:"onContentLoad,omitempty"`
	OnLoad        float64 `json:"onLoad,omitempty"`
	Comment       string  `json:"comment,omitempty"`
}
```


#### type Param

```go
type Param struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Comment     string `json:"comment,omitempty"`
}
```


#### type PostData

```go
type PostData struct {
	MIMEType string  `json:"mimeType"`
	Params   []Param `json:"params"`
	Text     string  `json:"text"`
	Comment  string  `json:"comment,omitempty"`
}
```


#### type QueryString

```go
type QueryString struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}
```


#### type Request

```go
type Request struct {
	Method string `json:"method"`
	URL    string `json:"url"`

	Cookies     []Cookie      `json:"cookies"`
	Headers     []Header      `json:"headers"`
	QueryString []QueryString `json:"queryString"`
	PostData    PostData      `json:"postData,omitempty"`
	HeadersSize int           `json:"headersSize"`
	BodySize    int           `json:"bodySize"`
	Comment     string        `json:"comment,omitempty"`
}
```


#### func (*Request) Request

```go
func (r *Request) Request() (httpreq *http.Request, err error)
```
Convert a HAR Request struct to an net/http.Request struct

#### type Response

```go
type Response struct {
	Status      int      `json:"status"`
	StatusText  string   `json:"statusText"`
	HTTPVersion string   `json:"httpVersion"`
	Cookies     []Cookie `json:"cookies"`
	Headers     []Header `json:"headers"`
	Content     Content  `json:"content"`
	RedirectURL string   `json:"redirectURL"`
	HeadersSize int      `json:"headersSize"`
	BodySize    int      `json:"bodySize"`
	Comment     string   `json:"comment,omitempty"`
}
```


#### type Timings

```go
type Timings struct {
	Blocked float64 `json:"blocked,omitempty"`
	DNS     float64 `json:"dns,omitempty"`
	Connect float64 `json:"connect,omitempty"`
	Send    float64 `json:"send"`
	Wait    float64 `json:"wait"`
	Receive float64 `json:"receive"`
	SSL     float64 `json:"ssl,omitempty"`
	Comment string  `json:"comment,omitempty"`
}
```
