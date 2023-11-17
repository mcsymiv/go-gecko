package driver

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RequestOptions struct {
	Method  string
	Url     string
	Payload []byte
	Headers map[string]string
}

type RequestOption func(*RequestOptions)

func Method(method string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Method = method
	}
}

func Url(url string) RequestOption {
	return func(ro *RequestOptions) {
		ro.Url = url
	}
}

func Payload(payload []byte) RequestOption {
	return func(ro *RequestOptions) {
		ro.Payload = payload
	}
}

func DefaultRequestOptions() RequestOptions {
	return RequestOptions{
		Method:  http.MethodGet,
		Url:     "http://localhost:4444",
		Payload: nil,
		Headers: map[string]string{
			"Accept": "json/application",
		},
	}
}

func MakeRequest(options ...RequestOption) ([]byte, error) {
	// Default options
	requestOptions := DefaultRequestOptions()
	log.Printf("Default req options: %+v", requestOptions)

	// Apply provided options
	for _, option := range options {
		option(&requestOptions)
	}

	log.Printf("Req options: %+v", requestOptions)

	req, err := http.NewRequest(
		requestOptions.Method,
		requestOptions.Url,
		bytes.NewBuffer(requestOptions.Payload),
	)
	if err != nil {
		log.Printf("Error creating request: %+v", err)
		return nil, err
	}

	req.Header.Add("Accept", "json/application")

	// Wrapper for RoundTripper Transport
	// Sets local logger for each request/response cycle
	c := &http.Client{
		Transport: &LoggingRoundTripper{
			next:   http.DefaultTransport,
			logger: log.New(os.Stdout, "[info]\t", log.Ldate|log.Ltime),
		},
	}
	res, err := c.Do(req)
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

// FormatActiveSessionUrl
func FormatActiveSessionUrl(url string, d *Driver) string {
	return fmt.Sprintf("%s%s/session/%s/%s",
		d.Capabilities.Host,
		d.Capabilities.Port,
		d.Session.SessionId,
		url,
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
