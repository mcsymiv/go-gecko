package driver

import (
	"github.com/mcsymiv/go-gecko/hooks"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/element"
)

func TestSendKeys(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("https://www.google.com")

<<<<<<< Updated upstream:test/driver/send_keys_test.go
	el, err := d.FindElement(element.ByCssSelector, "#APjFqb")
=======
	_, err := d.FindElement(driver.ByCssSelector, "#APjFqb")
>>>>>>> Stashed changes:test/drivertest/send_keys_test.go
	if err != nil {
		t.Errorf("Element not found: %+v", err)
	}

	//el.Click()
	//el.SendKeys("hello")

	time.Sleep(1 * time.Second)
}
