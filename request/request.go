package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const JsonContentType = "application/json"

// gecko default port
// const BaseUrl = "http://localhost:4444"

// chrome default port
const BaseUrl = "http://localhost:9515"

const (
	Session     = "session"
	Status      = "status"
	UrlPath     = "url"
	Element     = "element"
	Elements    = "elements"
	Click       = "click"
	Value       = "value"
	Attribute   = "attribute"
	Text        = "text"
	PageSource  = "source"
	Execute     = "execute"
	ScriptSync  = "sync"
	SwitchFrame = "frame"
)

func Url(arg string) string {
	return fmt.Sprintf("%s/%s", BaseUrl, arg)
}

func UrlArgs(args ...string) string {
	return fmt.Sprintf("%s/%s", BaseUrl, strings.Join(args, "/"))
}

type LoggingRoundTripper struct {
	next   http.RoundTripper
	logger *log.Logger
}

func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	l.logger.Printf("\n")
	l.logger.Printf("---------------------------------------------------------------------------")
	l.logger.Printf("%s Request: %s", r.Method, r.URL)
	if r.Body != nil {
		l.logger.Printf("Request data: %s", r.Body)
	}
	return l.next.RoundTrip(r)
}

// Do
// Perform http.Client request to the driver
// Prints req, res values to the stdout
func Do(method, url string, data []byte) (json.RawMessage, error) {
	req, err := New(strings.ToUpper(method), url, data)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Transport: &LoggingRoundTripper{
			next:   http.DefaultTransport,
			logger: log.New(os.Stdout, "[info]\t", log.Ldate|log.Ltime),
		},
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Response: %+v", string(body))
	return body, nil
}

// New NewRequest creates and returns http.Request
// Separates request logic into func as convenience method
func New(method, url string, data []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", JsonContentType)

	return request, nil
}

// Requester
// Implements strategy pattern
// For some of the repeated request cases:
//  1. Create url endpoint
//  2. Call request.Do(METHOD, URL, DATA)
//  3. Unmarshal response to new(struct{ Value string})
type Requester interface {
	Url() string
}

// ContextRequester
type ContextRequester struct {
	Requester
}

// NewRequester
// Creates new ContextRequester
// Accepts struct that implements Requester interface
func NewRequester(r Requester) *ContextRequester {
	return &ContextRequester{
		Requester: r,
	}
}

// SetRequester
// Setter mentod for new ContextRequester
func (ctx *ContextRequester) SetRequester(r Requester) *ContextRequester {
	ctx.Requester = r
	return ctx
}

// Get
// Wraps GET method to the driver
// And performs Default unmarshal for the response
// Response example: { "value": string }
//
// 1. request.Do(GET, url, nil)
// 2. unmarshal response
// 3. return Value string
func (ctx *ContextRequester) Get() string {

	// Url context wrapper for the path.UrlsArgs method
	url := ctx.Url()

	// Default GET method to the gecko driver
	res, err := Do(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error on request: %+v", err)
		return ""
	}

	// Unmarshal driver response for GET methods
	val := new(struct{ Value string })
	err = json.Unmarshal(res, val)
	if err != nil {
		log.Printf("Error on unmarshal response: %+v", err)
		return ""
	}

	return val.Value
}

// Post PostDefault
// Performs default POST method without returned value
// Or returned value from driver is null: { "value": null }
//
// 1. marshal data
// 2. request.Do(POST, url, data)
// 3. return []byte response for client to handle umarshal
func (ctx *ContextRequester) Post(d interface{}) []byte {

	url := ctx.Url()

	data, err := json.Marshal(d)
	if err != nil {
		log.Printf("Error marshal: %+v", err)
	}

	res, err := Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Error request: %+v", err)
	}

	// Response raw
	// Checks for { "value" : { "error": ... }}
	// Response from driver
	rr := new(struct{ Value map[string]interface{} })
	err = json.Unmarshal(res, rr)
	if rr.Value["error"] != nil {
		log.Printf("ERROR: %+v", rr)
	}

	return res
}
