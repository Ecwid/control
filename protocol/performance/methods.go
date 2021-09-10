package performance

import (
	"github.com/ecwid/control/protocol"
)

/*
	Disable collecting and reporting metrics.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Performance.disable", nil, nil)
}

/*
	Enable collecting and reporting metrics.
*/
func Enable(c protocol.Caller, args EnableArgs) error {
	return c.Call("Performance.enable", args, nil)
}

/*
	Retrieve current values of run-time metrics.
*/
func GetMetrics(c protocol.Caller) (*GetMetricsVal, error) {
	var val = &GetMetricsVal{}
	return val, c.Call("Performance.getMetrics", nil, val)
}
