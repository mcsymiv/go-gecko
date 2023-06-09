package driver

import (
	"log"
	"testing"

	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

func TestElementAttribute(t *testing.T) {

	d, err := driver.New()
	if err != nil {
		log.Fatal("start session", err)
	}

	defer d.Quit()

	_, err = d.Open("https://the-internet.herokuapp.com/")
	if err != nil {
		log.Fatal("open url", err)
	}

	el, err := d.FindElement(element.ByLinkText, "Typos")
	if err != nil {
		log.Fatal("find element", err)
	}

	attr := el.Attribute("href")
	if attr == "" {
		t.Errorf("element attribute: %+v", err)
	}

}
