package step

import (
	"testing"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
)

type WebStep interface {
	session.WebDriver
}

type Step struct {
	session.WebDriver
}

func New(s session.WebDriver) *Step {
	return &Step{
		WebDriver: s,
	}
}

// FindAndClick
// Convenience method
// Wraps FindElement and Click actions
func (s *Step) FindAndClick(by, val string, t *testing.T) element.WebElement {
	el, err := s.WebDriver.FindElement(by, val)
	if err != nil {
		t.Errorf("Error find element: %+v", err)
	}

	el.Click()

	return el
}
