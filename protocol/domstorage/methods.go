package domstorage

import (
	"github.com/ecwid/control/protocol"
)

/*

 */
func Clear(c protocol.Caller, args ClearArgs) error {
	return c.Call("DOMStorage.clear", args, nil)
}

/*
	Disables storage tracking, prevents storage events from being sent to the client.
*/
func Disable(c protocol.Caller) error {
	return c.Call("DOMStorage.disable", nil, nil)
}

/*
	Enables storage tracking, storage events will now be delivered to the client.
*/
func Enable(c protocol.Caller) error {
	return c.Call("DOMStorage.enable", nil, nil)
}

/*

 */
func GetDOMStorageItems(c protocol.Caller, args GetDOMStorageItemsArgs) (*GetDOMStorageItemsVal, error) {
	var val = &GetDOMStorageItemsVal{}
	return val, c.Call("DOMStorage.getDOMStorageItems", args, val)
}

/*

 */
func RemoveDOMStorageItem(c protocol.Caller, args RemoveDOMStorageItemArgs) error {
	return c.Call("DOMStorage.removeDOMStorageItem", args, nil)
}

/*

 */
func SetDOMStorageItem(c protocol.Caller, args SetDOMStorageItemArgs) error {
	return c.Call("DOMStorage.setDOMStorageItem", args, nil)
}
