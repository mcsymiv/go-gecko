package e2e

import (
	"log"

	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

type HomePage struct {
	driver driver.WebDriver

	ab element.WebElement
}

func NewHomePage(d driver.WebDriver) *HomePage {
	_, err := d.Open("https://the-internet.herokuapp.com/")
	if err != nil {
		log.Fatal("unable to open home page")
	}

	return &HomePage{
		driver: d,
		ab:     d.Init(element.ByCssSelector, "#content li a"),
	}
}

func (h *HomePage) ClickOnAbTestingLink() *ABPage {
	h.ab.Click()
	return &ABPage{
		title: h.driver.Init(element.ByCssSelector, "#content h3"),
	}
}
