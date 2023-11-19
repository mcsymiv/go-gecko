package driver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// FindStrategy
// Concrete strategy to find element
type FindStrategy struct{}

type FindResponse struct {
	Value []map[string]interface{}
	//WebElement
}

func (o FindResponse) GetElementId() interface{} {
	return o.Value
}

//func (o FindResponse) GetElement() WebElement {
//	return &Element{
//		ElementId: o.GetElementId(),
//	}
//}

func (s *FindStrategy) Execute(c *http.Client, req *http.Request) (WebElementResponse, error) {
	res, err := executeRequest(c, req)

	val := &FindResponse{}
	err = json.Unmarshal(res, val)
	if err != nil {
		log.Printf("Error on unmarshal response: %+v", err)
		return nil, err
	}

	//if err := json.Unmarshal(el, &res); err != nil {
	//	log.Printf("Find element unmarshal: %+v", err)
	//	return nil, err
	//}
	//
	//// Retrieves w3c element id
	//id := selenium.ElementID(res.Value)
	//
	//return &Element{
	//	Driver:    d,
	//	SessionId: d.Session.SessionId,
	//	ElementId: id,
	//}, nil

	return val, nil
}

// FindElement
// Finds single element by specifying selector strategy and its value
// Uses Selenium 3 protocol UUID-based string constant
func (d *Driver) FindElement(by, value string) (WebElementResponse, error) {
	url := formatActiveSessionUrl(d, "element")
	d.WebElementStrategy = &FindStrategy{}
	client := d.RequestOptions.Client

	data, err := json.Marshal(&FindUsing{
		Using: by,
		Value: value,
	})
	if err != nil {
		log.Printf("Find element marshal: %+v", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Find element request: %+v", err)
		return nil, err
	}

	res, err := d.WebElementStrategy.Execute(client, req)
	if err != nil {
		log.Printf("Find element request: %+v", err)
		return nil, err
	}

	return res, nil
}
