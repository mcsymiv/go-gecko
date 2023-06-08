package e2e

import (
	"log"

	"github.com/mcsymiv/go-gecko/element"
)

type ABPage struct {
	title element.WebElement
}

func (ab *ABPage) Title() string {
	attr, err := ab.title.Attribute("class")
	if err != nil {
		log.Println("failed attr", err)
	}

	return attr
}
