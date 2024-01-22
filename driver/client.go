package driver

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type RequestOptions struct {
	Client  *http.Client
	Method  string
	Url     string
	Payload []byte
	Headers map[string]string
	Port    int
}

type RequestOptionFunc func(*RequestOptions)

func WithMethod(method string) RequestOptionFunc {
	return func(ro *RequestOptions) {
		ro.Method = method
	}
}

func WithUrl(url string) RequestOptionFunc {
	return func(ro *RequestOptions) {
		ro.Url = url
	}
}

func WithPayload(payload []byte) RequestOptionFunc {
	return func(ro *RequestOptions) {
		ro.Payload = payload
	}
}

// Client
// Request option to set new http client
// Can be used to perform API calls in test
func Client(client *http.Client) RequestOptionFunc {
	return func(ro *RequestOptions) {
		ro.Client = client
	}
}

func DefaultRequestOptions() RequestOptions {
	return RequestOptions{
		Client: &http.Client{
			Transport: &LoggingRoundTripper{
				next:   http.DefaultTransport,
				logger: log.New(os.Stdout, "[info]\t", log.Ldate|log.Ltime),
			},
		},
		Method:  http.MethodGet,
		Url:     "http://localhost:4444",
		Payload: nil,
		Headers: map[string]string{
			"Accept": "json/application",
		},
	}
}

func makeReq(d *Driver, opts ...RequestOptionFunc) ([]byte, error) {
	// Default options set if none provided
	if d.RequestOptions == nil {
		ro := DefaultRequestOptions()
		d.RequestOptions = &ro
		return nil, errors.New("error driver Request options is not set")
	}

	// Clear payload from prev request
	d.RequestOptions.Payload = nil

	// Apply provided options
	for _, option := range opts {
		option(d.RequestOptions)
	}

	req, err := http.NewRequest(
		d.RequestOptions.Method,
		d.RequestOptions.Url,
		bytes.NewBuffer(d.RequestOptions.Payload),
	)
	if err != nil {
		log.Printf("Error creating request: %+v", err)
		return nil, err
	}

	// Apply headers from req options
	for k, v := range d.RequestOptions.Headers {
		req.Header.Add(k, v)
	}

	// Uses Client with RoundTripper Transport Wrapper
	// With Logger
	res, err := d.RequestOptions.Client.Do(req)
	if err != nil {
		log.Println("Error perform request:", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error read response body:", err)
		return nil, err
	}

	return body, nil
}

// formatActiveSessionUrl
// Return fully format driver url
// When active session is running
// TODO: add/update func to handle other driver endpoints
func formatActiveSessionUrl(d *Driver, args ...string) string {

	// 1st todo: adds check for args len,
	// if any, appends "/endpoint" like string
	// to active session
	var appendedArgs string
	if len(args) != 0 {
		appendedArgs = fmt.Sprintf("/%s", strings.Join(args, "/"))
	}
	return fmt.Sprintf("%s%s/session/%s%s",
		d.Capabilities.Host,
		d.Capabilities.Port,
		d.Session.SessionId,
		appendedArgs,
	)
}

// LoggingRoundTripper
// Wrapper for RoundTripper Transport
type LoggingRoundTripper struct {
	next   http.RoundTripper
	logger *log.Logger
}

// RoundTrip
// Local logger to output request and payload sent to the driver
func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {

	l.logger.Printf("\n")
	l.logger.Printf("-------")
	l.logger.Printf("%s Request: %s", r.Method, r.URL)
	if r.Body != nil {
		l.logger.Printf("Request data: %s", r.Body)
	}
	return l.next.RoundTrip(r)
}
