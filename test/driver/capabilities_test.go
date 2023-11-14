package driver

import (
	"log"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
)

func TestCaps(t *testing.T) {

	moz := &capabilities.MozOptions{
		Args: []string{"--profile", "/Users/mcs/Development/tools/selenium_profile"},
	}

	d, err := session.New(
		capabilities.ImplicitWait(3000),
		capabilities.Firefox(moz),
	)
	if err != nil {
		log.Fatal("session start err", err)
	}
	defer d.Quit()

	d.Open("https://the-internet.herokuapp.com/")

	el, err := d.FindElement(element.ByLinkText, "Inputs")
	if err != nil {
		t.Errorf("find el: %+v", err)
	}

	el.SendKeys(element.EnterKey)

	time.Sleep(15 * time.Second)
}
