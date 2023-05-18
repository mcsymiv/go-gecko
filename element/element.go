package element

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// Click
func (e *Element) Click() {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Click)
	_, err := request.Do(http.MethodPost, url, nil)
	if err != nil {
		fmt.Println("Error on click", err)
	}
}

// SendKeys
func (e *Element) SendKeys(s string) {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Value)
	k := &SendKeys{
		Text: s,
	}

	data, err := json.Marshal(k)
	fmt.Println(string(data))
	if err != nil {
		fmt.Println("Send keys on marshal error", err)
	}

	_, err = request.Do(http.MethodPost, url, data)
	if err != nil {
		fmt.Println("Error on click", err)
	}
}

// GetAttribute
func (e *Element) GetAttribute(a string) {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Attribute, a)
	r, err := request.Do(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Get attribute error", err)
	}

	fmt.Println(string(r))

	// d, err := json.Marshal(r)
}
