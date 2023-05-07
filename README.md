### go-gecko  
  
Testing framework based on [W3C Driver Protocol](https://w3c.github.io/webdriver/) with the use of [Firefox Gecko Driver](https://firefox-source-docs.mozilla.org/testing/geckodriver/index.html).  
In order to start test, ex. TestSession. Fire up gecko driver:
```
 ${PATH_TO_GECKO_DRIVER}/geckodriver > logs/gecko.session.logs 2>&1 &
```  
Here, driver is started and listening on `http://localhost:4444`.  
Aftet that, run the test:
```
go test -v -count=1 test/demo/session_test.go -run TestSession
```
`-v`, shows test output in verbose mode  
`-count=1`, discards test cache  
`test/demo/session_test.go`, specifies test directory  
`-run`, pattern for test name  
```
func TestSession(t *testing.T) {

	// Starts firefox browser
  // Returns POST /session response,
  // Which contains session id
	sc := session.New()
  
  // Assigns session id to the local struct
  // For use in test steps
	st := &SessionTest{
		Id: sc.Id,
	}
  
  // Closes created gecko session
  // After the test
	defer st.CloseSession(t)

	// Test steps
	st.GetSessionStatus(t)
	st.OpenUrl(t)
}
```
