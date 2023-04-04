package domsnapshot

import (
	"github.com/ecwid/control/protocol"
)

/*
Disables DOM snapshot agent for the given page.
*/
func Disable(c protocol.Caller) error {
	return c.Call("DOMSnapshot.disable", nil, nil)
}

/*
Enables DOM snapshot agent for the given page.
*/
func Enable(c protocol.Caller) error {
	return c.Call("DOMSnapshot.enable", nil, nil)
}

/*
	Returns a document snapshot, including the full DOM tree of the root node (including iframes,

template contents, and imported documents) in a flattened array, as well as layout and
white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is
flattened.
*/
func CaptureSnapshot(c protocol.Caller, args CaptureSnapshotArgs) (*CaptureSnapshotVal, error) {
	var val = &CaptureSnapshotVal{}
	return val, c.Call("DOMSnapshot.captureSnapshot", args, val)
}
