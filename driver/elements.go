package driver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
	"github.com/mcsymiv/go-gecko/selenium"
)

// FindElement
// Finds single element by specifying selector strategy and its value
// Uses Selenium 3 protocol UUID-based string constant
func (d *Driver) FindElement(by, value string) (element.WebElement, error) {
	p := &element.FindUsing{
		Using: by,
		Value: value,
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Printf("Find element marshal: %+v", err)
		return nil, err
	}

	url := path.UrlArgs(path.Session, d.Id, path.Element)
	el, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Find element request: %+v", err)
		return nil, err
	}

	res := new(struct{ Value map[string]string })
	if err := json.Unmarshal(el, &res); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
		return nil, err
	}

	id := selenium.ElementID(res.Value)

	return &element.Element{
		SessionId: d.Id,
		Id:        id,
	}, nil
}

func (d *Driver) FindElements(by, value string) (element.WebElements, error) {
	p := &element.FindUsing{
		Using: by,
		Value: value,
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Printf("Find element marshal: %+v", err)
		return nil, err
	}

	url := path.UrlArgs(path.Session, d.Id, path.Elements)
	el, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Find element request: %+v", err)
		return nil, err
	}

	res := new(struct{ Value []map[string]string })
	if err := json.Unmarshal(el, &res); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
		return nil, err
	}

	els := selenium.ElementsID(res.Value)
	if els == nil {
		log.Printf("No elements found. Empty slice. Elements ids: %+v", els)
	}

	return &element.Elements{
		SessionId: d.Id,
		Ids:       els,
	}, nil
}
