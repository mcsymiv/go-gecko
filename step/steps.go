package step

import (
	"log"
	"testing"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
)

type WebStep interface {
	session.WebDriver
  Element() element.WebElement
  Elements() element.WebElements
  SendAndSubmit(input string) WebStep
}

type Step struct {
	session.WebDriver
  StepElement element.WebElement
  StepElements element.WebElements
}

func New(s session.WebDriver) *Step {
	return &Step{
		WebDriver: s,
	}
}

func (s Step) Element() element.WebElement {
  return s.StepElement
}

func (s Step) Elements() element.WebElements {
  return s.StepElements
}

// FindAndClick
// Convenience method
// Wraps FindElement and Click actions
func (s Step) FindAndClick(by, val string, t *testing.T) WebStep {
	el, err := s.WebDriver.FindElement(by, val)
	if err != nil {
		t.Errorf("Error find element: %+v", err)
	}

	el.Click()

	return s
}

func (s Step) FindX(val string) WebStep {
  el, err := s.WebDriver.FindElement(element.ByXPath, val)
  if err != nil {
    log.Println("Unable to find element by xpath", err)
  }

  newStep := Step{
    StepElement: el,
  }

  return newStep
}

func (s Step) FindAllCss(val string) WebStep {
  el, err := s.WebDriver.FindElements(element.ByCssSelector, val)
  if err != nil {
    log.Println("Unable to find element by css", err)
  }

  newStep := Step{
    StepElements: el,
  }

  return newStep
}


func (s Step) FindCss(val string) WebStep {
  el, err := s.WebDriver.FindElement(element.ByCssSelector, val)
  if err != nil {
    log.Println("Unable to find element by css", err)
  }

  newStep := Step{
    StepElement: el,
  }

  return newStep
}

func (s Step) SendAndSubmit(input string) WebStep {
  el := s.Element()

  el.SendKeys(input)
  el.SendKeys(element.EnterKey)

  newStep := Step{
    StepElement: el,
  }
  return newStep
}
