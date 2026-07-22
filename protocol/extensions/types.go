package extensions

/*
Storage areas.
*/
type StorageArea string

/*
Detailed information about an extension.
*/
type ExtensionInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
	Enabled bool   `json:"enabled"`
}

type TriggerActionArgs struct {
	Id       string `json:"id"`
	TargetId string `json:"targetId"`
}

type LoadUnpackedArgs struct {
	Path              string `json:"path"`
	EnableInIncognito bool   `json:"enableInIncognito,omitempty"`
}

type LoadUnpackedVal struct {
	Id string `json:"id"`
}

type GetExtensionsVal struct {
	Extensions []*ExtensionInfo `json:"extensions"`
}

type UninstallArgs struct {
	Id string `json:"id"`
}

type GetStorageItemsArgs struct {
	Id          string      `json:"id"`
	StorageArea StorageArea `json:"storageArea"`
	Keys        []string    `json:"keys,omitempty"`
}

type GetStorageItemsVal struct {
	Data any `json:"data"`
}

type RemoveStorageItemsArgs struct {
	Id          string      `json:"id"`
	StorageArea StorageArea `json:"storageArea"`
	Keys        []string    `json:"keys"`
}

type ClearStorageItemsArgs struct {
	Id          string      `json:"id"`
	StorageArea StorageArea `json:"storageArea"`
}

type SetStorageItemsArgs struct {
	Id          string      `json:"id"`
	StorageArea StorageArea `json:"storageArea"`
	Values      any         `json:"values"`
}
