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
//  2. Call request.Do(GET, url, nil)
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
func (ctx *ContextRequester) Get() string {

	// Url context wrapper for the path.UrlsArgs method
	url := ctx.Url()

	// Default GET method to the gecko driver
	res, err := request.Do(http.MethodGet, url, nil)
	if err != nil {

		// TODO: add error handling
		log.Printf("Error on request: %+v", err)
		return ""
	}

	// Unmarshal driver response for GET methods
	val := new(struct{ Value string })
	err = json.Unmarshal(res, val)
	if err != nil {

		// TODO: add error handling
		log.Printf("Error on unmarshal response: %+v", err)
		return ""
	}

	return val.Value
}

func (ctx *ContextRequester) Post(d interface{}) {
	url := ctx.Url()
	data, err := json.Marshal(d)
	if err != nil {
		log.Printf("Error marshal: %+v", err)
	}

	// Performs default POST method without returned value
	// Or returned value from driver is null: { "value": nul }
	_, err = request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Error request: %+v", err)
	}
}
