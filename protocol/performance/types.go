package performance

/*
	Run-time execution metric.
*/
type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type EnableArgs struct {
	TimeDomain string `json:"timeDomain,omitempty"`
}

type GetMetricsVal struct {
	Metrics []*Metric `json:"metrics"`
}
