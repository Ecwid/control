package webauthn

/*
Triggered when a credential is added to an authenticator.
*/
type CredentialAdded struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	Credential      *Credential     `json:"credential"`
}

/*
	Triggered when a credential is deleted, e.g. through

PublicKeyCredential.signalUnknownCredential().
*/
type CredentialDeleted struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	CredentialId    []byte          `json:"credentialId"`
}

/*
	Triggered when a credential is updated, e.g. through

PublicKeyCredential.signalCurrentUserDetails().
*/
type CredentialUpdated struct {
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
