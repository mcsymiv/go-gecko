package driver

import (
	"log"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

func TestDriver(t *testing.T) {

	d, err := driver.New(capabilities.ImplicitWait(3000))
	if err != nil {
		log.Fatal("session start err", err)
	}
	defer d.Quit()

	_, err = d.Open("https://www.google.com")
	if err != nil {
		t.Errorf("Url: %+v", err)
	}

	el, err := d.FindElement(element.ByCssSelector, "#APjFqb")
	if err != nil {
		t.Errorf("Element not found: %+v", err)
	}

	el.Click()

	attr, err := el.Attribute("id")

	if err != nil && attr == "" {
		t.Errorf("No attribute: %+v", err)
	}

	el.SendKeys("hello")
	el.SendKeys(string("\ue007"))
}
