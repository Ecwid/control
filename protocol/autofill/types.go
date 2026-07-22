package autofill

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
)

/*
 */
type CreditCard struct {
	Number      string `json:"number"`
	Name        string `json:"name"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	Cvc         string `json:"cvc"`
}

/*
 */
type AddressField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
A list of address fields.
*/
type AddressFields struct {
	Fields []*AddressField `json:"fields"`
}

/*
 */
type Address struct {
	Fields []*AddressField `json:"fields"`
}

/*
	Defines how an address can be displayed like in chrome://settings/addresses.

Address UI is a two dimensional array, each inner array is an "address information line", and when rendered in a UI surface should be displayed as such.
The following address UI for instance:
[[{name: "GIVE_NAME", value: "Jon"}, {name: "FAMILY_NAME", value: "Doe"}], [{name: "CITY", value: "Munich"}, {name: "ZIP", value: "81456"}]]
should allow the receiver to render:
Jon Doe
Munich 81456
*/
type AddressUI struct {
	AddressFields []*AddressFields `json:"addressFields"`
}

/*
Specified whether a filled field was done so by using the html autocomplete attribute or autofill heuristics.
*/
type FillingStrategy string

/*
 */
type FilledField struct {
	HtmlType        string            `json:"htmlType"`
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	Value           string            `json:"value"`
	AutofillType    string            `json:"autofillType"`
	FillingStrategy FillingStrategy   `json:"fillingStrategy"`
	FrameId         common.FrameId    `json:"frameId"`
	FieldId         dom.BackendNodeId `json:"fieldId"`
}

type TriggerArgs struct {
	FieldId dom.BackendNodeId `json:"fieldId"`
	FrameId common.FrameId    `json:"frameId,omitempty"`
	Card    *CreditCard       `json:"card,omitempty"`
	Address *Address          `json:"address,omitempty"`
}

type SetAddressesArgs struct {
	Addresses []*Address `json:"addresses"`
}
