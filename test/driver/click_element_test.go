package driver

import (
	"log"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

func TestClick(t *testing.T) {
	d, err := driver.New(capabilities.ImplicitWait(3000))
	if err != nil {
		log.Fatal("session start err", err)
	}
	defer d.Quit()

	_, err = d.Open("https://www.google.com")
	if err != nil {
		t.Errorf("Url: %+v", err)
	}

	el, err := d.FindElement(element.ByCssSelector, "[class='FPdoLc lJ9FBc'] [class='RNmpXc']")
	if err != nil {
		t.Errorf("Element not found: %+v", err)
	}

	el.Click()

	time.Sleep(5 * time.Second)
}
