package element

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// Click
func (e *Element) Click() error {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Click)

	data, err := json.Marshal(&Empty{})
	if err != nil {
		log.Printf("Error on empty click marshal: %+v", err)
	}

	_, err = request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Error on click: %+v", err)
		return err
	}

	return nil
}

// SendKeys
func (e *Element) SendKeys(s string) error {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Value)
	k := &SendKeys{
		Text: s,
	}

	data, err := json.Marshal(k)
	if err != nil {
		log.Printf("Send keys on marshal: %+v", err)
		return err
	}

	_, err = request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Click: %+v", err)
		return err
	}

	return nil
}

// GetAttribute
func (e *Element) Attribute(a string) (string, error) {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Attribute, a)
	r, err := request.Do(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Get attribute: %+v", err)
		return "", err
	}

	attr := new(struct{ Value string })
	err = json.Unmarshal(r, attr)
	if err != nil {
		log.Printf("Marshal attribute: %+v", err)
		return "", nil
	}

	return attr.Value, nil
}

func (e *Element) ElementId() (string, error) {
	if e.Id == "" {
		return "", fmt.Errorf("No id for element: %+v", e)
	}

	return e.Id, nil
}

// Elements
func (els *Elements) Elements() ([]WebElement, error) {
	var wels []WebElement

	if len(els.Ids) == 0 {
		return nil, fmt.Errorf("No element ids. Empty slice of web elements: %+v", els)
	}

	for _, el := range els.Ids {
		wels = append(wels, &Element{
			SessionId: els.SessionId,
			Id:        el,
		})
	}

	return wels, nil
}
