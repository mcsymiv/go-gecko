package driver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// GetStatus
func GetStatus() {
	rr, err := request.Do(http.MethodGet, path.Url(path.Status), nil)
	if err != nil {
		fmt.Println("Status request error", err)
	}

	reply := new(struct{ Value Status })
	if err := json.Unmarshal(rr, reply); err != nil {
		fmt.Println("Status unmarshal error", err)
	}

	fmt.Println(reply)
}

// Open
// Goes to url
func (d *Driver) Open(u string) {
	url := map[string]string{
		"url": u,
	}
	param, err := json.Marshal(url)
	if err != nil {
		fmt.Println("Url marshal error", err)
	}

	_, err = request.Do(http.MethodPost, path.UrlArgs(path.Session, d.Id, path.UrlPath), param)
	if err != nil {
		fmt.Println("Open url POST Error", err)
	}

	rr, err := request.Do(http.MethodGet, path.UrlArgs(path.Session, d.Id, path.UrlPath), param)
	if err != nil {
		fmt.Println("Open url GET Error", err)
	}

	fmt.Println(string(rr))
}

// Closes session
func (d *Driver) Quit() {
	request.Do(http.MethodDelete, path.UrlArgs(path.Session, d.Id), nil)
}

// FindElement
// Finds single element by specifying selector strategy and its value
// Uses Selenium 3 protocol UUID-based string constant
func (d *Driver) FindElement(by, value string) element.WebElement {
	p := &element.FindUsing{
		Using: by,
		Value: value,
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Find element error marshal", err)
	}

	url := path.UrlArgs(path.Session, d.Id, path.Element)
	el, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		fmt.Println("Find element request error", err)
	}

	res := new(struct{ Value map[string]string })
	if err := json.Unmarshal(el, &res); err != nil {
		fmt.Println("Find element unmarshal error", err)
	}

	id := elementID(res.Value)

	return &element.Element{
		SessionId: d.Id,
		Id:        id,
	}
}
