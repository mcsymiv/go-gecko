package request

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const JsonContentType = "application/json"

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

// NewRequest creates and returns http.Request
// Separetes request logic into func as convenience method
func New(method, url string, data []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", JsonContentType)

	return request, nil
}
