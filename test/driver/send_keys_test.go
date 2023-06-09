package driver

import (
	"log"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

func TestSendKeys(t *testing.T) {

	d, err := driver.New(capabilities.ImplicitWait(3000))
	if err != nil {
		log.Fatal("session start err", err)
	}
	defer d.Quit()

	d.Open("https://www.google.com")

	el, err := d.FindElement(element.ByCssSelector, "#APjFqb")
	if err != nil {
		t.Errorf("Element not found: %+v", err)
	}

	el.Click()

	el.SendKeys("hello")

	time.Sleep(1 * time.Second)
}
