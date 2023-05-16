package request

import (
	"fmt"
	"strings"
)

const (
	BaseUrl = "http://localhost:4444"
	Session = "session"
	Status  = "status"
	UrlPath = "url"
	Element = "element"
	Click   = "click"
)

func Url(arg string) string {
	return fmt.Sprintf("%s/%s", BaseUrl, arg)
}

func UrlArgs(args ...string) string {
	return fmt.Sprintf("%s/%s", BaseUrl, strings.Join(args, "/"))
}
