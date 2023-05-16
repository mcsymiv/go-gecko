package element

const (
	ById              = "id"
	ByXPath           = "xpath"
	ByLinkText        = "link text"
	ByPartialLinkText = "partial link text"
	ByName            = "name"
	ByTagName         = "tag name"
	ByClassName       = "class name"
	ByCssSelector     = "css selector"
)

type WebElement interface {
	Click()
	// SendKeys(keys string) error
}

type Element struct {
	SessionId string
	Id        string
}

type FindUsing struct {
	Using string `json:"using"`
	Value string `json:"value"`
}
