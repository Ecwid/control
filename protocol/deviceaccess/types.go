package deviceaccess

/*
Device request id.
*/
type RequestId string

/*
A device id.
*/
type DeviceId string

/*
Device information displayed in a user prompt to select a device.
*/
type PromptDevice struct {
	Id   DeviceId `json:"id"`
	Name string   `json:"name"`
}

type SelectPromptArgs struct {
	Id       RequestId `json:"id"`
	DeviceId DeviceId  `json:"deviceId"`
}

type CancelPromptArgs struct {
	Id RequestId `json:"id"`
}
