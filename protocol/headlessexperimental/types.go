package headlessexperimental

/*
	Encoding options for a screenshot.
*/
type ScreenshotParams struct {
	Format  string `json:"format,omitempty"`
	Quality int    `json:"quality,omitempty"`
}

type BeginFrameArgs struct {
	FrameTimeTicks   float64           `json:"frameTimeTicks,omitempty"`
	Interval         float64           `json:"interval,omitempty"`
	NoDisplayUpdates bool              `json:"noDisplayUpdates,omitempty"`
	Screenshot       *ScreenshotParams `json:"screenshot,omitempty"`
}

type BeginFrameVal struct {
	HasDamage      bool   `json:"hasDamage"`
	ScreenshotData []byte `json:"screenshotData,omitempty"`
}
