package tethering

/*
	Informs that port was successfully bound and got a specified connection id.
*/
type Accepted struct {
	Port         int    `json:"port"`
	ConnectionId string `json:"connectionId"`
}
