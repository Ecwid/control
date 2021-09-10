package animation

import (
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/runtime"
)

/*
	Animation instance.
*/
type Animation struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	PausedState  bool             `json:"pausedState"`
	PlayState    string           `json:"playState"`
	PlaybackRate float64          `json:"playbackRate"`
	StartTime    float64          `json:"startTime"`
	CurrentTime  float64          `json:"currentTime"`
	Type         string           `json:"type"`
	Source       *AnimationEffect `json:"source,omitempty"`
	CssId        string           `json:"cssId,omitempty"`
}

/*
	AnimationEffect instance
*/
type AnimationEffect struct {
	Delay          float64           `json:"delay"`
	EndDelay       float64           `json:"endDelay"`
	IterationStart float64           `json:"iterationStart"`
	Iterations     float64           `json:"iterations"`
	Duration       float64           `json:"duration"`
	Direction      string            `json:"direction"`
	Fill           string            `json:"fill"`
	BackendNodeId  dom.BackendNodeId `json:"backendNodeId,omitempty"`
	KeyframesRule  *KeyframesRule    `json:"keyframesRule,omitempty"`
	Easing         string            `json:"easing"`
}

/*
	Keyframes Rule
*/
type KeyframesRule struct {
	Name      string           `json:"name,omitempty"`
	Keyframes []*KeyframeStyle `json:"keyframes"`
}

/*
	Keyframe Style
*/
type KeyframeStyle struct {
	Offset string `json:"offset"`
	Easing string `json:"easing"`
}

type GetCurrentTimeArgs struct {
	Id string `json:"id"`
}

type GetCurrentTimeVal struct {
	CurrentTime float64 `json:"currentTime"`
}

type GetPlaybackRateVal struct {
	PlaybackRate float64 `json:"playbackRate"`
}

type ReleaseAnimationsArgs struct {
	Animations []string `json:"animations"`
}

type ResolveAnimationArgs struct {
	AnimationId string `json:"animationId"`
}

type ResolveAnimationVal struct {
	RemoteObject *runtime.RemoteObject `json:"remoteObject"`
}

type SeekAnimationsArgs struct {
	Animations  []string `json:"animations"`
	CurrentTime float64  `json:"currentTime"`
}

type SetPausedArgs struct {
	Animations []string `json:"animations"`
	Paused     bool     `json:"paused"`
}

type SetPlaybackRateArgs struct {
	PlaybackRate float64 `json:"playbackRate"`
}

type SetTimingArgs struct {
	AnimationId string  `json:"animationId"`
	Duration    float64 `json:"duration"`
	Delay       float64 `json:"delay"`
}
