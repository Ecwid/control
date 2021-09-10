package performance

/*
	Current values of the metrics.
*/
type Metrics struct {
	Metrics []*Metric `json:"metrics"`
	Title   string    `json:"title"`
}
