package smartcardemulation

/*
	Fired when |SCardEstablishContext| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gaa1b8970169fd4883a6dc4a8f43f19b67
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardestablishcontext
*/
type EstablishContextRequested struct {
	RequestId string `json:"requestId"`
}

/*
	Fired when |SCardReleaseContext| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga6aabcba7744c5c9419fdd6404f73a934
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardreleasecontext
*/
type ReleaseContextRequested struct {
	RequestId string `json:"requestId"`
	ContextId int    `json:"contextId"`
}

/*
	Fired when |SCardListReaders| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga93b07815789b3cf2629d439ecf20f0d9
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardlistreadersa
*/
type ListReadersRequested struct {
	RequestId string `json:"requestId"`
	ContextId int    `json:"contextId"`
}

/*
	Fired when |SCardGetStatusChange| is called. Timeout is specified in milliseconds.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga33247d5d1257d59e55647c3bb717db24
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardgetstatuschangea
*/
type GetStatusChangeRequested struct {
	RequestId    string           `json:"requestId"`
	ContextId    int              `json:"contextId"`
	ReaderStates []*ReaderStateIn `json:"readerStates"`
	Timeout      int              `json:"timeout,omitempty"`
}

/*
	Fired when |SCardCancel| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gaacbbc0c6d6c0cbbeb4f4debf6fbeeee6
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardcancel
*/
type CancelRequested struct {
	RequestId string `json:"requestId"`
	ContextId int    `json:"contextId"`
}

/*
	Fired when |SCardConnect| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga4e515829752e0a8dbc4d630696a8d6a5
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardconnecta
*/
type ConnectRequested struct {
	RequestId          string       `json:"requestId"`
	ContextId          int          `json:"contextId"`
	Reader             string       `json:"reader"`
	ShareMode          ShareMode    `json:"shareMode"`
	PreferredProtocols *ProtocolSet `json:"preferredProtocols"`
}

/*
	Fired when |SCardDisconnect| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga4be198045c73ec0deb79e66c0ca1738a
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scarddisconnect
*/
type DisconnectRequested struct {
	RequestId   string      `json:"requestId"`
	Handle      int         `json:"handle"`
	Disposition Disposition `json:"disposition"`
}

/*
	Fired when |SCardTransmit| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga9a2d77242a271310269065e64633ab99
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardtransmit
*/
type TransmitRequested struct {
	RequestId string   `json:"requestId"`
	Handle    int      `json:"handle"`
	Data      []byte   `json:"data"`
	Protocol  Protocol `json:"protocol,omitempty"`
}

/*
	Fired when |SCardControl| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gac3454d4657110fd7f753b2d3d8f4e32f
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardcontrol
*/
type ControlRequested struct {
	RequestId   string `json:"requestId"`
	Handle      int    `json:"handle"`
	ControlCode int    `json:"controlCode"`
	Data        []byte `json:"data"`
}

/*
	Fired when |SCardGetAttrib| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gaacfec51917255b7a25b94c5104961602
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardgetattrib
*/
type GetAttribRequested struct {
	RequestId string `json:"requestId"`
	Handle    int    `json:"handle"`
	AttribId  int    `json:"attribId"`
}

/*
	Fired when |SCardSetAttrib| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga060f0038a4ddfd5dd2b8fadf3c3a2e4f
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardsetattrib
*/
type SetAttribRequested struct {
	RequestId string `json:"requestId"`
	Handle    int    `json:"handle"`
	AttribId  int    `json:"attribId"`
	Data      []byte `json:"data"`
}

/*
	Fired when |SCardStatus| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gae49c3c894ad7ac12a5b896bde70d0382
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardstatusa
*/
type StatusRequested struct {
	RequestId string `json:"requestId"`
	Handle    int    `json:"handle"`
}

/*
	Fired when |SCardBeginTransaction| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gaddb835dce01a0da1d6ca02d33ee7d861
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardbegintransaction
*/
type BeginTransactionRequested struct {
	RequestId string `json:"requestId"`
	Handle    int    `json:"handle"`
}

/*
	Fired when |SCardEndTransaction| is called.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gae8742473b404363e5c587f570d7e2f3b
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardendtransaction
*/
type EndTransactionRequested struct {
	RequestId   string      `json:"requestId"`
	Handle      int         `json:"handle"`
	Disposition Disposition `json:"disposition"`
}
