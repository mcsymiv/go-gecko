package util

import (
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/hooks"
	"github.com/mcsymiv/go-gecko/step"
)

func TestPortForward(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("http://192.168.0.1")
	d.IsPageLoaded()

	time.Sleep(2 * time.Second)
	inputBtn, _ := d.FindElement(element.ByCssSelector, "[id='local-login-button'] [title='LOG IN']")
	inputBtn.Click()

	time.Sleep(2 * time.Second)
	advanced, _ := d.FindElement(element.ByXPath, ".//span[contains(text(),'Advanced')]")
	advanced.Click()

	time.Sleep(2 * time.Second)
	natForward, _ := d.FindElement(element.ByXPath, ".//span[text()='NAT Forwarding']")
	natForward.Click()

	time.Sleep(2 * time.Second)
	portForward, _ := d.FindElement(element.ByXPath, ".//span[text()='Port Forwarding']")
	portForward.Click()

	time.Sleep(2 * time.Second)
	editIcon, _ := d.FindElement(element.ByXPath, ".//a[contains(@class,'btn-edit')]")
	editIcon.Click()

	time.Sleep(2 * time.Second)
	connDevices, _ := d.FindElement(element.ByXPath, ".//span[text()='VIEW CONNECTED DEVICES']")
	connDevices.Click()

	time.Sleep(2 * time.Second)
	myPcDevice, _ := d.FindElement(element.ByXPath, ".//span[contains(text(), 'mcs-pc')]/..")
	myPcDevice.Click()

	time.Sleep(2 * time.Second)
	saveBtn, _ := d.FindElement(element.ByXPath, ".//div[contains(@class,'msg-content-wrap')]//span[text()='SAVE']")
	saveBtn.Click()
}

