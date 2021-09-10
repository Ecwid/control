package animation

import (
	"github.com/ecwid/control/protocol"
)

/*
	Disables animation domain notifications.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Animation.disable", nil, nil)
}

/*
	Enables animation domain notifications.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Animation.enable", nil, nil)
}

/*
	Returns the current time of the an animation.
*/
func GetCurrentTime(c protocol.Caller, args GetCurrentTimeArgs) (*GetCurrentTimeVal, error) {
	var val = &GetCurrentTimeVal{}
	return val, c.Call("Animation.getCurrentTime", args, val)
}

/*
	Gets the playback rate of the document timeline.
*/
func GetPlaybackRate(c protocol.Caller) (*GetPlaybackRateVal, error) {
	var val = &GetPlaybackRateVal{}
	return val, c.Call("Animation.getPlaybackRate", nil, val)
}

/*
	Releases a set of animations to no longer be manipulated.
*/
func ReleaseAnimations(c protocol.Caller, args ReleaseAnimationsArgs) error {
	return c.Call("Animation.releaseAnimations", args, nil)
}

/*
	Gets the remote object of the Animation.
*/
func ResolveAnimation(c protocol.Caller, args ResolveAnimationArgs) (*ResolveAnimationVal, error) {
	var val = &ResolveAnimationVal{}
	return val, c.Call("Animation.resolveAnimation", args, val)
}

/*
	Seek a set of animations to a particular time within each animation.
*/
func SeekAnimations(c protocol.Caller, args SeekAnimationsArgs) error {
	return c.Call("Animation.seekAnimations", args, nil)
}

/*
	Sets the paused state of a set of animations.
*/
func SetPaused(c protocol.Caller, args SetPausedArgs) error {
	return c.Call("Animation.setPaused", args, nil)
}

/*
	Sets the playback rate of the document timeline.
*/
func SetPlaybackRate(c protocol.Caller, args SetPlaybackRateArgs) error {
	return c.Call("Animation.setPlaybackRate", args, nil)
}

/*
	Sets the timing of an animation node.
*/
func SetTiming(c protocol.Caller, args SetTimingArgs) error {
	return c.Call("Animation.setTiming", args, nil)
}
