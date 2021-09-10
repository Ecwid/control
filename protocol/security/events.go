package security

/*
	The security state of the page changed.
*/
type VisibleSecurityStateChanged struct {
	VisibleSecurityState *VisibleSecurityState `json:"visibleSecurityState"`
}

/*
	The security state of the page changed.
*/
type SecurityStateChanged struct {
	SecurityState SecurityState               `json:"securityState"`
	Explanations  []*SecurityStateExplanation `json:"explanations"`
	Summary       string                      `json:"summary,omitempty"`
}
