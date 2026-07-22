package ads

/*
Ad metrics for a page.
*/
type AdMetrics struct {
	ViewportAdDensityByArea        int     `json:"viewportAdDensityByArea"`
	AverageViewportAdDensityByArea float64 `json:"averageViewportAdDensityByArea"`
	ViewportAdCount                int     `json:"viewportAdCount"`
	AverageViewportAdCount         float64 `json:"averageViewportAdCount"`
	TotalAdCpuTime                 float64 `json:"totalAdCpuTime"`
	TotalAdNetworkBytes            float64 `json:"totalAdNetworkBytes"`
}

type GetAdMetricsVal struct {
	Metrics *AdMetrics `json:"metrics"`
}
