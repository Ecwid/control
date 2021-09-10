package storage

import (
	"github.com/ecwid/control/protocol"
)

/*
	Clears storage for origin.
*/
func ClearDataForOrigin(c protocol.Caller, args ClearDataForOriginArgs) error {
	return c.Call("Storage.clearDataForOrigin", args, nil)
}

/*
	Returns all browser cookies.
*/
func GetCookies(c protocol.Caller, args GetCookiesArgs) (*GetCookiesVal, error) {
	var val = &GetCookiesVal{}
	return val, c.Call("Storage.getCookies", args, val)
}

/*
	Sets given cookies.
*/
func SetCookies(c protocol.Caller, args SetCookiesArgs) error {
	return c.Call("Storage.setCookies", args, nil)
}

/*
	Clears cookies.
*/
func ClearCookies(c protocol.Caller, args ClearCookiesArgs) error {
	return c.Call("Storage.clearCookies", args, nil)
}

/*
	Returns usage and quota in bytes.
*/
func GetUsageAndQuota(c protocol.Caller, args GetUsageAndQuotaArgs) (*GetUsageAndQuotaVal, error) {
	var val = &GetUsageAndQuotaVal{}
	return val, c.Call("Storage.getUsageAndQuota", args, val)
}

/*
	Override quota for the specified origin
*/
func OverrideQuotaForOrigin(c protocol.Caller, args OverrideQuotaForOriginArgs) error {
	return c.Call("Storage.overrideQuotaForOrigin", args, nil)
}

/*
	Registers origin to be notified when an update occurs to its cache storage list.
*/
func TrackCacheStorageForOrigin(c protocol.Caller, args TrackCacheStorageForOriginArgs) error {
	return c.Call("Storage.trackCacheStorageForOrigin", args, nil)
}

/*
	Registers origin to be notified when an update occurs to its IndexedDB.
*/
func TrackIndexedDBForOrigin(c protocol.Caller, args TrackIndexedDBForOriginArgs) error {
	return c.Call("Storage.trackIndexedDBForOrigin", args, nil)
}

/*
	Unregisters origin from receiving notifications for cache storage.
*/
func UntrackCacheStorageForOrigin(c protocol.Caller, args UntrackCacheStorageForOriginArgs) error {
	return c.Call("Storage.untrackCacheStorageForOrigin", args, nil)
}

/*
	Unregisters origin from receiving notifications for IndexedDB.
*/
func UntrackIndexedDBForOrigin(c protocol.Caller, args UntrackIndexedDBForOriginArgs) error {
	return c.Call("Storage.untrackIndexedDBForOrigin", args, nil)
}

/*
	Returns the number of stored Trust Tokens per issuer for the
current browsing context.
*/
func GetTrustTokens(c protocol.Caller) (*GetTrustTokensVal, error) {
	var val = &GetTrustTokensVal{}
	return val, c.Call("Storage.getTrustTokens", nil, val)
}

/*
	Removes all Trust Tokens issued by the provided issuerOrigin.
Leaves other stored data, including the issuer's Redemption Records, intact.
*/
func ClearTrustTokens(c protocol.Caller, args ClearTrustTokensArgs) (*ClearTrustTokensVal, error) {
	var val = &ClearTrustTokensVal{}
	return val, c.Call("Storage.clearTrustTokens", args, val)
}
