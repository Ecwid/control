package heapprofiler

/*

 */
type AddHeapSnapshotChunk struct {
	Chunk string `json:"chunk"`
}

/*
	If heap objects tracking has been started then backend may send update for one or more fragments
*/
type HeapStatsUpdate struct {
	StatsUpdate []int `json:"statsUpdate"`
}

/*
	If heap objects tracking has been started then backend regularly sends a current value for last
seen object id and corresponding timestamp. If the were changes in the heap since last event
then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
*/
type LastSeenObjectId struct {
	LastSeenObjectId int     `json:"lastSeenObjectId"`
	Timestamp        float64 `json:"timestamp"`
}

/*

 */
type ReportHeapSnapshotProgress struct {
	Done     int  `json:"done"`
	Total    int  `json:"total"`
	Finished bool `json:"finished,omitempty"`
}

/*

 */
type ResetProfiles interface{}
