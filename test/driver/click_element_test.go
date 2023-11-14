package driver

import (
	"log"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
)

func TestClick(t *testing.T) {
	d, err := session.New(capabilities.ImplicitWait(3000))
	if err != nil {
		log.Fatal("session start err", err)
	}
	defer d.Quit()

	d.Open("https://www.google.com")

	el, err := d.FindElement(element.ByCssSelector, "[class='FPdoLc lJ9FBc'] [class='RNmpXc']")
	if err != nil {
		t.Errorf("Element not found: %+v", err)
	}

	el.Click()

	u, _ := d.GetUrl()
	if u == "" {
		t.Errorf("Unable to get URL: %+v", err)
	}

	time.Sleep(2 * time.Second)
}
