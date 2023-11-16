package drivertest

import (
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/hooks"
	"log"
	"testing"
)

func TestElementAttribute(t *testing.T) {
	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	el, err := d.FindElement(element.ByLinkText, "Typos")
	if err != nil {
		log.Fatal("find element", err)
	}

	attr, _ := el.Attribute("href")
	if attr == "" {
		t.Errorf("element attribute: %+v", attr)
	}

}
