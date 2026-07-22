package smartcardemulation

import (
	"github.com/ecwid/control/protocol"
)

/*
Enables the |SmartCardEmulation| domain.
*/
func Enable(c protocol.Caller) error {
	return c.Call("SmartCardEmulation.enable", nil, nil)
}

/*
Disables the |SmartCardEmulation| domain.
*/
func Disable(c protocol.Caller) error {
	return c.Call("SmartCardEmulation.disable", nil, nil)
}

/*
	Reports the successful result of a |SCardEstablishContext| call.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gaa1b8970169fd4883a6dc4a8f43f19b67
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardestablishcontext
*/
func ReportEstablishContextResult(c protocol.Caller, args ReportEstablishContextResultArgs) error {
	return c.Call("SmartCardEmulation.reportEstablishContextResult", args, nil)
}

/*
	Reports the successful result of a |SCardReleaseContext| call.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga6aabcba7744c5c9419fdd6404f73a934
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardreleasecontext
*/
func ReportReleaseContextResult(c protocol.Caller, args ReportReleaseContextResultArgs) error {
	return c.Call("SmartCardEmulation.reportReleaseContextResult", args, nil)
}

/*
	Reports the successful result of a |SCardListReaders| call.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga93b07815789b3cf2629d439ecf20f0d9
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardlistreadersa
*/
func ReportListReadersResult(c protocol.Caller, args ReportListReadersResultArgs) error {
	return c.Call("SmartCardEmulation.reportListReadersResult", args, nil)
}

/*
	Reports the successful result of a |SCardGetStatusChange| call.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga33247d5d1257d59e55647c3bb717db24
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardgetstatuschangea
*/
func ReportGetStatusChangeResult(c protocol.Caller, args ReportGetStatusChangeResultArgs) error {
	return c.Call("SmartCardEmulation.reportGetStatusChangeResult", args, nil)
}

/*
	Reports the result of a |SCardBeginTransaction| call.

On success, this creates a new transaction object.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gaddb835dce01a0da1d6ca02d33ee7d861
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardbegintransaction
*/
func ReportBeginTransactionResult(c protocol.Caller, args ReportBeginTransactionResultArgs) error {
	return c.Call("SmartCardEmulation.reportBeginTransactionResult", args, nil)
}

/*
	Reports the successful result of a call that returns only a result code.

Used for: |SCardCancel|, |SCardDisconnect|, |SCardSetAttrib|, |SCardEndTransaction|.

This maps to:
 1. SCardCancel
    PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gaacbbc0c6d6c0cbbeb4f4debf6fbeeee6
    Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardcancel

 2. SCardDisconnect
    PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga4be198045c73ec0deb79e66c0ca1738a
    Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scarddisconnect

 3. SCardSetAttrib
    PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga060f0038a4ddfd5dd2b8fadf3c3a2e4f
    Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardsetattrib

 4. SCardEndTransaction
    PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gae8742473b404363e5c587f570d7e2f3b
    Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardendtransaction
*/
func ReportPlainResult(c protocol.Caller, args ReportPlainResultArgs) error {
	return c.Call("SmartCardEmulation.reportPlainResult", args, nil)
}

/*
	Reports the successful result of a |SCardConnect| call.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga4e515829752e0a8dbc4d630696a8d6a5
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardconnecta
*/
func ReportConnectResult(c protocol.Caller, args ReportConnectResultArgs) error {
	return c.Call("SmartCardEmulation.reportConnectResult", args, nil)
}

/*
	Reports the successful result of a call that sends back data on success.

Used for |SCardTransmit|, |SCardControl|, and |SCardGetAttrib|.

This maps to:
 1. SCardTransmit
    PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#ga9a2d77242a271310269065e64633ab99
    Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardtransmit

 2. SCardControl
    PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gac3454d4657110fd7f753b2d3d8f4e32f
    Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardcontrol

 3. SCardGetAttrib
    PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gaacfec51917255b7a25b94c5104961602
    Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardgetattrib
*/
func ReportDataResult(c protocol.Caller, args ReportDataResultArgs) error {
	return c.Call("SmartCardEmulation.reportDataResult", args, nil)
}

/*
	Reports the successful result of a |SCardStatus| call.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__API.html#gae49c3c894ad7ac12a5b896bde70d0382
Microsoft: https://learn.microsoft.com/en-us/windows/win32/api/winscard/nf-winscard-scardstatusa
*/
func ReportStatusResult(c protocol.Caller, args ReportStatusResultArgs) error {
	return c.Call("SmartCardEmulation.reportStatusResult", args, nil)
}

/*
Reports an error result for the given request.
*/
func ReportError(c protocol.Caller, args ReportErrorArgs) error {
	return c.Call("SmartCardEmulation.reportError", args, nil)
}
