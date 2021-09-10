package log

/*
	Issued when new message was logged.
*/
type EntryAdded struct {
	Entry *LogEntry `json:"entry"`
}
