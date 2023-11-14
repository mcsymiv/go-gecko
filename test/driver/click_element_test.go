package driver

import (
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/element"
)

func TestClick(t *testing.T) {
	d, cmd := SetupTest()
	defer cmd()

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
