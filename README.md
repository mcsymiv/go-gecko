### go-gecko  
  
Testing framework based on [W3C Driver Protocol](https://w3c.github.io/webdriver/) with the use of [Firefox Gecko Driver](https://firefox-source-docs.mozilla.org/testing/geckodriver/index.html).  
In order to start test, ex. TestSession. Fire up gecko driver:
```
 ${PATH_TO_GECKO_DRIVER}/geckodriver > logs/gecko.session.logs 2>&1 &
```  
Here, driver is started and listening on `http://localhost:4444`.  
Aftet that, run the test:
```
go test -v -count=1 test/driver/driver_test.go -run TestDriver
```
`-v`, shows test output in verbose mode  
`-count=1`, discards test cache  
`test/driver/driver_test.go`, specifies test directory  
`-run`, pattern for test name  
  
```
func TestDriver(t *testing.T) {

	// Starts firefox browser
	s := driver.New()
	
	// Closes current session
	defer s.Quit()
	
	// Goes to the page
	s.Open("https://www.google.com")
	
	// Finds and returns an webelement
	el := s.FindElement(element.ByCssSelector, "#APjFqb")
	
	// Yep
	el.Click()
	
	// Sends text to the element, in this case it's is a google search input
	el.SendKeys("hello")
}
```
