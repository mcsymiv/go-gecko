package drivertest

import (
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/hooks"
	"testing"
)

func TestFindElement(t *testing.T) {

	d, tear := hooks.Driver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	el, err := d.FindElement(driver.ByLinkText, "A/B Testing")
	if err != nil {
		t.Errorf("Unable to find element: %+v", err)
	}

	id, err := el.Id()
	if err != nil || id == "" {
		t.Errorf("No element found: %s", id)
	}
}

func TestFindElements(t *testing.T) {

	d, tear := hooks.Driver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	els, err := d.FindElements(driver.ByCssSelector, "#content li a")
	if err != nil {
		t.Errorf("Unable to find element: %+v", err)
	}

	wels, _ := els.Elements()
	if len(wels) == 0 {
		t.Errorf("No elements found: %+v", wels)
	}
}
