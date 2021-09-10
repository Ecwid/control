package webauthn

/*

 */
type AuthenticatorId string

/*

 */
type AuthenticatorProtocol string

/*

 */
type Ctap2Version string

/*

 */
type AuthenticatorTransport string

/*

 */
type VirtualAuthenticatorOptions struct {
	Protocol                    AuthenticatorProtocol  `json:"protocol"`
	Ctap2Version                Ctap2Version           `json:"ctap2Version,omitempty"`
	Transport                   AuthenticatorTransport `json:"transport"`
	HasResidentKey              bool                   `json:"hasResidentKey,omitempty"`
	HasUserVerification         bool                   `json:"hasUserVerification,omitempty"`
	HasLargeBlob                bool                   `json:"hasLargeBlob,omitempty"`
	AutomaticPresenceSimulation bool                   `json:"automaticPresenceSimulation,omitempty"`
	IsUserVerified              bool                   `json:"isUserVerified,omitempty"`
}

/*

 */
type Credential struct {
	CredentialId         []byte `json:"credentialId"`
	IsResidentCredential bool   `json:"isResidentCredential"`
	RpId                 string `json:"rpId,omitempty"`
	PrivateKey           []byte `json:"privateKey"`
	UserHandle           []byte `json:"userHandle,omitempty"`
	SignCount            int    `json:"signCount"`
	LargeBlob            []byte `json:"largeBlob,omitempty"`
}

type AddVirtualAuthenticatorArgs struct {
	Options *VirtualAuthenticatorOptions `json:"options"`
}

type AddVirtualAuthenticatorVal struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
}

type RemoveVirtualAuthenticatorArgs struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
}

type AddCredentialArgs struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	Credential      *Credential     `json:"credential"`
}

type GetCredentialArgs struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	CredentialId    []byte          `json:"credentialId"`
}

type GetCredentialVal struct {
	Credential *Credential `json:"credential"`
}

type GetCredentialsArgs struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
}

type GetCredentialsVal struct {
	Credentials []*Credential `json:"credentials"`
}

type RemoveCredentialArgs struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	CredentialId    []byte          `json:"credentialId"`
}

type ClearCredentialsArgs struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
}

type SetUserVerifiedArgs struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	IsUserVerified  bool            `json:"isUserVerified"`
}

type SetAutomaticPresenceSimulationArgs struct {
	AuthenticatorId AuthenticatorId `json:"authenticatorId"`
	Enabled         bool            `json:"enabled"`
}
