package crashreportcontext

import (
	"github.com/ecwid/control/protocol"
)

/*
Returns all entries in the CrashReportContext across all frames in the page.
*/
func GetEntries(c protocol.Caller) (*GetEntriesVal, error) {
	var val = &GetEntriesVal{}
	return val, c.Call("CrashReportContext.getEntries", nil, val)
}
