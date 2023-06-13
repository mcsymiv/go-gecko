package driver

import (
	"log"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
)

func TestFindElement(t *testing.T) {

	caps := capabilities.ImplicitWait(3000)
	d, err := session.New(caps)
	if err != nil {
		log.Fatal("Create session session", err)
	}

	defer d.Quit()

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

	caps := capabilities.ImplicitWait(3000)
	d, err := session.New(caps)
	if err != nil {
		log.Fatal("Create session session", err)
	}

	defer d.Quit()

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
