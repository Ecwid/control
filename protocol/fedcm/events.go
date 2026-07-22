package fedcm

/*
 */
type DialogShown struct {
	DialogId   string     `json:"dialogId"`
	DialogType DialogType `json:"dialogType"`
	Accounts   []*Account `json:"accounts"`
	Title      string     `json:"title"`
	Subtitle   string     `json:"subtitle,omitempty"`
}

/*
	Triggered when a dialog is closed, either by user action, JS abort,

or a command below.
*/
type DialogClosed struct {
	DialogId string `json:"dialogId"`
}
