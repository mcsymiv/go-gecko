package element

import (
	"fmt"

	"github.com/mcsymiv/go-gecko/selenium"
)

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
// For example: {"value": {"webelementIdentifier": "element-id"}}
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
