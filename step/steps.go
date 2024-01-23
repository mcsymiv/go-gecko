package step

import "github.com/mcsymiv/go-gecko/driver"

type WebStep interface {
	driver.WebDriver
	Element() driver.WebElement
	SendAndSubmit(input string) WebStep
}

type Step struct {
	driver.WebDriver
	StepElement driver.WebElement
}

func New(s driver.WebDriver) *Step {
	return &Step{
		WebDriver: s,
	}
}

func (s Step) Element() driver.WebElement {
	return s.StepElement
}

/*
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

*/
