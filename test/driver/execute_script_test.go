package driver

import (
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
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

	err = d.ExecuteScriptSync("document.querySelector('#content li a').click()", nil)
	if err != nil {
		t.Errorf("Failed script: %+v", err)
	}

	time.Sleep(3 * time.Second)

}
