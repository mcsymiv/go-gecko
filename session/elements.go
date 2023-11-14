package session

import (
	"encoding/json"
	"github.com/mcsymiv/go-gecko/request"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/selenium"
)

// FindElement
// Finds single element by specifying selector strategy and its value
// Uses Selenium 3 protocol UUID-based string constant
func (s *Session) FindElement(by, value string) (element.WebElement, error) {
	data, err := json.Marshal(&element.FindUsing{
		Using: by,
		Value: value,
	})
	if err != nil {
		log.Printf("Find element marshal: %+v", err)
		return nil, err
	}

	url := path.UrlArgs(path.Session, s.Id, path.Element)
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

	// Retrieves w3c element id
	id := selenium.ElementID(res.Value)

	return &element.Element{
		SessionId: s.Id,
		Id:        id,
	}, nil
}

func (s *Session) FindElements(by, value string) (element.WebElements, error) {
	data, err := json.Marshal(&element.FindUsing{
		Using: by,
		Value: value,
	})
	if err != nil {
		log.Printf("Find elements marshal: %+v", err)
		return nil, err
	}

	url := path.UrlArgs(path.Session, s.Id, path.Elements)
	el, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Find elements request: %+v", err)
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
		SessionId: s.Id,
		Ids:       els,
	}, nil
}

// Init
func (s *Session) Init(by, val string) element.WebElement {
	el, err := s.FindElement(by, val)
	if err != nil {
		log.Println("unable to find element", err, by, val)
		return nil
	}
	return el
}
