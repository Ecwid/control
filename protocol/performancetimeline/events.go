package performancetimeline

/*
	Sent when a performance timeline event is added. See reportPerformanceTimeline method.
*/
type TimelineEventAdded struct {
	Event *TimelineEvent `json:"event"`
}
