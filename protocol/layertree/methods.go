package layertree

import (
	"github.com/ecwid/control/protocol"
)

/*
Provides the reasons why the given layer was composited.
*/
func CompositingReasons(c protocol.Caller, args CompositingReasonsArgs) (*CompositingReasonsVal, error) {
	var val = &CompositingReasonsVal{}
	return val, c.Call("LayerTree.compositingReasons", args, val)
}

/*
Disables compositing tree inspection.
*/
func Disable(c protocol.Caller) error {
	return c.Call("LayerTree.disable", nil, nil)
}

/*
Enables compositing tree inspection.
*/
func Enable(c protocol.Caller) error {
	return c.Call("LayerTree.enable", nil, nil)
}

/*
Returns the snapshot identifier.
*/
func LoadSnapshot(c protocol.Caller, args LoadSnapshotArgs) (*LoadSnapshotVal, error) {
	var val = &LoadSnapshotVal{}
	return val, c.Call("LayerTree.loadSnapshot", args, val)
}

/*
Returns the layer snapshot identifier.
*/
func MakeSnapshot(c protocol.Caller, args MakeSnapshotArgs) (*MakeSnapshotVal, error) {
	var val = &MakeSnapshotVal{}
	return val, c.Call("LayerTree.makeSnapshot", args, val)
}

/*
 */
func ProfileSnapshot(c protocol.Caller, args ProfileSnapshotArgs) (*ProfileSnapshotVal, error) {
	var val = &ProfileSnapshotVal{}
	return val, c.Call("LayerTree.profileSnapshot", args, val)
}

/*
Releases layer snapshot captured by the back-end.
*/
func ReleaseSnapshot(c protocol.Caller, args ReleaseSnapshotArgs) error {
	return c.Call("LayerTree.releaseSnapshot", args, nil)
}

/*
Replays the layer snapshot and returns the resulting bitmap.
*/
func ReplaySnapshot(c protocol.Caller, args ReplaySnapshotArgs) (*ReplaySnapshotVal, error) {
	var val = &ReplaySnapshotVal{}
	return val, c.Call("LayerTree.replaySnapshot", args, val)
}

/*
Replays the layer snapshot and returns canvas log.
*/
func SnapshotCommandLog(c protocol.Caller, args SnapshotCommandLogArgs) (*SnapshotCommandLogVal, error) {
	var val = &SnapshotCommandLogVal{}
	return val, c.Call("LayerTree.snapshotCommandLog", args, val)
}
