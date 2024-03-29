package element

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// Click
// Performs click on an element
// Returns error
func (e *Element) Click() error {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Click)
	data, err := json.Marshal(&Empty{})
	if err != nil {
		log.Printf("Error on empty click marshal: %+v", err)
	}
  rr, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Error on click: %+v", err)
		return err
	}

  res := new(struct{ Value map[string]string })
  err = json.Unmarshal(rr, res)
  if res.Value["error"] != "" || err != nil {
    return err
  }

	return nil
}

// Special keys, for SendKeys.
const (
	NullKey       = string('\ue000')
	CancelKey     = string('\ue001')
	HelpKey       = string('\ue002')
	BackspaceKey  = string('\ue003')
	TabKey        = string('\ue004')
	ClearKey      = string('\ue005')
	ReturnKey     = string('\ue006')
	EnterKey      = string('\ue007')
	ShiftKey      = string('\ue008')
	ControlKey    = string('\ue009')
	AltKey        = string('\ue00a')
	PauseKey      = string('\ue00b')
	EscapeKey     = string('\ue00c')
	SpaceKey      = string('\ue00d')
	PageUpKey     = string('\ue00e')
	PageDownKey   = string('\ue00f')
	EndKey        = string('\ue010')
	HomeKey       = string('\ue011')
	LeftArrowKey  = string('\ue012')
	UpArrowKey    = string('\ue013')
	RightArrowKey = string('\ue014')
	DownArrowKey  = string('\ue015')
	InsertKey     = string('\ue016')
	DeleteKey     = string('\ue017')
	SemicolonKey  = string('\ue018')
	EqualsKey     = string('\ue019')
	Numpad0Key    = string('\ue01a')
	Numpad1Key    = string('\ue01b')
	Numpad2Key    = string('\ue01c')
	Numpad3Key    = string('\ue01d')
	Numpad4Key    = string('\ue01e')
	Numpad5Key    = string('\ue01f')
	Numpad6Key    = string('\ue020')
	Numpad7Key    = string('\ue021')
	Numpad8Key    = string('\ue022')
	Numpad9Key    = string('\ue023')
	MultiplyKey   = string('\ue024')
	AddKey        = string('\ue025')
	SeparatorKey  = string('\ue026')
	SubstractKey  = string('\ue027')
	DecimalKey    = string('\ue028')
	DivideKey     = string('\ue029')
	F1Key         = string('\ue031')
	F2Key         = string('\ue032')
	F3Key         = string('\ue033')
	F4Key         = string('\ue034')
	F5Key         = string('\ue035')
	F6Key         = string('\ue036')
	F7Key         = string('\ue037')
	F8Key         = string('\ue038')
	F9Key         = string('\ue039')
	F10Key        = string('\ue03a')
	F11Key        = string('\ue03b')
	F12Key        = string('\ue03c')
	MetaKey       = string('\ue03d')
)

// SendKeys
// Send keys
func (e *Element) SendKeys(s string) error {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Value)
	k := &SendKeys{
		Text: s,
	}

	data, err := json.Marshal(k)
	if err != nil {
		log.Printf("Send keys on marshal: %+v", err)
		return err
	}
	_, err = request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Click: %+v", err)
		return err
	}

	return nil
}
