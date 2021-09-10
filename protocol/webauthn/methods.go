package webauthn

import (
	"github.com/ecwid/control/protocol"
)

/*
	Enable the WebAuthn domain and start intercepting credential storage and
retrieval with a virtual authenticator.
*/
func Enable(c protocol.Caller) error {
	return c.Call("WebAuthn.enable", nil, nil)
}

/*
	Disable the WebAuthn domain.
*/
func Disable(c protocol.Caller) error {
	return c.Call("WebAuthn.disable", nil, nil)
}

/*
	Creates and adds a virtual authenticator.
*/
func AddVirtualAuthenticator(c protocol.Caller, args AddVirtualAuthenticatorArgs) (*AddVirtualAuthenticatorVal, error) {
	var val = &AddVirtualAuthenticatorVal{}
	return val, c.Call("WebAuthn.addVirtualAuthenticator", args, val)
}

/*
	Removes the given authenticator.
*/
func RemoveVirtualAuthenticator(c protocol.Caller, args RemoveVirtualAuthenticatorArgs) error {
	return c.Call("WebAuthn.removeVirtualAuthenticator", args, nil)
}

/*
	Adds the credential to the specified authenticator.
*/
func AddCredential(c protocol.Caller, args AddCredentialArgs) error {
	return c.Call("WebAuthn.addCredential", args, nil)
}

/*
	Returns a single credential stored in the given virtual authenticator that
matches the credential ID.
*/
func GetCredential(c protocol.Caller, args GetCredentialArgs) (*GetCredentialVal, error) {
	var val = &GetCredentialVal{}
	return val, c.Call("WebAuthn.getCredential", args, val)
}

/*
	Returns all the credentials stored in the given virtual authenticator.
*/
func GetCredentials(c protocol.Caller, args GetCredentialsArgs) (*GetCredentialsVal, error) {
	var val = &GetCredentialsVal{}
	return val, c.Call("WebAuthn.getCredentials", args, val)
}

/*
	Removes a credential from the authenticator.
*/
func RemoveCredential(c protocol.Caller, args RemoveCredentialArgs) error {
	return c.Call("WebAuthn.removeCredential", args, nil)
}

/*
	Clears all the credentials from the specified device.
*/
func ClearCredentials(c protocol.Caller, args ClearCredentialsArgs) error {
	return c.Call("WebAuthn.clearCredentials", args, nil)
}

/*
	Sets whether User Verification succeeds or fails for an authenticator.
The default is true.
*/
func SetUserVerified(c protocol.Caller, args SetUserVerifiedArgs) error {
	return c.Call("WebAuthn.setUserVerified", args, nil)
}

/*
	Sets whether tests of user presence will succeed immediately (if true) or fail to resolve (if false) for an authenticator.
The default is true.
*/
func SetAutomaticPresenceSimulation(c protocol.Caller, args SetAutomaticPresenceSimulationArgs) error {
	return c.Call("WebAuthn.setAutomaticPresenceSimulation", args, nil)
}
