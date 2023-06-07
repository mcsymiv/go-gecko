package path

import (
	"fmt"
	"strings"
)

const (
	BaseUrl  = "http://localhost:4444"
	Session  = "session"
	Status   = "status"
	UrlPath  = "url"
	Element  = "element"
	Elements = "elements"
	Click    = "click"
	Value    = "value"

	// Element
	Attribute = "attribute"

	// Document
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
