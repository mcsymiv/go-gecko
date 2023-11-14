package driver

import (
	"github.com/mcsymiv/go-gecko/hooks"
	"testing"

	"github.com/mcsymiv/go-gecko/element"
)

func TestExecuteScriptSync(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	err := d.ExecuteScriptSync("document.querySelector('#content li a').click()")
	if err != nil {
		t.Errorf("Failed script: %+v", err)
	}

	el, err := d.FindElement(element.ByLinkText, "Elemental Selenium")
	if err != nil {
		t.Errorf("Failed find el: %+v:", err)
	}

	err = d.ExecuteScriptSync("arguments[0].click()", el.ElementIdentifier())
	if err != nil {
		t.Errorf("Failed to pass el as arg: %+v", err)
	}
}
