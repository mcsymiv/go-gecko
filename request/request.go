package request

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const JsonContentType = "application/json"

// Do
// Performs http.Client request to the driver
// Prints req, res values to the stdout
func Do(method, url string, data []byte) (json.RawMessage, error) {
	req, err := New(strings.ToUpper(method), url, data)
	if err != nil {
		return nil, err
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("\n")
	log.Printf("---------------------------------------------------------------------------")
	log.Printf("%s Request: %s", method, url)
	if data != nil {
		log.Printf("Request data: %s", string(data))
	}
	log.Printf("-----------------------------")
	log.Printf("Response: %+v", string(body))
	log.Printf("---------------------------------------------------------------------------")

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
