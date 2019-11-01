package har

import (
	"net/url"
	"time"
)

func parseHeaders(headers map[string]interface{}) []*NVP {
	h := make([]*NVP, 0)
	for k, v := range headers {
		h = append(h, &NVP{
			Name:  k,
			Value: v.(string),
		})
	}
	return h
}

func parseURL(strURL string) ([]*NVP, error) {
	nvps := make([]*NVP, 0)
	reqURL, err := url.Parse(strURL)
	if err != nil {
		return nil, err
	}
	for k, v := range reqURL.Query() {
		for _, value := range v {
			nvps = append(nvps, &NVP{
				Name:  k,
				Value: value,
			})
		}
	}
	return nvps, nil
}

func epoch(epoch float64) time.Time {
	return time.Unix(0, int64(epoch*1000)*int64(time.Millisecond))
}
