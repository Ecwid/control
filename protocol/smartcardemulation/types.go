package smartcardemulation

/*
	Indicates the PC/SC error code.

This maps to:
PC/SC Lite: https://pcsclite.apdu.fr/api/group__ErrorCodes.html
Microsoft: https://learn.microsoft.com/en-us/windows/win32/secauthn/authentication-return-values
*/
type ResultCode string

/*
Maps to the |SCARD_SHARE_*| values.
*/
type ShareMode string

/*
Indicates what the reader should do with the card.
*/
type Disposition string

/*
Maps to |SCARD_*| connection state values.
*/
type ConnectionState string

/*
Maps to the |SCARD_STATE_*| flags.
*/
type ReaderStateFlags struct {
	Unaware     bool `json:"unaware,omitempty"`
	Ignore      bool `json:"ignore,omitempty"`
	Changed     bool `json:"changed,omitempty"`
	Unknown     bool `json:"unknown,omitempty"`
	Unavailable bool `json:"unavailable,omitempty"`
	Empty       bool `json:"empty,omitempty"`
	Present     bool `json:"present,omitempty"`
	Exclusive   bool `json:"exclusive,omitempty"`
	Inuse       bool `json:"inuse,omitempty"`
	Mute        bool `json:"mute,omitempty"`
	Unpowered   bool `json:"unpowered,omitempty"`
}

/*
Maps to the |SCARD_PROTOCOL_*| flags.
*/
type ProtocolSet struct {
	T0  bool `json:"t0,omitempty"`
	T1  bool `json:"t1,omitempty"`
	Raw bool `json:"raw,omitempty"`
}

/*
Maps to the |SCARD_PROTOCOL_*| values.
*/
type Protocol string

/*
 */
type ReaderStateIn struct {
	Reader                string            `json:"reader"`
	CurrentState          *ReaderStateFlags `json:"currentState"`
	CurrentInsertionCount int               `json:"currentInsertionCount"`
}

/*
 */
type ReaderStateOut struct {
	Reader     string            `json:"reader"`
	EventState *ReaderStateFlags `json:"eventState"`
	EventCount int               `json:"eventCount"`
	Atr        []byte            `json:"atr"`
}

type ReportEstablishContextResultArgs struct {
	RequestId string `json:"requestId"`
	ContextId int    `json:"contextId"`
}

type ReportReleaseContextResultArgs struct {
	RequestId string `json:"requestId"`
}

type ReportListReadersResultArgs struct {
	RequestId string   `json:"requestId"`
	Readers   []string `json:"readers"`
}

type ReportGetStatusChangeResultArgs struct {
	RequestId    string            `json:"requestId"`
	ReaderStates []*ReaderStateOut `json:"readerStates"`
}

type ReportBeginTransactionResultArgs struct {
	RequestId string `json:"requestId"`
	Handle    int    `json:"handle"`
}

type ReportPlainResultArgs struct {
	RequestId string `json:"requestId"`
}

type ReportConnectResultArgs struct {
	RequestId      string   `json:"requestId"`
	Handle         int      `json:"handle"`
	ActiveProtocol Protocol `json:"activeProtocol,omitempty"`
}

type ReportDataResultArgs struct {
	RequestId string `json:"requestId"`
	Data      []byte `json:"data"`
}

type ReportStatusResultArgs struct {
	RequestId  string          `json:"requestId"`
	ReaderName string          `json:"readerName"`
	State      ConnectionState `json:"state"`
	Atr        []byte          `json:"atr"`
	Protocol   Protocol        `json:"protocol,omitempty"`
}

type ReportErrorArgs struct {
	RequestId  string     `json:"requestId"`
	ResultCode ResultCode `json:"resultCode"`
}
