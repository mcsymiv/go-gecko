package driver

import (
	"encoding/json"
	"log"
	"net/http"
)

type UrlStrategy struct{}

type UrlResponse struct {
	Value string
}

func (o UrlResponse) GetValue() interface{} {
	return o.Value
}

func (s *UrlStrategy) Execute(c *http.Client, req *http.Request) (WebDriverResponse, error) {
	res, err := executeRequest(c, req)

	val := &UrlResponse{}
	err = json.Unmarshal(res, val)
	if err != nil {
		log.Printf("Error on unmarshal response: %+v", err)
		return nil, err
	}

	return val, nil
}

func (d *Driver) GetUrl() (WebDriverResponse, error) {
	d.WebDriverStrategy = &UrlStrategy{}
	client := d.RequestOptions.Client
	url := formatActiveSessionUrl(d, "url")

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
