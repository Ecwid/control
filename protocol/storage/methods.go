package storage

import (
	"github.com/ecwid/control/protocol"
)

/*
Returns a storage key given a frame id.
*/
func GetStorageKeyForFrame(c protocol.Caller, args GetStorageKeyForFrameArgs) (*GetStorageKeyForFrameVal, error) {
	var val = &GetStorageKeyForFrameVal{}
	return val, c.Call("Storage.getStorageKeyForFrame", args, val)
}

/*
Clears storage for origin.
*/
func ClearDataForOrigin(c protocol.Caller, args ClearDataForOriginArgs) error {
	return c.Call("Storage.clearDataForOrigin", args, nil)
}

/*
Clears storage for storage key.
*/
func ClearDataForStorageKey(c protocol.Caller, args ClearDataForStorageKeyArgs) error {
	return c.Call("Storage.clearDataForStorageKey", args, nil)
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
Registers storage key to be notified when an update occurs to its cache storage list.
*/
func TrackCacheStorageForStorageKey(c protocol.Caller, args TrackCacheStorageForStorageKeyArgs) error {
	return c.Call("Storage.trackCacheStorageForStorageKey", args, nil)
}

/*
Registers origin to be notified when an update occurs to its IndexedDB.
*/
func TrackIndexedDBForOrigin(c protocol.Caller, args TrackIndexedDBForOriginArgs) error {
	return c.Call("Storage.trackIndexedDBForOrigin", args, nil)
}

/*
Registers storage key to be notified when an update occurs to its IndexedDB.
*/
func TrackIndexedDBForStorageKey(c protocol.Caller, args TrackIndexedDBForStorageKeyArgs) error {
	return c.Call("Storage.trackIndexedDBForStorageKey", args, nil)
}

/*
Unregisters origin from receiving notifications for cache storage.
*/
func UntrackCacheStorageForOrigin(c protocol.Caller, args UntrackCacheStorageForOriginArgs) error {
	return c.Call("Storage.untrackCacheStorageForOrigin", args, nil)
}

/*
Unregisters storage key from receiving notifications for cache storage.
*/
func UntrackCacheStorageForStorageKey(c protocol.Caller, args UntrackCacheStorageForStorageKeyArgs) error {
	return c.Call("Storage.untrackCacheStorageForStorageKey", args, nil)
}

/*
Unregisters origin from receiving notifications for IndexedDB.
*/
func UntrackIndexedDBForOrigin(c protocol.Caller, args UntrackIndexedDBForOriginArgs) error {
	return c.Call("Storage.untrackIndexedDBForOrigin", args, nil)
}

/*
Unregisters storage key from receiving notifications for IndexedDB.
*/
func UntrackIndexedDBForStorageKey(c protocol.Caller, args UntrackIndexedDBForStorageKeyArgs) error {
	return c.Call("Storage.untrackIndexedDBForStorageKey", args, nil)
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

/*
Gets details for a named interest group.
*/
func GetInterestGroupDetails(c protocol.Caller, args GetInterestGroupDetailsArgs) (*GetInterestGroupDetailsVal, error) {
	var val = &GetInterestGroupDetailsVal{}
	return val, c.Call("Storage.getInterestGroupDetails", args, val)
}

/*
Enables/Disables issuing of interestGroupAccessed events.
*/
func SetInterestGroupTracking(c protocol.Caller, args SetInterestGroupTrackingArgs) error {
	return c.Call("Storage.setInterestGroupTracking", args, nil)
}

/*
Gets metadata for an origin's shared storage.
*/
func GetSharedStorageMetadata(c protocol.Caller, args GetSharedStorageMetadataArgs) (*GetSharedStorageMetadataVal, error) {
	var val = &GetSharedStorageMetadataVal{}
	return val, c.Call("Storage.getSharedStorageMetadata", args, val)
}

/*
Gets the entries in an given origin's shared storage.
*/
func GetSharedStorageEntries(c protocol.Caller, args GetSharedStorageEntriesArgs) (*GetSharedStorageEntriesVal, error) {
	var val = &GetSharedStorageEntriesVal{}
	return val, c.Call("Storage.getSharedStorageEntries", args, val)
}

/*
Sets entry with `key` and `value` for a given origin's shared storage.
*/
func SetSharedStorageEntry(c protocol.Caller, args SetSharedStorageEntryArgs) error {
	return c.Call("Storage.setSharedStorageEntry", args, nil)
}

/*
Deletes entry for `key` (if it exists) for a given origin's shared storage.
*/
func DeleteSharedStorageEntry(c protocol.Caller, args DeleteSharedStorageEntryArgs) error {
	return c.Call("Storage.deleteSharedStorageEntry", args, nil)
}

/*
Clears all entries for a given origin's shared storage.
*/
func ClearSharedStorageEntries(c protocol.Caller, args ClearSharedStorageEntriesArgs) error {
	return c.Call("Storage.clearSharedStorageEntries", args, nil)
}

/*
Resets the budget for `ownerOrigin` by clearing all budget withdrawals.
*/
func ResetSharedStorageBudget(c protocol.Caller, args ResetSharedStorageBudgetArgs) error {
	return c.Call("Storage.resetSharedStorageBudget", args, nil)
}

/*
Enables/disables issuing of sharedStorageAccessed events.
*/
func SetSharedStorageTracking(c protocol.Caller, args SetSharedStorageTrackingArgs) error {
	return c.Call("Storage.setSharedStorageTracking", args, nil)
}
