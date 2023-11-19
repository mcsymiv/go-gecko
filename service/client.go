package service

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type WebClient struct {
	Client         *http.Client
	RequestOptions *RequestOptions
}

type RequestOptions struct {
	Method  string
	Url     string
	Payload []byte
	Headers map[string]string
}

type RequestOptionFunc func(*RequestOptions)

func (r *RequestOptions) WithMethod(method string) RequestOptionFunc {
	return func(ro *RequestOptions) {
		ro.Method = method
	}
}

func (r *RequestOptions) WithUrl(url string) RequestOptionFunc {
	return func(ro *RequestOptions) {
		ro.Url = url
	}
}

func (r *RequestOptions) WithPayload(payload []byte) RequestOptionFunc {
	return func(ro *RequestOptions) {
		ro.Payload = payload
	}
}

// NewClient
// Request option to set new http client
// Can be used to perform API calls in test
func NewClient() *WebClient {
	return &WebClient{
		Client: &http.Client{
			Transport: &LoggingRoundTripper{
				next:   http.DefaultTransport,
				logger: log.New(os.Stdout, "[info]\t", log.Ldate|log.Ltime),
			},
		},
		RequestOptions: DefaultRequestOptions(),
	}
}

func DefaultRequestOptions() *RequestOptions {
	return &RequestOptions{
		Method:  http.MethodGet,
		Url:     "http://localhost:4444",
		Payload: nil,
		Headers: map[string]string{
			"Accept": "json/application",
		},
	}
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
	l.logger.Printf(`
		\n
		-------
		%s Request: %s`,
		r.Method, r.URL,
	)
	if r.Body != nil {
		l.logger.Printf("Request data: %s", r.Body)
	}
	return l.next.RoundTrip(r)
}

func (w *WebClient) MakeRequest(opts ...RequestOptionFunc) ([]byte, error) {
	// Default options set if none provided
	if w.RequestOptions == nil {
		w.RequestOptions = DefaultRequestOptions()
		return nil, errors.New("error driver Request options is not set")
	}

	// Clear payload from prev request
	w.RequestOptions.Payload = nil

	// Apply provided options
	for _, option := range opts {
		option(w.RequestOptions)
	}

	req, err := http.NewRequest(
		w.RequestOptions.Method,
		w.RequestOptions.Url,
		bytes.NewBuffer(w.RequestOptions.Payload),
	)
	if err != nil {
		log.Printf("Error creating request: %+v", err)
		return nil, err
	}

	// Apply headers from req options
	for k, v := range w.RequestOptions.Headers {
		req.Header.Add(k, v)
	}

	// Uses Client with RoundTripper Transport Wrapper
	// With Logger
	res, err := w.Client.Do(req)
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
