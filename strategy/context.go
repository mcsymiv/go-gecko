package strategy

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/request"
)

// Requester
// Implements strategy pattern
// For some of the repeated request cases:
//  1. Create url endpoint
//  2. Call request.Do(METHOD, URL, DATA)
//  3. Unmarshal response to new(struct{ Value string})
type Requester interface {
	Url() string
}

// ContextRequester
type ContextRequester struct {
	Requester
}

// NewRequester
// Creates new ContextRequester
// Accepts struct that implements Requester interface
func NewRequester(r Requester) *ContextRequester {
	return &ContextRequester{
		Requester: r,
	}
}

// SetRequester
// Setter mentod for new ContextRequester
func (ctx *ContextRequester) SetRequester(r Requester) *ContextRequester {
	ctx.Requester = r
	return ctx
}

// Get
// Wraps GET method to the driver
// And performs Default unmarshal for the response
// Response example: { "value": string }
//
// 1. request.Do(GET, url, nil)
// 2. unmarshal response
// 3. return Value string
func (ctx *ContextRequester) Get() string {

	// Url context wrapper for the path.UrlsArgs method
	url := ctx.Url()

	// Default GET method to the gecko driver
	res, err := request.Do(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error on request: %+v", err)
		return ""
	}

	// Unmarshal driver response for GET methods
	val := new(struct{ Value string })
	err = json.Unmarshal(res, val)
	if err != nil {
		log.Printf("Error on unmarshal response: %+v", err)
		return ""
	}

	return val.Value
}

// Post PostDefault
// Performs default POST method without returned value
// Or returned value from driver is null: { "value": null }
//
// 1. marshal data
// 2. request.Do(POST, url, data)
// 3. return []byte response for client to handle umarshal
func (ctx *ContextRequester) Post(d interface{}) []byte {

	url := ctx.Url()

	data, err := json.Marshal(d)
	if err != nil {
		log.Printf("Error marshal: %+v", err)
	}

	res, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Error request: %+v", err)
	}

	// Response raw
	// Checks for { "value" : { "error": ... }}
	// Response from driver
	rr := new(struct{ Value map[string]interface{} })
	err = json.Unmarshal(res, rr)
	if rr.Value["error"] != nil {
		log.Printf("ERROR: %+v", rr)
	}

	return res
}
