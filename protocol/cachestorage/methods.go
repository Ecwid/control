package cachestorage

import (
	"github.com/ecwid/control/protocol"
)

/*
Deletes a cache.
*/
func DeleteCache(c protocol.Caller, args DeleteCacheArgs) error {
	return c.Call("CacheStorage.deleteCache", args, nil)
}

/*
Deletes a cache entry.
*/
func DeleteEntry(c protocol.Caller, args DeleteEntryArgs) error {
	return c.Call("CacheStorage.deleteEntry", args, nil)
}

/*
Requests cache names.
*/
func RequestCacheNames(c protocol.Caller, args RequestCacheNamesArgs) (*RequestCacheNamesVal, error) {
	var val = &RequestCacheNamesVal{}
	return val, c.Call("CacheStorage.requestCacheNames", args, val)
}

/*
Fetches cache entry.
*/
func RequestCachedResponse(c protocol.Caller, args RequestCachedResponseArgs) (*RequestCachedResponseVal, error) {
	var val = &RequestCachedResponseVal{}
	return val, c.Call("CacheStorage.requestCachedResponse", args, val)
}

/*
Requests data from cache.
*/
func RequestEntries(c protocol.Caller, args RequestEntriesArgs) (*RequestEntriesVal, error) {
	var val = &RequestEntriesVal{}
	return val, c.Call("CacheStorage.requestEntries", args, val)
}
