package security

/*
The security state of the page changed.
*/
type VisibleSecurityStateChanged struct {
	VisibleSecurityState *VisibleSecurityState `json:"visibleSecurityState"`
}
