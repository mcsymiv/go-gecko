package drivertest

import (
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/hooks"
	"testing"
	"time"
)

func TestSendKeys(t *testing.T) {

	d, tear := hooks.Driver()
	defer tear()

	d.Open("https://www.google.com")

	el, err := d.FindElement(driver.ByCssSelector, "#APjFqb")
	if err != nil {
		t.Errorf("Element not found: %+v", err)
	}

	el.Click()

	el.SendKeys("hello")

	time.Sleep(1 * time.Second)
}
