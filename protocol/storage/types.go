package storage

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/network"
)

/*
 */
type SerializedStorageKey string

/*
Enum of possible storage types.
*/
type StorageType string

/*
Usage for a storage type.
*/
type UsageForType struct {
	StorageType StorageType `json:"storageType"`
	Usage       float64     `json:"usage"`
}

/*
	Pair of issuer origin and number of available (signed, but not used) Trust

Tokens from that issuer.
*/
type TrustTokens struct {
	IssuerOrigin string  `json:"issuerOrigin"`
	Count        float64 `json:"count"`
}

/*
Enum of interest group access types.
*/
type InterestGroupAccessType string

/*
Ad advertising element inside an interest group.
*/
type InterestGroupAd struct {
	RenderUrl string `json:"renderUrl"`
	Metadata  string `json:"metadata,omitempty"`
}

/*
The full details of an interest group.
*/
type InterestGroupDetails struct {
	OwnerOrigin               string                `json:"ownerOrigin"`
	Name                      string                `json:"name"`
	ExpirationTime            common.TimeSinceEpoch `json:"expirationTime"`
	JoiningOrigin             string                `json:"joiningOrigin"`
	BiddingUrl                string                `json:"biddingUrl,omitempty"`
	BiddingWasmHelperUrl      string                `json:"biddingWasmHelperUrl,omitempty"`
	UpdateUrl                 string                `json:"updateUrl,omitempty"`
	TrustedBiddingSignalsUrl  string                `json:"trustedBiddingSignalsUrl,omitempty"`
	TrustedBiddingSignalsKeys []string              `json:"trustedBiddingSignalsKeys"`
	UserBiddingSignals        string                `json:"userBiddingSignals,omitempty"`
	Ads                       []*InterestGroupAd    `json:"ads"`
	AdComponents              []*InterestGroupAd    `json:"adComponents"`
}

/*
Enum of shared storage access types.
*/
type SharedStorageAccessType string

/*
Struct for a single key-value pair in an origin's shared storage.
*/
type SharedStorageEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

/*
Details for an origin's shared storage.
*/
type SharedStorageMetadata struct {
	CreationTime    common.TimeSinceEpoch `json:"creationTime"`
	Length          int                   `json:"length"`
	RemainingBudget float64               `json:"remainingBudget"`
}

/*
Pair of reporting metadata details for a candidate URL for `selectURL()`.
*/
type SharedStorageReportingMetadata struct {
	EventType    string `json:"eventType"`
	ReportingUrl string `json:"reportingUrl"`
}

/*
Bundles a candidate URL with its reporting metadata.
*/
type SharedStorageUrlWithMetadata struct {
	Url               string                            `json:"url"`
	ReportingMetadata []*SharedStorageReportingMetadata `json:"reportingMetadata"`
}

/*
	Bundles the parameters for shared storage access events whose

presence/absence can vary according to SharedStorageAccessType.
*/
type SharedStorageAccessParams struct {
	ScriptSourceUrl  string                          `json:"scriptSourceUrl,omitempty"`
	OperationName    string                          `json:"operationName,omitempty"`
	SerializedData   string                          `json:"serializedData,omitempty"`
	UrlsWithMetadata []*SharedStorageUrlWithMetadata `json:"urlsWithMetadata,omitempty"`
	Key              string                          `json:"key,omitempty"`
	Value            string                          `json:"value,omitempty"`
	IgnoreIfPresent  bool                            `json:"ignoreIfPresent,omitempty"`
}

type GetStorageKeyForFrameArgs struct {
	FrameId common.FrameId `json:"frameId"`
}

type GetStorageKeyForFrameVal struct {
	StorageKey SerializedStorageKey `json:"storageKey"`
}

type ClearDataForOriginArgs struct {
	Origin       string `json:"origin"`
	StorageTypes string `json:"storageTypes"`
}

type ClearDataForStorageKeyArgs struct {
	StorageKey   string `json:"storageKey"`
	StorageTypes string `json:"storageTypes"`
}

type GetCookiesArgs struct {
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
}

type GetCookiesVal struct {
	Cookies []*network.Cookie `json:"cookies"`
}

type SetCookiesArgs struct {
	Cookies          []*network.CookieParam  `json:"cookies"`
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
}

type ClearCookiesArgs struct {
	BrowserContextId common.BrowserContextID `json:"browserContextId,omitempty"`
}

type GetUsageAndQuotaArgs struct {
	Origin string `json:"origin"`
}

type GetUsageAndQuotaVal struct {
	Usage          float64         `json:"usage"`
	Quota          float64         `json:"quota"`
	OverrideActive bool            `json:"overrideActive"`
	UsageBreakdown []*UsageForType `json:"usageBreakdown"`
}

type OverrideQuotaForOriginArgs struct {
	Origin    string  `json:"origin"`
	QuotaSize float64 `json:"quotaSize,omitempty"`
}

type TrackCacheStorageForOriginArgs struct {
	Origin string `json:"origin"`
}

type TrackCacheStorageForStorageKeyArgs struct {
	StorageKey string `json:"storageKey"`
}

type TrackIndexedDBForOriginArgs struct {
	Origin string `json:"origin"`
}

type TrackIndexedDBForStorageKeyArgs struct {
	StorageKey string `json:"storageKey"`
}

type UntrackCacheStorageForOriginArgs struct {
	Origin string `json:"origin"`
}

type UntrackCacheStorageForStorageKeyArgs struct {
	StorageKey string `json:"storageKey"`
}

type UntrackIndexedDBForOriginArgs struct {
	Origin string `json:"origin"`
}

type UntrackIndexedDBForStorageKeyArgs struct {
	StorageKey string `json:"storageKey"`
}

type GetTrustTokensVal struct {
	Tokens []*TrustTokens `json:"tokens"`
}

type ClearTrustTokensArgs struct {
	IssuerOrigin string `json:"issuerOrigin"`
}

type ClearTrustTokensVal struct {
	DidDeleteTokens bool `json:"didDeleteTokens"`
}

type GetInterestGroupDetailsArgs struct {
	OwnerOrigin string `json:"ownerOrigin"`
	Name        string `json:"name"`
}

type GetInterestGroupDetailsVal struct {
	Details *InterestGroupDetails `json:"details"`
}

type SetInterestGroupTrackingArgs struct {
	Enable bool `json:"enable"`
}

type GetSharedStorageMetadataArgs struct {
	OwnerOrigin string `json:"ownerOrigin"`
}

type GetSharedStorageMetadataVal struct {
	Metadata *SharedStorageMetadata `json:"metadata"`
}

type GetSharedStorageEntriesArgs struct {
	OwnerOrigin string `json:"ownerOrigin"`
}

type GetSharedStorageEntriesVal struct {
	Entries []*SharedStorageEntry `json:"entries"`
}

type SetSharedStorageEntryArgs struct {
	OwnerOrigin     string `json:"ownerOrigin"`
	Key             string `json:"key"`
	Value           string `json:"value"`
	IgnoreIfPresent bool   `json:"ignoreIfPresent,omitempty"`
}

type DeleteSharedStorageEntryArgs struct {
	OwnerOrigin string `json:"ownerOrigin"`
	Key         string `json:"key"`
}

type ClearSharedStorageEntriesArgs struct {
	OwnerOrigin string `json:"ownerOrigin"`
}

type ResetSharedStorageBudgetArgs struct {
	OwnerOrigin string `json:"ownerOrigin"`
}

type SetSharedStorageTrackingArgs struct {
	Enable bool `json:"enable"`
}
