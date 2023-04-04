package storage

import (
	"github.com/ecwid/control/protocol/common"
)

/*
A cache's contents have been modified.
*/
type CacheStorageContentUpdated struct {
	Origin     string `json:"origin"`
	StorageKey string `json:"storageKey"`
	CacheName  string `json:"cacheName"`
}

/*
A cache has been added/deleted.
*/
type CacheStorageListUpdated struct {
	Origin     string `json:"origin"`
	StorageKey string `json:"storageKey"`
}

/*
The origin's IndexedDB object store has been modified.
*/
type IndexedDBContentUpdated struct {
	Origin          string `json:"origin"`
	StorageKey      string `json:"storageKey"`
	DatabaseName    string `json:"databaseName"`
	ObjectStoreName string `json:"objectStoreName"`
}

/*
The origin's IndexedDB database list has been modified.
*/
type IndexedDBListUpdated struct {
	Origin     string `json:"origin"`
	StorageKey string `json:"storageKey"`
}

/*
One of the interest groups was accessed by the associated page.
*/
type InterestGroupAccessed struct {
	AccessTime  common.TimeSinceEpoch   `json:"accessTime"`
	Type        InterestGroupAccessType `json:"type"`
	OwnerOrigin string                  `json:"ownerOrigin"`
	Name        string                  `json:"name"`
}

/*
	Shared storage was accessed by the associated page.

The following parameters are included in all events.
*/
type SharedStorageAccessed struct {
	AccessTime  common.TimeSinceEpoch      `json:"accessTime"`
	Type        SharedStorageAccessType    `json:"type"`
	MainFrameId common.FrameId             `json:"mainFrameId"`
	OwnerOrigin string                     `json:"ownerOrigin"`
	Params      *SharedStorageAccessParams `json:"params"`
}
