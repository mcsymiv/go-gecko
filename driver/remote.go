package driver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/request"
)

// Status
func (d *Driver) Status() {
	rr, err := request.Do(http.MethodGet, request.Url(request.Status), nil)
	if err != nil {
		fmt.Println("Status request error", err)
	}

	reply := &StatusResponse{}
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

	_, err = request.Do(http.MethodPost, request.UrlArgs(request.Session, d.Id, request.UrlPath), param)
	if err != nil {
		fmt.Println("Open url POST Error", err)
	}

	rr, err := request.Do(http.MethodGet, request.UrlArgs(request.Session, d.Id, request.UrlPath), param)
	if err != nil {
		fmt.Println("Open url GET Error", err)
	}

	fmt.Println(string(rr))
}

// Closes session
func (d *Driver) Quit() {
	request.Do(http.MethodDelete, request.UrlArgs(request.Session, d.Id), nil)
}

//- func elementIDFromValue(v map[string]string) string {
//-	for _, key := range []string{webElementIdentifier, legacyWebElementIdentifier} {
//-		v, ok := v[key]
//-		if !ok || v == "" {
//-			continue
//-		}
//-		return v
//-	}
//-	return ""
//- }

func (d *Driver) FindElement(by, value string) element.WebElement {
	// element.Find(element.ByCSSSelector, "#APjFqb", d.Id)
	params := map[string]string{
		"using": by,
		"value": value,
	}

	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Find element error marshal", err)
	}

	url := request.UrlArgs(request.Session, d.Id, request.Element)
	_, err = request.Do(http.MethodPost, url, data)
	if err != nil {
		fmt.Println("Find element request error", err)
	}

	res := new(struct{ Value map[string]string })
	if err := json.Unmarshal(data, &res); err != nil {
		fmt.Println("Find element unmarshal error", err)
	}

	return &element.Element{}
}
