package autofill

/*
Emitted when an address form is filled.
*/
type AddressFormFilled struct {
	FilledFields []*FilledField `json:"filledFields"`
	AddressUi    *AddressUI     `json:"addressUi"`
}
