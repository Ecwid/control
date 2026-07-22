package crashreportcontext

import (
	"github.com/ecwid/control/protocol/common"
)

/*
Key-value pair in CrashReportContext.
*/
type CrashReportContextEntry struct {
	Key     string         `json:"key"`
	Value   string         `json:"value"`
	FrameId common.FrameId `json:"frameId"`
}

type GetEntriesVal struct {
	Entries []*CrashReportContextEntry `json:"entries"`
}
