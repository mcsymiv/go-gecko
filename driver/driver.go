package driver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/config"
	"github.com/mcsymiv/go-gecko/request"
)

type WebDriver interface {
	Status()
	Open(u string)
	Quit()
}

type Driver struct {
	Id string
}

func (d *Driver) Status() {
	rr, err := request.Do(http.MethodGet, request.Url(request.Status), nil)
	if err != nil {
		fmt.Println("Status request error", err)
	}

	fmt.Println(string(rr))
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

func New(capsFn ...config.CapabilitiesFunc) WebDriver {

	c := config.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}
	fmt.Printf("%+v", c)

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}

	rr, err := request.Do(http.MethodPost, request.Url(request.Session), data)
	if err != nil {
		fmt.Println(err)
	}

	var res config.RemoteResponse

	err = json.Unmarshal(rr, &res)
	if err != nil {
		fmt.Println("error:", err)
	}

	return &Driver{
		Id: res.Value.SessionId,
	}
}
