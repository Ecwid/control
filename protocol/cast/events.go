package cast

/*
	This is fired whenever the list of available sinks changes. A sink is a

device or a software surface that you can cast to.
*/
type SinksUpdated struct {
	Sinks []*Sink `json:"sinks"`
}

/*
	This is fired whenever the outstanding issue/error message changes.

|issueMessage| is empty if there is no issue.
*/
type IssueUpdated struct {
	IssueMessage string `json:"issueMessage"`
}
