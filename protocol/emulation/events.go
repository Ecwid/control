package emulation

/*
Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
*/
type VirtualTimeBudgetExpired any

/*
	Fired when a page calls screen.orientation.lock() or screen.orientation.unlock()

while device emulation is enabled. This allows the DevTools frontend to update the
emulated device orientation accordingly.
*/
type ScreenOrientationLockChanged struct {
	Locked      bool               `json:"locked"`
	Orientation *ScreenOrientation `json:"orientation,omitempty"`
}
