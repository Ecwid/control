package common

/*
 */
type BrowserContextID string

/*
Rectangle.
*/
type Rect struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

/*
	Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints

Missing optional values will be filled in by the target with what it would normally use.
*/
type UserAgentMetadata struct {
	Brands          []*UserAgentBrandVersion `json:"brands,omitempty"`
	FullVersionList []*UserAgentBrandVersion `json:"fullVersionList,omitempty"`
	Platform        string                   `json:"platform"`
	PlatformVersion string                   `json:"platformVersion"`
	Architecture    string                   `json:"architecture"`
	Model           string                   `json:"model"`
	Mobile          bool                     `json:"mobile"`
	Bitness         string                   `json:"bitness,omitempty"`
	Wow64           bool                     `json:"wow64,omitempty"`
}

/*
Used to specify User Agent Cient Hints to emulate. See https://wicg.github.io/ua-client-hints
*/
type UserAgentBrandVersion struct {
	Brand   string `json:"brand"`
	Version string `json:"version"`
}

/*
UTC time in seconds, counted from January 1, 1970.
*/
type TimeSinceEpoch float64

/*
Unique frame identifier.
*/
type FrameId string
