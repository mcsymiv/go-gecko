package path

import (
	"fmt"
	"strings"
)

// gecko default port
// const BaseUrl = "http://localhost:4444"
// chrome default port
const BaseUrl = "http://localhost:9515"

const (
	Session    = "session"
	Status     = "status"
	UrlPath    = "url"
	Element    = "element"
	Elements   = "elements"
	Click      = "click"
	Value      = "value"
	Attribute  = "attribute"
	Text       = "text"
	PageSource = "source"
	Execute    = "execute"
	ScriptSync = "sync"
)

func Url(arg string) string {
	return fmt.Sprintf("%s/%s", BaseUrl, arg)
}

func UrlArgs(args ...string) string {
	return fmt.Sprintf("%s/%s", BaseUrl, strings.Join(args, "/"))
}
