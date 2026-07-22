package extensions

import (
	"github.com/ecwid/control/protocol"
)

/*
Runs an extension default action.
*/
func TriggerAction(c protocol.Caller, args TriggerActionArgs) error {
	return c.Call("Extensions.triggerAction", args, nil)
}

/*
	Installs an unpacked extension from the filesystem similar to

--load-extension CLI flags. Returns extension ID once the extension
has been installed.
*/
func LoadUnpacked(c protocol.Caller, args LoadUnpackedArgs) (*LoadUnpackedVal, error) {
	var val = &LoadUnpackedVal{}
	return val, c.Call("Extensions.loadUnpacked", args, val)
}

/*
Gets a list of all unpacked extensions.
*/
func GetExtensions(c protocol.Caller) (*GetExtensionsVal, error) {
	var val = &GetExtensionsVal{}
	return val, c.Call("Extensions.getExtensions", nil, val)
}

/*
Uninstalls an unpacked extension (others not supported) from the profile.
*/
func Uninstall(c protocol.Caller, args UninstallArgs) error {
	return c.Call("Extensions.uninstall", args, nil)
}

/*
	Gets data from extension storage in the given `storageArea`. If `keys` is

specified, these are used to filter the result.
*/
func GetStorageItems(c protocol.Caller, args GetStorageItemsArgs) (*GetStorageItemsVal, error) {
	var val = &GetStorageItemsVal{}
	return val, c.Call("Extensions.getStorageItems", args, val)
}

/*
Removes `keys` from extension storage in the given `storageArea`.
*/
func RemoveStorageItems(c protocol.Caller, args RemoveStorageItemsArgs) error {
	return c.Call("Extensions.removeStorageItems", args, nil)
}

/*
Clears extension storage in the given `storageArea`.
*/
func ClearStorageItems(c protocol.Caller, args ClearStorageItemsArgs) error {
	return c.Call("Extensions.clearStorageItems", args, nil)
}

/*
	Sets `values` in extension storage in the given `storageArea`. The provided `values`

will be merged with existing values in the storage area.
*/
func SetStorageItems(c protocol.Caller, args SetStorageItemsArgs) error {
	return c.Call("Extensions.setStorageItems", args, nil)
}
