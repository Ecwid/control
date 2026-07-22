package fedcm

/*
	Whether this is a sign-up or sign-in action for this account, i.e.

whether this account has ever been used to sign in to this RP before.
*/
type LoginState string

/*
The types of FedCM dialogs.
*/
type DialogType string

/*
The buttons on the FedCM dialog.
*/
type DialogButton string

/*
The URLs that each account has
*/
type AccountUrlType string

/*
Corresponds to IdentityRequestAccount
*/
type Account struct {
	AccountId         string     `json:"accountId"`
	Email             string     `json:"email"`
	Name              string     `json:"name"`
	GivenName         string     `json:"givenName"`
	PictureUrl        string     `json:"pictureUrl"`
	IdpConfigUrl      string     `json:"idpConfigUrl"`
	IdpLoginUrl       string     `json:"idpLoginUrl"`
	LoginState        LoginState `json:"loginState"`
	TermsOfServiceUrl string     `json:"termsOfServiceUrl,omitempty"`
	PrivacyPolicyUrl  string     `json:"privacyPolicyUrl,omitempty"`
}

type EnableArgs struct {
	DisableRejectionDelay bool `json:"disableRejectionDelay,omitempty"`
}

type SelectAccountArgs struct {
	DialogId     string `json:"dialogId"`
	AccountIndex int    `json:"accountIndex"`
}

type ClickDialogButtonArgs struct {
	DialogId     string       `json:"dialogId"`
	DialogButton DialogButton `json:"dialogButton"`
}

type OpenUrlArgs struct {
	DialogId       string         `json:"dialogId"`
	AccountIndex   int            `json:"accountIndex"`
	AccountUrlType AccountUrlType `json:"accountUrlType"`
}

type DismissDialogArgs struct {
	DialogId        string `json:"dialogId"`
	TriggerCooldown bool   `json:"triggerCooldown,omitempty"`
}
