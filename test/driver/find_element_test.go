package driver

import (
	"github.com/mcsymiv/go-gecko/hooks"
	"testing"

	"github.com/mcsymiv/go-gecko/element"
)

func TestFindElement(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	el, err := d.FindElement(element.ByLinkText, "A/B Testing")
	if err != nil {
		t.Errorf("Unable to find element: %+v", err)
	}

	id, err := el.ElementId()
	if err != nil || id == "" {
		t.Errorf("No element found: %s", id)
	}
}

func TestFindElements(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	els, err := d.FindElements(element.ByCssSelector, "#content li a")
	if err != nil {
		t.Errorf("Unable to find element: %+v", err)
	}

	wels, _ := els.Elements()
	if len(wels) == 0 {
		t.Errorf("No elements found: %+v", wels)
	}
}
