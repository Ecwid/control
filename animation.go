package control

import (
	"github.com/ecwid/control/protocol"
	"github.com/ecwid/control/protocol/animation"
)

type Animation struct {
	s *Session
}

func (a Animation) Disable() error {
	return animation.Disable(a.s)
}

func (a Animation) Enable(c protocol.Caller) error {
	return animation.Enable(a.s)
}

func (a Animation) GetCurrentTime(args animation.GetCurrentTimeArgs) (*animation.GetCurrentTimeVal, error) {
	return animation.GetCurrentTime(a.s, args)
}

func (a Animation) GetPlaybackRate() (*animation.GetPlaybackRateVal, error) {
	return animation.GetPlaybackRate(a.s)
}

func (a Animation) ReleaseAnimations(args animation.ReleaseAnimationsArgs) error {
	return animation.ReleaseAnimations(a.s, args)
}

func (a Animation) ResolveAnimation(args animation.ResolveAnimationArgs) (*animation.ResolveAnimationVal, error) {
	return animation.ResolveAnimation(a.s, args)
}

func (a Animation) SeekAnimations(args animation.SeekAnimationsArgs) error {
	return animation.SeekAnimations(a.s, args)
}

func (a Animation) SetPaused(args animation.SetPausedArgs) error {
	return animation.SetPaused(a.s, args)
}

func (a Animation) SetPlaybackRate(args animation.SetPlaybackRateArgs) error {
	return animation.SetPlaybackRate(a.s, args)
}

func (a Animation) SetTiming(args animation.SetTimingArgs) error {
	return animation.SetTiming(a.s, args)
}
