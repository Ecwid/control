package domstorage

/*

 */
type DomStorageItemAdded struct {
	StorageId *StorageId `json:"storageId"`
	Key       string     `json:"key"`
	NewValue  string     `json:"newValue"`
}

/*

 */
type DomStorageItemRemoved struct {
	StorageId *StorageId `json:"storageId"`
	Key       string     `json:"key"`
}

/*

 */
type DomStorageItemUpdated struct {
	StorageId *StorageId `json:"storageId"`
	Key       string     `json:"key"`
	OldValue  string     `json:"oldValue"`
	NewValue  string     `json:"newValue"`
}

/*

 */
type DomStorageItemsCleared struct {
	StorageId *StorageId `json:"storageId"`
}
