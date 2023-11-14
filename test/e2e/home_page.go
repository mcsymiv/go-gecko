package e2e

import (
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
)

type HomePage struct {
	session session.WebDriver

	ab     element.WebElement
	ABPage ABPage
}

type ABPage struct {
	title element.WebElement
}

func NewHomePage(s session.WebDriver) *HomePage {
	s.Open("https://the-internet.herokuapp.com/")

	return &HomePage{
		session: s,
		ab:      s.Init(element.ByCssSelector, "#content li a"),
	}
}

func (h *HomePage) ClickOnAbTestingLink() *ABPage {
	h.ab.Click()
	return &ABPage{
		title: h.session.Init(element.ByCssSelector, "#content h3"),
	}
}
