package deviceaccess

/*
	A device request opened a user prompt to select a device. Respond with the

selectPrompt or cancelPrompt command.
*/
type DeviceRequestPrompted struct {
	Id      RequestId       `json:"id"`
	Devices []*PromptDevice `json:"devices"`
}
