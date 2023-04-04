package webauthn

/*
Triggered when a credential is added to an authenticator.
*/
type CredentialAdded struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	Credential      *Credential     `json:"credential"`
}

/*
Triggered when a credential is used in a webauthn assertion.
*/
type CredentialAsserted struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	Credential      *Credential     `json:"credential"`
}
