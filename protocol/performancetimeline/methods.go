package performancetimeline

import (
	"github.com/ecwid/control/protocol"
)

/*
	Previously buffered events would be reported before method returns.
See also: timelineEventAdded
*/
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("PerformanceTimeline.enable", args, nil)
}
