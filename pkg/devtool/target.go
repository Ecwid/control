package devtool

// TargetInfos https://chromedevtools.github.io/devtools-protocol/tot/Target#method-getTargets
type TargetInfos struct {
	TargetInfos []*TargetInfo `json:"targetInfos"`
}

// TargetInfo https://chromedevtools.github.io/devtools-protocol/tot/Target#type-TargetInfo
type TargetInfo struct {
	TargetID         string `json:"targetId"`
	Type             string `json:"type"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	Attached         bool   `json:"attached"`
	OpenerID         string `json:"openerId"`
	BrowserContextID string `json:"browserContextId"`
}

// TargetCreated https://chromedevtools.github.io/devtools-protocol/tot/Target#event-targetCreated
type TargetCreated struct {
	TargetInfo *TargetInfo `json:"targetInfo"`
}

// TargetCrashed https://chromedevtools.github.io/devtools-protocol/tot/Target#event-targetCrashed
type TargetCrashed struct {
	TargetID  string `json:"targetId"`
	Status    string `json:"status"`
	ErrorCode int64  `json:"errorCode"`
}

// TargetDestroyed https://chromedevtools.github.io/devtools-protocol/tot/Target#event-targetDestroyed
type TargetDestroyed struct {
	TargetID string `json:"targetId"`
}
