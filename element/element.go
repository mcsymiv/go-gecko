package element

import (
	"fmt"

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
	SessionId string
	Id        string
}

type Elements struct {
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
func (e Element) ElementId() (string, error) {
	if e.Id == "" {
		return "", fmt.Errorf("No id for element: %+v", e)
	}
	return e.Id, nil
}

// ElementIdentifier
// Returns w3c element full value
// For example: {"value": {"webelementIdentifier": "element-id"}}
func (e Element) ElementIdentifier() map[string]string {
	return map[string]string{
		selenium.WebElementIdentifier: e.Id,
	}
}

// Elements
// Converts returned elements ids by driver
// Into usable slice of WebElements
func (els Elements) Elements() ([]WebElement, error) {
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
