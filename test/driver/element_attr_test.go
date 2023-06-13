package driver

import (
	"log"
	"testing"
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
)

func TestElementAttribute(t *testing.T) {

	d, err := session.New()
	if err != nil {
		log.Fatal("start session", err)
	}

	defer d.Quit()

	d.Open("https://the-internet.herokuapp.com/")

	el, err := d.FindElement(element.ByLinkText, "Typos")
	if err != nil {
		log.Fatal("find element", err)
	}

	attr := el.Attribute("href")
	if attr == "" {
		t.Errorf("element attribute: %+v", attr)
	}

}
