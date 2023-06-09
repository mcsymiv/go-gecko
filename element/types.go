package element

const (
	ById              = "id" // not speciied by w3c
	ByXPath           = "xpath"
	ByLinkText        = "link text"
	ByPartialLinkText = "partial link text"
	ByName            = "name" // not specified by w3c
	ByTagName         = "tag name"
	ByClassName       = "class name" // not specified by w3c
	ByCssSelector     = "css selector"
)

type WebElement interface {
	ElementId() (string, error)
	ElementIdentifier() map[string]string
	Click() error
	SendKeys(keys string) error

	// State
	Attribute(attr string) string
	Text() string
}

type WebElements interface {
	Elements() ([]WebElement, error)
}

type Element struct {
	SessionId string
	Id        string
}

type Elements struct {
	SessionId string
	Ids       []string
}

type SendKeys struct {
	Text string `json:"text"`
}

// Empty
// Due to geckodriver bug: https://github.com/webdriverio/webdriverio/pull/3208
// "where Geckodriver requires POST requests to have a valid JSON body"
// Used in POST requests that don't require data to be passed by W3C
type Empty struct{}

type FindUsing struct {
	Using string `json:"using"`
	Value string `json:"value"`
}
