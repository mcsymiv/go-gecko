package driver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/selenium"
)

const (
	ById              = "id" // not speciied by w3c
	ByXPath           = "xpath"
	ByLinkText        = "link text"
	ByPartialLinkText = "partial link text"
	ByName            = "name" // not specified by w3c
	ByTagName         = "tag name"
	ByClassName       = "class name" // not specified by w3c
	ByCssSelector     = "css selector"
)

type WebElement interface {
	Id() (string, error)
	ElementIdentifier() map[string]string
	Click() error
	SendKeys(keys string) error
	Attribute(attr string) (string, error)
	Text() (string, error)
}

type WebElements interface {
	Elements() ([]WebElement, error)
}

type Element struct {
	Driver    *Driver
	SessionId string
	ElementId string
}

type Elements struct {
	Driver     *Driver
	SessionId  string
	ElementsId []string
}

type SendKeys struct {
	Text string `json:"text"`
}

// Empty
// Due to geckodriver bug: https://github.com/webdriverio/webdriverio/pull/3208
// "where Geckodriver requires POST requests to have a valid JSON body"
// Used in POST requests that don't require data to be passed by W3C
type Empty struct{}

type FindUsing struct {
	Using string `json:"using"`
	Value string `json:"value"`
}

// Id
// Returns Element w3c id
// Which is bound to the selenium.WebElementIdentifier
func (e *Element) Id() (string, error) {
	if e.ElementId == "" {
		return "", fmt.Errorf("No id for element: %+v", e)
	}
	return e.ElementId, nil
}

// ElementIdentifier
// Returns w3c element full value
// For example: {"value": {"webElementIdentifier": "element-id"}}
func (e *Element) ElementIdentifier() map[string]string {
	return map[string]string{
		selenium.WebElementIdentifier: e.ElementId,
	}
}

// Elements
// Converts returned elements ids by driver
// Into usable slice of WebElements
func (els *Elements) Elements() ([]WebElement, error) {
	var wels []WebElement

	if len(els.ElementsId) == 0 {
		return nil, fmt.Errorf("No element ids. Empty slice of web elements: %+v", els)
	}

	for _, el := range els.ElementsId {
		wels = append(wels, &Element{
			SessionId: els.SessionId,
			ElementId: el,
		})
	}

	return wels, nil
}

// FindElement
// Finds single element by specifying selector strategy and its value
// Uses Selenium 3 protocol UUID-based string constant
func (d *Driver) FindElement(by, value string) (WebElement, error) {
	url := formatActiveSessionUrl(d, "element")
	data, err := json.Marshal(&FindUsing{
		Using: by,
		Value: value,
	})
	if err != nil {
		log.Printf("Find element marshal: %+v", err)
		return nil, err
	}

	el, err := makeReq(d, WithUrl(url), WithMethod(http.MethodPost), WithPayload(data))
	if err != nil {
		log.Printf("Find element request: %+v", err)
		return nil, err
	}

	res := new(struct{ Value map[string]string })
	if err := json.Unmarshal(el, &res); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
		return nil, err
	}

	// Retrieves w3c element id
	id := selenium.ElementID(res.Value)

	return &Element{
		Driver:    d,
		SessionId: d.Session.SessionId,
		ElementId: id,
	}, nil
}

func (d *Driver) FindElements(by, value string) (WebElements, error) {
	url := formatActiveSessionUrl(d, "elements")
	data, err := json.Marshal(&FindUsing{
		Using: by,
		Value: value,
	})
	if err != nil {
		log.Printf("Find elements marshal: %+v", err)
		return nil, err
	}

	elsResponse, err := makeReq(d, WithMethod(http.MethodPost), WithUrl(url), WithPayload(data))
	if err != nil {
		log.Printf("Find elements request: %+v", err)
		return nil, err
	}

	res := new(struct{ Value []map[string]string })
	if err := json.Unmarshal(elsResponse, &res); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
		return nil, err
	}

	els := selenium.ElementsID(res.Value)
	if els == nil {
		log.Printf("No elements found. Empty slice. Elements ids: %+v", els)
	}

	return &Elements{
		Driver:     d,
		SessionId:  d.Session.SessionId,
		ElementsId: els,
	}, nil
}

func (e *Element) Click() error {
	url := formatActiveSessionUrl(e.Driver, "element", e.ElementId, "click")
	data, err := json.Marshal(&Empty{})
	if err != nil {
		log.Printf("Error on empty click marshal: %+v", err)
		return err
	}
	//rr, err := request.Do(http.MethodPost, url, data)
	rr, err := makeReq(e.Driver, WithMethod(http.MethodPost), WithUrl(url), WithPayload(data))
	if err != nil {
		log.Printf("Error on click: %+v", err)
		return err
	}

	res := new(struct{ Value map[string]string })
	err = json.Unmarshal(rr, res)
	if res.Value["error"] != "" || err != nil {
		return err
	}

	return nil
}

// Attribute
// Returns elements attribute value
func (e *Element) Attribute(a string) (string, error) {
	url := formatActiveSessionUrl(e.Driver, "element", e.ElementId, "attribute", a)
	rr, err := makeReq(e.Driver, WithMethod(http.MethodGet), WithUrl(url))
	if err != nil {
		log.Printf("Get attribute: %+v", err)
		return "", err
	}

	attr := new(struct{ Value string })
	err = json.Unmarshal(rr, attr)
	if err != nil {
		log.Printf("Marshal attribute: %+v", err)
		return "", nil
	}

	return attr.Value, nil
}

func (e *Element) SendKeys(s string) error {
	url := formatActiveSessionUrl(e.Driver, "element", e.ElementId, "value")
	k := &SendKeys{
		Text: s,
	}

	data, err := json.Marshal(k)
	if err != nil {
		log.Printf("Send keys on marshal: %+v", err)
		return err
	}
	_, err = makeReq(e.Driver, WithMethod(http.MethodPost), WithUrl(url), WithPayload(data))
	if err != nil {
		log.Printf("Click: %+v", err)
		return err
	}

	return nil
}

// Text
// Returns an element’s text “as rendered”
func (e *Element) Text() (string, error) {
	url := formatActiveSessionUrl(e.Driver, "element", e.ElementId, "text")
	r, err := makeReq(e.Driver, WithMethod(http.MethodPost), WithUrl(url))
	if err != nil {
		log.Printf("Get text: %+v", err)
		return "", err
	}

	text := new(struct{ Value string })
	err = json.Unmarshal(r, text)
	if err != nil {
		log.Printf("Marshal text: %+v", err)
		return "", nil
	}

	return text.Value, nil
}
