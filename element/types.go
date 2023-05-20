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
	Click() error
	SendKeys(keys string) error
	Attribute(attr string) (string, error)
}

type Element struct {
	SessionId string
	Id        string
}

type SendKeys struct {
	Text string `json:"text"`
}

type FindUsing struct {
	Using string `json:"using"`
	Value string `json:"value"`
}
