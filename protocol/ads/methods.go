package ads

import (
	"github.com/ecwid/control/protocol"
)

/*
Retrieves ad metrics for the current page.
*/
func GetAdMetrics(c protocol.Caller) (*GetAdMetricsVal, error) {
	var val = &GetAdMetricsVal{}
	return val, c.Call("Ads.getAdMetrics", nil, val)
}
