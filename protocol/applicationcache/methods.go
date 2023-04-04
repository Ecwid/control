package applicationcache

import (
	"github.com/ecwid/control/protocol"
)

/*
Enables application cache domain notifications.
*/
func Enable(c protocol.Caller) error {
	return c.Call("ApplicationCache.enable", nil, nil)
}

/*
Returns relevant application cache data for the document in given frame.
*/
func GetApplicationCacheForFrame(c protocol.Caller, args GetApplicationCacheForFrameArgs) (*GetApplicationCacheForFrameVal, error) {
	var val = &GetApplicationCacheForFrameVal{}
	return val, c.Call("ApplicationCache.getApplicationCacheForFrame", args, val)
}

/*
	Returns array of frame identifiers with manifest urls for each frame containing a document

associated with some application cache.
*/
func GetFramesWithManifests(c protocol.Caller) (*GetFramesWithManifestsVal, error) {
	var val = &GetFramesWithManifestsVal{}
	return val, c.Call("ApplicationCache.getFramesWithManifests", nil, val)
}

/*
Returns manifest URL for document in the given frame.
*/
func GetManifestForFrame(c protocol.Caller, args GetManifestForFrameArgs) (*GetManifestForFrameVal, error) {
	var val = &GetManifestForFrameVal{}
	return val, c.Call("ApplicationCache.getManifestForFrame", args, val)
}
