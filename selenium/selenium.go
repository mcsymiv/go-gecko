package selenium

import (
	"log"
)

const (
	// LegacyWebElementIdentifier is the string constant used in the old Selenium 2 protocol
	// WebDriver JSON protocol that is the key for the map that contains an
	// unique element identifier.
	// This value is ignored in element id retreival
	LegacyWebElementIdentifier = "ELEMENT"

	// WebElementIdentifier is the string constant defined by the W3C Selenium 3 protocol
	// specification that is the key for the map that contains a unique element identifier.
	WebElementIdentifier = "element-6066-11e4-a52e-4f735466cecf"

	// ShadowRootIdentifier A shadow root is an abstraction used to identify a shadow root when
	// it is transported via the protocol, between remote and local ends.
	ShadowRootIdentifier = "shadow-6066-11e4-a52e-4f735466cecf"
)

func ElementID(v map[string]string) string {
	id, ok := v[WebElementIdentifier]
	if !ok || id == "" {
		log.Println("Error on find element", v)
	}
	return id
}

func ElementsID(v []map[string]string) []string {
	var els []string

	for _, el := range v {
		id, ok := el[WebElementIdentifier]
		if !ok || id == "" {
			log.Println("Error on find elements", v)
		}
		els = append(els, id)
	}

	return els
}
