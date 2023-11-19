package driver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// OpenStrategy
// concrete strategy for /url endpoint
type OpenStrategy struct{}

type OpenResponse struct {
	Value string
}

func (o OpenResponse) GetValue() interface{} {
	return o.Value
}

func (s *OpenStrategy) Execute(c *http.Client, req *http.Request) (WebDriverResponse, error) {
	res, err := executeRequest(c, req)

	val := &OpenResponse{}
	err = json.Unmarshal(res, val)
	if err != nil {
		log.Printf("Error on unmarshal response: %+v", err)
		return nil, err
	}

	return val, nil
}

func (d *Driver) Open(u string) (WebDriverResponse, error) {
	d.WebDriverStrategy = &OpenStrategy{}
	client := d.RequestOptions.Client
	url := formatActiveSessionUrl(d, "url")
	data, _ := json.Marshal(map[string]string{
		"url": u,
	})

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Error make request: %+v", err)
		return nil, err
	}

	// Execute the strategy
	res, err := d.WebDriverStrategy.Execute(client, req)
	if err != nil {
		log.Printf("execute open error: %+v", err)
		return nil, err
	}

	return res, nil
}
