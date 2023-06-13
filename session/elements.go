package session

import (
	"encoding/json"
	"log"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/selenium"
	"github.com/mcsymiv/go-gecko/strategy"
)

type ElementRequest struct {
	ElementUrl string
}

func (e *ElementRequest) Url() string {
	return e.ElementUrl
}

// FindElement
// Finds single element by specifying selector strategy and its value
// Uses Selenium 3 protocol UUID-based string constant
func (s *Session) FindElement(by, value string) (element.WebElement, error) {

	st := strategy.NewRequester(&ElementRequest{
		ElementUrl: path.UrlArgs(path.Session, s.Id, path.Element),
	})

	el := st.Post(&element.FindUsing{
		Using: by,
		Value: value,
	})

	res := new(struct{ Value map[string]string })
	if err := json.Unmarshal(el, &res); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
		return nil, err
	}

	// Retrieves w3c element id
	id := selenium.ElementID(res.Value)

	return &element.Element{
		SessionId: s.Id,
		Id:        id,
	}, nil
}

// FindElements
func (s *Session) FindElements(by, value string) (element.WebElements, error) {

	st := strategy.NewRequester(&ElementRequest{
		ElementUrl: path.UrlArgs(path.Session, s.Id, path.Elements),
	})

	el := st.Post(&element.FindUsing{
		Using: by,
		Value: value,
	})

	res := new(struct{ Value []map[string]string })
	if err := json.Unmarshal(el, &res); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
		return nil, err
	}

	els := selenium.ElementsID(res.Value)
	if els == nil {
		log.Printf("No elements found. Empty slice. Elements ids: %+v", els)
	}

	return &element.Elements{
		SessionId: s.Id,
		Ids:       els,
	}, nil
}

// Init
func (s *Session) Init(by, val string) element.WebElement {
	el, err := s.FindElement(by, val)
	if err != nil {
		log.Println("unable to find element", err, by, val)
		return nil
	}
	return el
}
