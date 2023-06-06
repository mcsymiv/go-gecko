### go-gecko  
  
Testing framework based on [W3C Driver Protocol](https://w3c.github.io/webdriver/) with the use of [Firefox Gecko Driver](https://firefox-source-docs.mozilla.org/testing/geckodriver/index.html).  

Run the test:
```
go test -v -count=1 test/driver/*.go -run TestDriver
```
`-v`, shows test output in verbose mode  
`-count=1`, discards test cache  
`test/driver/driver_test.go`, specifies test directory  
`-run`, pattern for test name  
  
This command will start TestDriver with TestMain setup routine first.
`TestMain`, starts Gecko driver for zsh, with default GeckoDriverPath  
and will redirect driver stdout/stderr into `logs/gecko.session.logs` file.	
```
exec.Command("zsh", "-c", GeckoDriverPath, ">", "logs/gecko.session.logs", "2>&1", "&")
```

TestDriver  
```
func TestDriver(t *testing.T) {

	caps := capabilities.ImplicitWait(3000)
	d, err := driver.New(caps)
	if err != nil {
		log.Fatal("session start err", err)
	}
	defer d.Quit()

	_, err = d.Open("https://www.google.com")
	if err != nil {
		t.Errorf("Url: %+v", err)
	}
}
```
