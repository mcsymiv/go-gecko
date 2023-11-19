package driver

import (
<<<<<<< Updated upstream:test/driver/click_element_test.go
=======
	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
>>>>>>> Stashed changes:test/drivertest/click_element_test.go
	"github.com/mcsymiv/go-gecko/hooks"
	"log"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/element"
)

func TestClick(t *testing.T) {
<<<<<<< Updated upstream:test/driver/click_element_test.go
	d, tear := hooks.StartDriver()
=======
	d, tear := hooks.Driver(
		capabilities.ImplicitWait(3000),
		capabilities.Port(":9515"),
		capabilities.BrowserName("chrome"),
	)
>>>>>>> Stashed changes:test/drivertest/click_element_test.go
	defer tear()

	d.Open("https://www.google.com")

	el, err := d.FindElement(element.ByCssSelector, "[class='FPdoLc lJ9FBc'] [class='RNmpXc']")
	if err != nil {
		t.Errorf("Element not found: %+v", err)
	}

	log.Printf("el: %+v", el)

	//el.Click()

	//u, _ := d.GetUrl()
	//if u == "" {
	//	t.Errorf("Unable to get URL: %+v", err)
	//}

	time.Sleep(2 * time.Second)
}
