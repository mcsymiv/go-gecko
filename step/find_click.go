package step

import (
	"testing"

	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

type WebStep interface {
	driver.WebDriver
}

type Step struct {
	driver.WebDriver
}

func New(d driver.WebDriver) *Step {
	return &Step{
		WebDriver: d,
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
