package indexeddb

import (
	"github.com/ecwid/control/protocol/runtime"
)

/*
	Database with an array of object stores.
*/
type DatabaseWithObjectStores struct {
	Name         string         `json:"name"`
	Version      float64        `json:"version"`
	ObjectStores []*ObjectStore `json:"objectStores"`
}

/*
	Object store.
*/
type ObjectStore struct {
	Name          string              `json:"name"`
	KeyPath       *KeyPath            `json:"keyPath"`
	AutoIncrement bool                `json:"autoIncrement"`
	Indexes       []*ObjectStoreIndex `json:"indexes"`
}

/*
	Object store index.
*/
type ObjectStoreIndex struct {
	Name       string   `json:"name"`
	KeyPath    *KeyPath `json:"keyPath"`
	Unique     bool     `json:"unique"`
	MultiEntry bool     `json:"multiEntry"`
}

/*
	Key.
*/
type Key struct {
	Type   string  `json:"type"`
	Number float64 `json:"number,omitempty"`
	String string  `json:"string,omitempty"`
	Date   float64 `json:"date,omitempty"`
	Array  []*Key  `json:"array,omitempty"`
}

/*
	Key range.
*/
type KeyRange struct {
	Lower     *Key `json:"lower,omitempty"`
	Upper     *Key `json:"upper,omitempty"`
	LowerOpen bool `json:"lowerOpen"`
	UpperOpen bool `json:"upperOpen"`
}

/*
	Data entry.
*/
type DataEntry struct {
	Key        *runtime.RemoteObject `json:"key"`
	PrimaryKey *runtime.RemoteObject `json:"primaryKey"`
	Value      *runtime.RemoteObject `json:"value"`
}

/*
	Key path.
*/
type KeyPath struct {
	Type   string   `json:"type"`
	String string   `json:"string,omitempty"`
	Array  []string `json:"array,omitempty"`
}

type ClearObjectStoreArgs struct {
	SecurityOrigin  string `json:"securityOrigin"`
	DatabaseName    string `json:"databaseName"`
	ObjectStoreName string `json:"objectStoreName"`
}

type DeleteDatabaseArgs struct {
	SecurityOrigin string `json:"securityOrigin"`
	DatabaseName   string `json:"databaseName"`
}

type DeleteObjectStoreEntriesArgs struct {
	SecurityOrigin  string    `json:"securityOrigin"`
	DatabaseName    string    `json:"databaseName"`
	ObjectStoreName string    `json:"objectStoreName"`
	KeyRange        *KeyRange `json:"keyRange"`
}

type RequestDataArgs struct {
	SecurityOrigin  string    `json:"securityOrigin"`
	DatabaseName    string    `json:"databaseName"`
	ObjectStoreName string    `json:"objectStoreName"`
	IndexName       string    `json:"indexName"`
	SkipCount       int       `json:"skipCount"`
	PageSize        int       `json:"pageSize"`
	KeyRange        *KeyRange `json:"keyRange,omitempty"`
}

type RequestDataVal struct {
	ObjectStoreDataEntries []*DataEntry `json:"objectStoreDataEntries"`
	HasMore                bool         `json:"hasMore"`
}

type GetMetadataArgs struct {
	SecurityOrigin  string `json:"securityOrigin"`
	DatabaseName    string `json:"databaseName"`
	ObjectStoreName string `json:"objectStoreName"`
}

type GetMetadataVal struct {
	EntriesCount      float64 `json:"entriesCount"`
	KeyGeneratorValue float64 `json:"keyGeneratorValue"`
}

type RequestDatabaseArgs struct {
	SecurityOrigin string `json:"securityOrigin"`
	DatabaseName   string `json:"databaseName"`
}

type RequestDatabaseVal struct {
	DatabaseWithObjectStores *DatabaseWithObjectStores `json:"databaseWithObjectStores"`
}

type RequestDatabaseNamesArgs struct {
	SecurityOrigin string `json:"securityOrigin"`
}

type RequestDatabaseNamesVal struct {
	DatabaseNames []string `json:"databaseNames"`
}
