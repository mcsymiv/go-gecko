package models

type Capabilities struct {
	AlwaysMatch
}

type AlwaysMatch struct {
	AcceptInsecureCerts bool
}
