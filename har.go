// HAR 1.2 implementation
// See [Spec](http://www.softwareishard.com/blog/har-12-spec/)
package har

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"commment,omitempty"`
}

type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

type PageTimings struct {
	OnContentLoad float64 `json:"onContentLoad,omitempty"`
	OnLoad        float64 `json:"onLoad,omitempty"`
	Comment       string  `json:"comment,omitempty"`
}

type Page struct {
	StartedDateTime string      `json:"startedDateTime"`
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	PageTimings     PageTimings `json:"pageTimings"`
	Comment         string      `json:"comment,omitempty"`
}

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

type Header struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

type QueryString struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

type Param struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

func (p *Param) Pair() string {
	if len(p.Value) == 0 {
		return p.Name
	} else {
		return p.Name + "=" + p.Value
	}
}

type PostData struct {
	MIMEType string  `json:"mimeType"`
	Params   []Param `json:"params"`
	Text     string  `json:"text"`
	Comment  string  `json:"comment,omitempty"`
}

func (p *PostData) Data() string {
	if len(p.Text) > 0 {
		return p.Text
	} else if len(p.Params) > 0 {
		var b bytes.Buffer
		for _, p := range p.Params {
			io.WriteString(&b, p.Pair())
		}

		return b.String()
	}

	return ""
}

type Request struct {
	Method      string        `json:"method"`
	URL         string        `json:"url"`
	httpVersion string        `json:"httpVersion"`
	Cookies     []Cookie      `json:"cookies"`
	Headers     []Header      `json:"headers"`
	QueryString []QueryString `json:"queryString"`
	PostData    PostData      `json:"postData,omitempty"`
	HeadersSize int           `json:"headersSize"`
	BodySize    int           `json:"bodySize"`
	Comment     string        `json:"comment,omitempty"`
}

// Convert a HAR Request struct to an net/http.Request struct
func (r *Request) Request() (httpreq *http.Request, err error) {

	dstr := r.PostData.Data()
	if len(dstr) > 0 {
		var body *strings.Reader
		body = strings.NewReader(dstr)
		httpreq, err = http.NewRequest(r.Method, r.URL, body)
	} else {
		httpreq, err = http.NewRequest(r.Method, r.URL, nil)
	}

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(r.Headers); i++ {
		if !strings.Contains(r.Headers[i].Name, ":") {
			httpreq.Header.Add(r.Headers[i].Name, r.Headers[i].Value)
		}
	}

	return httpreq, nil
}

type Content struct {
	Size        int    `json:"size"`
	Compression int    `json:"compression,omitempty"`
	MIMEType    string `json:"mimeType"`
	Text        string `json:"text,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

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

type CacheRequest struct {
	Expires    string `json:"expires,omitempty"`
	LastAccess string `json:"lastAccess"`
	ETag       string `json:"eTag"`
	HitCount   int    `json:"hitCount"`
	Comment    string `json:"comment,omitempty"`
}

type Cache struct {
	BeforeRequest *CacheRequest `json:"beforeRequest,omitempty"`
	AfterRequest  *CacheRequest `json:"afterRequest,omitempty"`
	Comment       string        `json:"comment,omitempty"`
}

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

type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Browser Browser `json:"browser,omitempty"`
	Pages   []Page  `json:"pages,omitempty"`
	Entries []Entry `json:"entries"`
	Comment string  `json:"comment,omitempty"`
}

type HAR struct {
	Log Log `json:"log"`
}

// Load a HAR from a reader
func Load(r io.Reader) (*HAR, error) {
	dec := json.NewDecoder(r)
	x := &HAR{}
	if err := dec.Decode(&x); err != nil {
		return nil, err
	}

	return x, nil
}

// Load a HAR from a file
func LoadFile(path string) (*HAR, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	har, err := Load(fh)
	fh.Close()
	return har, err
}
