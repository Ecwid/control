package storage

/*
	A cache's contents have been modified.
*/
type CacheStorageContentUpdated struct {
	Origin    string `json:"origin"`
	CacheName string `json:"cacheName"`
}

/*
	A cache has been added/deleted.
*/
type CacheStorageListUpdated struct {
	Origin string `json:"origin"`
}

/*
	The origin's IndexedDB object store has been modified.
*/
type IndexedDBContentUpdated struct {
	Origin          string `json:"origin"`
	DatabaseName    string `json:"databaseName"`
	ObjectStoreName string `json:"objectStoreName"`
}

/*
	The origin's IndexedDB database list has been modified.
*/
type IndexedDBListUpdated struct {
	Origin string `json:"origin"`
}
