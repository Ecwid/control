package css

/*
	Fires whenever a web font is updated.  A non-empty font parameter indicates a successfully loaded
web font
*/
type FontsUpdated struct {
	Font *FontFace `json:"font,omitempty"`
}

/*
	Fires whenever a MediaQuery result changes (for example, after a browser window has been
resized.) The current implementation considers only viewport-dependent media features.
*/
type MediaQueryResultChanged interface{}

/*
	Fired whenever an active document stylesheet is added.
*/
type StyleSheetAdded struct {
	Header *CSSStyleSheetHeader `json:"header"`
}

/*
	Fired whenever a stylesheet is changed as a result of the client operation.
*/
type StyleSheetChanged struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
}

/*
	Fired whenever an active document stylesheet is removed.
*/
type StyleSheetRemoved struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
}
