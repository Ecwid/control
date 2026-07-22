package filesystem

import (
	"github.com/ecwid/control/protocol"
)

/*
 */
func GetDirectory(c protocol.Caller, args GetDirectoryArgs) (*GetDirectoryVal, error) {
	var val = &GetDirectoryVal{}
	return val, c.Call("FileSystem.getDirectory", args, val)
}
