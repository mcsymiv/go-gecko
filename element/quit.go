package driver

import (
	"encoding/json"
	"log"
	"net/http"
)

// QuitStrategy
// Concrete strategy for quit/close driver session endpoint
type QuitStrategy struct{}

type QuitResponse struct {
	Value string
}

func (o QuitResponse) GetValue() interface{} {
	return o.Value
}

func (s *QuitStrategy) Execute(c *http.Client, req *http.Request) (WebDriverResponse, error) {
	res, err := executeRequest(c, req)

	val := &QuitResponse{}
	err = json.Unmarshal(res, val)
	if err != nil {
		log.Printf("Error on unmarshal response: %+v", err)
		return nil, err
	}

	return val, nil
}

func (d *Driver) Quit() {
	url := formatActiveSessionUrl(d)
	d.WebDriverStrategy = &QuitStrategy{}
	client := d.RequestOptions.Client

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Printf("Error make request: %+v", err)
	}

	_, err = d.WebDriverStrategy.Execute(client, req)
	if err != nil {
		log.Printf("Error quit request: %+v", err)
	}
}
