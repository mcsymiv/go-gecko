package driver

import (
	"log"
)

const (
	// legacyWebElementIdentifier is the string constant used in the old Selenium 2 protocol
	// WebDriver JSON protocol that is the key for the map that contains an
	// unique element identifier.
	// This value is ignored in element id retreival
	legacyWebElementIdentifier = "ELEMENT"

	// webElementIdentifier is the string constant defined by the W3C Selenium 3 protocol
	// specification that is the key for the map that contains a unique element identifier.
	webElementIdentifier = "element-6066-11e4-a52e-4f735466cecf"
)

func elementID(v map[string]string) string {
	id, ok := v[webElementIdentifier]
	if !ok || id == "" {
		log.Println("Error on find element", v)
	}
	return id
}

func elementsID(v []map[string]string) []string {
	var els []string

	for _, el := range v {
		id, ok := el[webElementIdentifier]
		if !ok || id == "" {
			log.Println("Error on find element", v)
		}
		els = append(els, id)
	}

	return els
}
