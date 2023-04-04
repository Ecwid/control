package input

/*
	Emitted only when `Input.setInterceptDrags` is enabled. Use this data with `Input.dispatchDragEvent` to

restore normal drag and drop behavior.
*/
type DragIntercepted struct {
	Data *DragData `json:"data"`
}
