package driver

import (
	"log"
	"testing"
	"github.com/mcsymiv/go-gecko/element"
)

func TestElementAttribute(t *testing.T) {
  d, tear := SetupTest()
  defer tear()

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
