package e2e

import (
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

type HomePage struct {
	session driver.WebDriver

	ab     element.WebElement
	ABPage ABPage
}

type ABPage struct {
	title element.WebElement
}

func NewHomePage(d driver.WebDriver) *HomePage {
	d.Open("https://the-internet.herokuapp.com/")

	return &HomePage{
		session: d,
		ab:      d.Init(element.ByCssSelector, "#content li a"),
	}
}

func (h *HomePage) ClickOnAbTestingLink() *ABPage {
	h.ab.Click()
	return &ABPage{
		title: h.session.Init(element.ByCssSelector, "#content h3"),
	}
}
