package applicationcache

import (
	"github.com/ecwid/control/protocol/common"
)

/*
Detailed application cache resource information.
*/
type ApplicationCacheResource struct {
	Url  string `json:"url"`
	Size int    `json:"size"`
	Type string `json:"type"`
}

/*
Detailed application cache information.
*/
type ApplicationCache struct {
	ManifestURL  string                      `json:"manifestURL"`
	Size         float64                     `json:"size"`
	CreationTime float64                     `json:"creationTime"`
	UpdateTime   float64                     `json:"updateTime"`
	Resources    []*ApplicationCacheResource `json:"resources"`
}

/*
Frame identifier - manifest URL pair.
*/
type FrameWithManifest struct {
	FrameId     common.FrameId `json:"frameId"`
	ManifestURL string         `json:"manifestURL"`
	Status      int            `json:"status"`
}

type GetApplicationCacheForFrameArgs struct {
	FrameId common.FrameId `json:"frameId"`
}

type GetApplicationCacheForFrameVal struct {
	ApplicationCache *ApplicationCache `json:"applicationCache"`
}

type GetFramesWithManifestsVal struct {
	FrameIds []*FrameWithManifest `json:"frameIds"`
}

type GetManifestForFrameArgs struct {
	FrameId common.FrameId `json:"frameId"`
}

type GetManifestForFrameVal struct {
	ManifestURL string `json:"manifestURL"`
}
