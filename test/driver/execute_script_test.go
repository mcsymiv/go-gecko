package driver

import (
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
)

func TestExecuteScriptSync(t *testing.T) {

	d, err := driver.New(capabilities.ImplicitWait(3000))
	if err != nil {
		t.Errorf("Failed session: %+v", err)
	}

	defer d.Quit()

	_, err = d.Open("https://the-internet.herokuapp.com/")
	if err != nil {
		t.Errorf("Failed open url: %+v", err)
	}

	err = d.ExecuteScriptSync("document.querySelector('#content li a').click()")
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

	time.Sleep(1 * time.Second)
}
