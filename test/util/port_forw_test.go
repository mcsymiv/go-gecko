package util

import (
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/hooks"
)

func TestPortForward(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("http://192.168.0.1")
  time.Sleep(5*time.Second)

  _, _ = d.FindElement(element.ByCssSelector, "[id='local-pwd-tb'] [type='password']")
  time.Sleep(1*time.Second)

  inputBtn, _ := d.FindElement(element.ByCssSelector, "[id='local-login-button'] [title='LOG IN']")
  inputBtn.Click()
  time.Sleep(1*time.Second)
}
