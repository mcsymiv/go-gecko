package existing

import (
	"log"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
)

func TestGecko(t *testing.T) {

	moz := &capabilities.MozOptions{
		Args: []string{"--profile", "/Users/mcs/Library/Application Support/Firefox/Profiles/uupibms2.default-release"},
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
		log.Fatal("find element", err)
	}

	el.Click()

	time.Sleep(5 * time.Second)

}
