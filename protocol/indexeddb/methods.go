package indexeddb

import (
	"github.com/ecwid/control/protocol"
)

/*
Clears all entries from an object store.
*/
func ClearObjectStore(c protocol.Caller, args ClearObjectStoreArgs) error {
	return c.Call("IndexedDB.clearObjectStore", args, nil)
}

/*
Deletes a database.
*/
func DeleteDatabase(c protocol.Caller, args DeleteDatabaseArgs) error {
	return c.Call("IndexedDB.deleteDatabase", args, nil)
}

/*
Delete a range of entries from an object store
*/
func DeleteObjectStoreEntries(c protocol.Caller, args DeleteObjectStoreEntriesArgs) error {
	return c.Call("IndexedDB.deleteObjectStoreEntries", args, nil)
}

/*
Disables events from backend.
*/
func Disable(c protocol.Caller) error {
	return c.Call("IndexedDB.disable", nil, nil)
}

/*
Enables events from backend.
*/
func Enable(c protocol.Caller) error {
	return c.Call("IndexedDB.enable", nil, nil)
}

/*
Requests data from object store or index.
*/
func RequestData(c protocol.Caller, args RequestDataArgs) (*RequestDataVal, error) {
	var val = &RequestDataVal{}
	return val, c.Call("IndexedDB.requestData", args, val)
}

/*
Gets metadata of an object store
*/
func GetMetadata(c protocol.Caller, args GetMetadataArgs) (*GetMetadataVal, error) {
	var val = &GetMetadataVal{}
	return val, c.Call("IndexedDB.getMetadata", args, val)
}

/*
Requests database with given name in given frame.
*/
func RequestDatabase(c protocol.Caller, args RequestDatabaseArgs) (*RequestDatabaseVal, error) {
	var val = &RequestDatabaseVal{}
	return val, c.Call("IndexedDB.requestDatabase", args, val)
}

/*
Requests database names for given security origin.
*/
func RequestDatabaseNames(c protocol.Caller, args RequestDatabaseNamesArgs) (*RequestDatabaseNamesVal, error) {
	var val = &RequestDatabaseNamesVal{}
	return val, c.Call("IndexedDB.requestDatabaseNames", args, val)
}
