package driver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// GetStatus
//
// Status returns information about whether a remote end is
// in a state in which it can create new sessions,
// but may additionally include arbitrary meta information
// that is specific to the implementation.
func GetStatus() (*Status, error) {
	rr, err := request.Do(http.MethodGet, path.Url(path.Status), nil)
	if err != nil {
		log.Println("Status request error", err)
		return nil, err
	}

	reply := new(struct{ Value Status })
	if err := json.Unmarshal(rr, reply); err != nil {
		log.Println("Status unmarshal error", err)
		return &reply.Value, err
	}

	return &reply.Value, nil
}

// Open
// Goes to url
func (d *Driver) Open(u string) (string, error) {
	url := map[string]string{
		"url": u,
	}
	param, err := json.Marshal(url)
	if err != nil {
		log.Println("Url marshal error", err)
		return "", err
	}

	// POST /url method returns null as value
	_, err = request.Do(http.MethodPost, path.UrlArgs(path.Session, d.Id, path.UrlPath), param)
	if err != nil {
		log.Println("Open url POST Error", err)
		return "", err
	}

	r, err := request.Do(http.MethodGet, path.UrlArgs(path.Session, d.Id, path.UrlPath), param)
	if err != nil {
		log.Printf("Open url GET error: %+v", err)
		return "", err
	}

	ur := new(struct{ Value string })
	err = json.Unmarshal(r, ur)
	if err != nil {
		log.Printf("Url GET error: %+v", err)
		return "", err
	}

	return ur.Value, nil
}

// GetUrl
func (d *Driver) GetUrl() (string, error) {

	r, err := request.Do(http.MethodGet, path.UrlArgs(path.Session, d.Id, path.UrlPath), nil)
	if err != nil {
		log.Printf("Open url GET error: %+v", err)
		return "", err
	}

	ur := new(struct{ Value string })
	err = json.Unmarshal(r, ur)
	if err != nil {
		log.Printf("Url GET error: %+v", err)
		return "", err
	}

	return ur.Value, nil
}

// Closes session
func (d *Driver) Quit() {
	request.Do(http.MethodDelete, path.UrlArgs(path.Session, d.Id), nil)
}

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

	id := elementID(res.Value)

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

	els := elementsID(res.Value)
	if els == nil {
		log.Printf("No elements found. Empty slice. Elements ids: %+v", els)
	}

	return &element.Elements{
		SessionId: d.Id,
		Ids:       els,
	}, nil
}
