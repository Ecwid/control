package domstorage

/*
	DOM Storage identifier.
*/
type StorageId struct {
	SecurityOrigin string `json:"securityOrigin"`
	IsLocalStorage bool   `json:"isLocalStorage"`
}

/*
	DOM Storage item.
*/
type Item []string

type ClearArgs struct {
	StorageId *StorageId `json:"storageId"`
}

type GetDOMStorageItemsArgs struct {
	StorageId *StorageId `json:"storageId"`
}

type GetDOMStorageItemsVal struct {
	Entries []Item `json:"entries"`
}

type RemoveDOMStorageItemArgs struct {
	StorageId *StorageId `json:"storageId"`
	Key       string     `json:"key"`
}

type SetDOMStorageItemArgs struct {
	StorageId *StorageId `json:"storageId"`
	Key       string     `json:"key"`
	Value     string     `json:"value"`
}
