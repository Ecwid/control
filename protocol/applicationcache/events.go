package applicationcache

import (
	"github.com/ecwid/control/protocol/common"
)

/*

 */
type ApplicationCacheStatusUpdated struct {
	FrameId     common.FrameId `json:"frameId"`
	ManifestURL string         `json:"manifestURL"`
	Status      int            `json:"status"`
}

/*

 */
type NetworkStateUpdated struct {
	IsNowOnline bool `json:"isNowOnline"`
}
