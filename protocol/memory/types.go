package memory

/*
Memory pressure level.
*/
type PressureLevel string

/*
Heap profile sample.
*/
type SamplingProfileNode struct {
	Size  float64  `json:"size"`
	Total float64  `json:"total"`
	Stack []string `json:"stack"`
}

/*
Array of heap profile samples.
*/
type SamplingProfile struct {
	Samples []*SamplingProfileNode `json:"samples"`
	Modules []*Module              `json:"modules"`
}

/*
Executable module information
*/
type Module struct {
	Name        string  `json:"name"`
	Uuid        string  `json:"uuid"`
	BaseAddress string  `json:"baseAddress"`
	Size        float64 `json:"size"`
}

type GetDOMCountersVal struct {
	Documents        int `json:"documents"`
	Nodes            int `json:"nodes"`
	JsEventListeners int `json:"jsEventListeners"`
}

type SetPressureNotificationsSuppressedArgs struct {
	Suppressed bool `json:"suppressed"`
}

type SimulatePressureNotificationArgs struct {
	Level PressureLevel `json:"level"`
}

type StartSamplingArgs struct {
	SamplingInterval   int  `json:"samplingInterval,omitempty"`
	SuppressRandomness bool `json:"suppressRandomness,omitempty"`
}

type GetAllTimeSamplingProfileVal struct {
	Profile *SamplingProfile `json:"profile"`
}

type GetBrowserSamplingProfileVal struct {
	Profile *SamplingProfile `json:"profile"`
}

type GetSamplingProfileVal struct {
	Profile *SamplingProfile `json:"profile"`
}
