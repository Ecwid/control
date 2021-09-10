package storage

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/network"
)

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

type ClearDataForOriginArgs struct {
	Origin       string `json:"origin"`
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

type TrackIndexedDBForOriginArgs struct {
	Origin string `json:"origin"`
}

type UntrackCacheStorageForOriginArgs struct {
	Origin string `json:"origin"`
}

type UntrackIndexedDBForOriginArgs struct {
	Origin string `json:"origin"`
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
