package storage

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/target"
)

/*
A cache's contents have been modified.
*/
type CacheStorageContentUpdated struct {
	Origin     string `json:"origin"`
	StorageKey string `json:"storageKey"`
	BucketId   string `json:"bucketId"`
	CacheName  string `json:"cacheName"`
}

/*
A cache has been added/deleted.
*/
type CacheStorageListUpdated struct {
	Origin     string `json:"origin"`
	StorageKey string `json:"storageKey"`
	BucketId   string `json:"bucketId"`
}

/*
The origin's IndexedDB object store has been modified.
*/
type IndexedDBContentUpdated struct {
	Origin          string `json:"origin"`
	StorageKey      string `json:"storageKey"`
	BucketId        string `json:"bucketId"`
	DatabaseName    string `json:"databaseName"`
	ObjectStoreName string `json:"objectStoreName"`
}

/*
The origin's IndexedDB database list has been modified.
*/
type IndexedDBListUpdated struct {
	Origin     string `json:"origin"`
	StorageKey string `json:"storageKey"`
	BucketId   string `json:"bucketId"`
}

/*
	One of the interest groups was accessed. Note that these events are global

to all targets sharing an interest group store.
*/
type InterestGroupAccessed struct {
	AccessTime            common.TimeSinceEpoch   `json:"accessTime"`
	Type                  InterestGroupAccessType `json:"type"`
	OwnerOrigin           string                  `json:"ownerOrigin"`
	Name                  string                  `json:"name"`
	ComponentSellerOrigin string                  `json:"componentSellerOrigin,omitempty"`
	Bid                   float64                 `json:"bid,omitempty"`
	BidCurrency           string                  `json:"bidCurrency,omitempty"`
	UniqueAuctionId       InterestGroupAuctionId  `json:"uniqueAuctionId,omitempty"`
}

/*
	An auction involving interest groups is taking place. These events are

target-specific.
*/
type InterestGroupAuctionEventOccurred struct {
	EventTime       common.TimeSinceEpoch         `json:"eventTime"`
	Type            InterestGroupAuctionEventType `json:"type"`
	UniqueAuctionId InterestGroupAuctionId        `json:"uniqueAuctionId"`
	ParentAuctionId InterestGroupAuctionId        `json:"parentAuctionId,omitempty"`
	AuctionConfig   any                           `json:"auctionConfig,omitempty"`
}

/*
	Specifies which auctions a particular network fetch may be related to, and

in what role. Note that it is not ordered with respect to
Network.requestWillBeSent (but will happen before loadingFinished
loadingFailed).
*/
type InterestGroupAuctionNetworkRequestCreated struct {
	Type      InterestGroupAuctionFetchType `json:"type"`
	RequestId network.RequestId             `json:"requestId"`
	Auctions  []InterestGroupAuctionId      `json:"auctions"`
}

/*
	Shared storage was accessed by the associated page.

The following parameters are included in all events.
*/
type SharedStorageAccessed struct {
	AccessTime  common.TimeSinceEpoch      `json:"accessTime"`
	Scope       SharedStorageAccessScope   `json:"scope"`
	Method      SharedStorageAccessMethod  `json:"method"`
	MainFrameId common.FrameId             `json:"mainFrameId"`
	OwnerOrigin string                     `json:"ownerOrigin"`
	OwnerSite   string                     `json:"ownerSite"`
	Params      *SharedStorageAccessParams `json:"params"`
}

/*
	A shared storage run or selectURL operation finished its execution.

The following parameters are included in all events.
*/
type SharedStorageWorkletOperationExecutionFinished struct {
	FinishedTime    common.TimeSinceEpoch     `json:"finishedTime"`
	ExecutionTime   int                       `json:"executionTime"`
	Method          SharedStorageAccessMethod `json:"method"`
	OperationId     string                    `json:"operationId"`
	WorkletTargetId target.TargetID           `json:"workletTargetId"`
	MainFrameId     common.FrameId            `json:"mainFrameId"`
	OwnerOrigin     string                    `json:"ownerOrigin"`
}

/*
 */
type StorageBucketCreatedOrUpdated struct {
	BucketInfo *StorageBucketInfo `json:"bucketInfo"`
}

/*
 */
type StorageBucketDeleted struct {
	BucketId string `json:"bucketId"`
}
