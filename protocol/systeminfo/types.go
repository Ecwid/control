package systeminfo

/*
	Describes a single graphics processor (GPU).
*/
type GPUDevice struct {
	VendorId      float64 `json:"vendorId"`
	DeviceId      float64 `json:"deviceId"`
	SubSysId      float64 `json:"subSysId,omitempty"`
	Revision      float64 `json:"revision,omitempty"`
	VendorString  string  `json:"vendorString"`
	DeviceString  string  `json:"deviceString"`
	DriverVendor  string  `json:"driverVendor"`
	DriverVersion string  `json:"driverVersion"`
}

/*
	Describes the width and height dimensions of an entity.
*/
type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

/*
	Describes a supported video decoding profile with its associated minimum and
maximum resolutions.
*/
type VideoDecodeAcceleratorCapability struct {
	Profile       string `json:"profile"`
	MaxResolution *Size  `json:"maxResolution"`
	MinResolution *Size  `json:"minResolution"`
}

/*
	Describes a supported video encoding profile with its associated maximum
resolution and maximum framerate.
*/
type VideoEncodeAcceleratorCapability struct {
	Profile                 string `json:"profile"`
	MaxResolution           *Size  `json:"maxResolution"`
	MaxFramerateNumerator   int    `json:"maxFramerateNumerator"`
	MaxFramerateDenominator int    `json:"maxFramerateDenominator"`
}

/*
	YUV subsampling type of the pixels of a given image.
*/
type SubsamplingFormat string

/*
	Image format of a given image.
*/
type ImageType string

/*
	Describes a supported image decoding profile with its associated minimum and
maximum resolutions and subsampling.
*/
type ImageDecodeAcceleratorCapability struct {
	ImageType     ImageType           `json:"imageType"`
	MaxDimensions *Size               `json:"maxDimensions"`
	MinDimensions *Size               `json:"minDimensions"`
	Subsamplings  []SubsamplingFormat `json:"subsamplings"`
}

/*
	Provides information about the GPU(s) on the system.
*/
type GPUInfo struct {
	Devices              []*GPUDevice                        `json:"devices"`
	AuxAttributes        interface{}                         `json:"auxAttributes,omitempty"`
	FeatureStatus        interface{}                         `json:"featureStatus,omitempty"`
	DriverBugWorkarounds []string                            `json:"driverBugWorkarounds"`
	VideoDecoding        []*VideoDecodeAcceleratorCapability `json:"videoDecoding"`
	VideoEncoding        []*VideoEncodeAcceleratorCapability `json:"videoEncoding"`
	ImageDecoding        []*ImageDecodeAcceleratorCapability `json:"imageDecoding"`
}

/*
	Represents process info.
*/
type ProcessInfo struct {
	Type    string  `json:"type"`
	Id      int     `json:"id"`
	CpuTime float64 `json:"cpuTime"`
}

type GetInfoVal struct {
	Gpu          *GPUInfo `json:"gpu"`
	ModelName    string   `json:"modelName"`
	ModelVersion string   `json:"modelVersion"`
	CommandLine  string   `json:"commandLine"`
}

type GetProcessInfoVal struct {
	ProcessInfo []*ProcessInfo `json:"processInfo"`
}
