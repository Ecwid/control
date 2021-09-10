package browser

import (
	"github.com/ecwid/control/protocol/common"
)

/*
	Fired when page is about to start a download.
*/
type DownloadWillBegin struct {
	FrameId           common.FrameId `json:"frameId"`
	Guid              string         `json:"guid"`
	Url               string         `json:"url"`
	SuggestedFilename string         `json:"suggestedFilename"`
}

/*
	Fired when download makes progress. Last call has |done| == true.
*/
type DownloadProgress struct {
	Guid          string  `json:"guid"`
	TotalBytes    float64 `json:"totalBytes"`
	ReceivedBytes float64 `json:"receivedBytes"`
	State         string  `json:"state"`
}
