package driver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/request"
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
	ElementId() (string, error)
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
	Driver    WebDriver
	SessionId string
	Id        string
}

type Elements struct {
	Driver    WebDriver
	SessionId string
	Ids       []string
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

// ElementId
// Returns Element w3c id
// Which is bound to the selenium.WebElementIdentifier
func (e *Element) ElementId() (string, error) {
	if e.Id == "" {
		return "", fmt.Errorf("No id for element: %+v", e)
	}
	return e.Id, nil
}

// ElementIdentifier
// Returns w3c element full value
// For example: {"value": {"webElementIdentifier": "element-id"}}
func (e *Element) ElementIdentifier() map[string]string {
	return map[string]string{
		selenium.WebElementIdentifier: e.Id,
	}
}

// Elements
// Converts returned elements ids by driver
// Into usable slice of WebElements
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

// FindElement
// Finds single element by specifying selector strategy and its value
// Uses Selenium 3 protocol UUID-based string constant
func (d *Driver) FindElement(by, value string) (WebElement, error) {
	data, err := json.Marshal(&FindUsing{
		Using: by,
		Value: value,
	})
	if err != nil {
		log.Printf("Find element marshal: %+v", err)
		return nil, err
	}

	url := request.UrlArgs(request.Session, d.Session.SessionId, request.Element)
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

	return &Element{
		SessionId: d.Session.SessionId,
		Id:        id,
	}, nil
}

func (d *Driver) FindElements(by, value string) (WebElements, error) {
	data, err := json.Marshal(&FindUsing{
		Using: by,
		Value: value,
	})
	if err != nil {
		log.Printf("Find elements marshal: %+v", err)
		return nil, err
	}

	url := request.UrlArgs(request.Session, d.Session.SessionId, request.Elements)
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

	return &Elements{
		SessionId: d.Session.SessionId,
		Ids:       els,
	}, nil
}

func (e *Element) Click() error {
	url := request.UrlArgs(request.Session, e.SessionId, request.Element, e.Id, request.Click)
	data, err := json.Marshal(&Empty{})
	if err != nil {
		log.Printf("Error on empty click marshal: %+v", err)
	}
	rr, err := request.Do(http.MethodPost, url, data)
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
	url := request.UrlArgs(request.Session, e.SessionId, request.Element, e.Id, request.Attribute, a)
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

const (
	NullKey       = string('\ue000')
	CancelKey     = string('\ue001')
	HelpKey       = string('\ue002')
	BackspaceKey  = string('\ue003')
	TabKey        = string('\ue004')
	ClearKey      = string('\ue005')
	ReturnKey     = string('\ue006')
	EnterKey      = string('\ue007')
	ShiftKey      = string('\ue008')
	ControlKey    = string('\ue009')
	AltKey        = string('\ue00a')
	PauseKey      = string('\ue00b')
	EscapeKey     = string('\ue00c')
	SpaceKey      = string('\ue00d')
	PageUpKey     = string('\ue00e')
	PageDownKey   = string('\ue00f')
	EndKey        = string('\ue010')
	HomeKey       = string('\ue011')
	LeftArrowKey  = string('\ue012')
	UpArrowKey    = string('\ue013')
	RightArrowKey = string('\ue014')
	DownArrowKey  = string('\ue015')
	InsertKey     = string('\ue016')
	DeleteKey     = string('\ue017')
	SemicolonKey  = string('\ue018')
	EqualsKey     = string('\ue019')
	Numpad0Key    = string('\ue01a')
	Numpad1Key    = string('\ue01b')
	Numpad2Key    = string('\ue01c')
	Numpad3Key    = string('\ue01d')
	Numpad4Key    = string('\ue01e')
	Numpad5Key    = string('\ue01f')
	Numpad6Key    = string('\ue020')
	Numpad7Key    = string('\ue021')
	Numpad8Key    = string('\ue022')
	Numpad9Key    = string('\ue023')
	MultiplyKey   = string('\ue024')
	AddKey        = string('\ue025')
	SeparatorKey  = string('\ue026')
	SubstractKey  = string('\ue027')
	DecimalKey    = string('\ue028')
	DivideKey     = string('\ue029')
	F1Key         = string('\ue031')
	F2Key         = string('\ue032')
	F3Key         = string('\ue033')
	F4Key         = string('\ue034')
	F5Key         = string('\ue035')
	F6Key         = string('\ue036')
	F7Key         = string('\ue037')
	F8Key         = string('\ue038')
	F9Key         = string('\ue039')
	F10Key        = string('\ue03a')
	F11Key        = string('\ue03b')
	F12Key        = string('\ue03c')
	MetaKey       = string('\ue03d')
)

func (e *Element) SendKeys(s string) error {
	url := request.UrlArgs(request.Session, e.SessionId, request.Element, e.Id, request.Value)
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

// Text
// Returns an element’s text “as rendered”
func (e *Element) Text() (string, error) {
	url := request.UrlArgs(request.Session, e.SessionId, request.Element, e.Id, request.Text)
	r, err := request.Do(http.MethodGet, url, nil)
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
