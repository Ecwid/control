package css

import (
	"github.com/ecwid/control/protocol"
)

/*
	Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the
position specified by `location`.
*/
func AddRule(c protocol.Caller, args AddRuleArgs) (*AddRuleVal, error) {
	var val = &AddRuleVal{}
	return val, c.Call("CSS.addRule", args, val)
}

/*
	Returns all class names from specified stylesheet.
*/
func CollectClassNames(c protocol.Caller, args CollectClassNamesArgs) (*CollectClassNamesVal, error) {
	var val = &CollectClassNamesVal{}
	return val, c.Call("CSS.collectClassNames", args, val)
}

/*
	Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
*/
func CreateStyleSheet(c protocol.Caller, args CreateStyleSheetArgs) (*CreateStyleSheetVal, error) {
	var val = &CreateStyleSheetVal{}
	return val, c.Call("CSS.createStyleSheet", args, val)
}

/*
	Disables the CSS agent for the given page.
*/
func Disable(c protocol.Caller) error {
	return c.Call("CSS.disable", nil, nil)
}

/*
	Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been
enabled until the result of this command is received.
*/
func Enable(c protocol.Caller) error {
	return c.Call("CSS.enable", nil, nil)
}

/*
	Ensures that the given node will have specified pseudo-classes whenever its style is computed by
the browser.
*/
func ForcePseudoState(c protocol.Caller, args ForcePseudoStateArgs) error {
	return c.Call("CSS.forcePseudoState", args, nil)
}

/*

 */
func GetBackgroundColors(c protocol.Caller, args GetBackgroundColorsArgs) (*GetBackgroundColorsVal, error) {
	var val = &GetBackgroundColorsVal{}
	return val, c.Call("CSS.getBackgroundColors", args, val)
}

/*
	Returns the computed style for a DOM node identified by `nodeId`.
*/
func GetComputedStyleForNode(c protocol.Caller, args GetComputedStyleForNodeArgs) (*GetComputedStyleForNodeVal, error) {
	var val = &GetComputedStyleForNodeVal{}
	return val, c.Call("CSS.getComputedStyleForNode", args, val)
}

/*
	Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM
attributes) for a DOM node identified by `nodeId`.
*/
func GetInlineStylesForNode(c protocol.Caller, args GetInlineStylesForNodeArgs) (*GetInlineStylesForNodeVal, error) {
	var val = &GetInlineStylesForNodeVal{}
	return val, c.Call("CSS.getInlineStylesForNode", args, val)
}

/*
	Returns requested styles for a DOM node identified by `nodeId`.
*/
func GetMatchedStylesForNode(c protocol.Caller, args GetMatchedStylesForNodeArgs) (*GetMatchedStylesForNodeVal, error) {
	var val = &GetMatchedStylesForNodeVal{}
	return val, c.Call("CSS.getMatchedStylesForNode", args, val)
}

/*
	Returns all media queries parsed by the rendering engine.
*/
func GetMediaQueries(c protocol.Caller) (*GetMediaQueriesVal, error) {
	var val = &GetMediaQueriesVal{}
	return val, c.Call("CSS.getMediaQueries", nil, val)
}

/*
	Requests information about platform fonts which we used to render child TextNodes in the given
node.
*/
func GetPlatformFontsForNode(c protocol.Caller, args GetPlatformFontsForNodeArgs) (*GetPlatformFontsForNodeVal, error) {
	var val = &GetPlatformFontsForNodeVal{}
	return val, c.Call("CSS.getPlatformFontsForNode", args, val)
}

/*
	Returns the current textual content for a stylesheet.
*/
func GetStyleSheetText(c protocol.Caller, args GetStyleSheetTextArgs) (*GetStyleSheetTextVal, error) {
	var val = &GetStyleSheetTextVal{}
	return val, c.Call("CSS.getStyleSheetText", args, val)
}

/*
	Starts tracking the given computed styles for updates. The specified array of properties
replaces the one previously specified. Pass empty array to disable tracking.
Use takeComputedStyleUpdates to retrieve the list of nodes that had properties modified.
The changes to computed style properties are only tracked for nodes pushed to the front-end
by the DOM agent. If no changes to the tracked properties occur after the node has been pushed
to the front-end, no updates will be issued for the node.
*/
func TrackComputedStyleUpdates(c protocol.Caller, args TrackComputedStyleUpdatesArgs) error {
	return c.Call("CSS.trackComputedStyleUpdates", args, nil)
}

/*
	Polls the next batch of computed style updates.
*/
func TakeComputedStyleUpdates(c protocol.Caller) (*TakeComputedStyleUpdatesVal, error) {
	var val = &TakeComputedStyleUpdatesVal{}
	return val, c.Call("CSS.takeComputedStyleUpdates", nil, val)
}

/*
	Find a rule with the given active property for the given node and set the new value for this
property
*/
func SetEffectivePropertyValueForNode(c protocol.Caller, args SetEffectivePropertyValueForNodeArgs) error {
	return c.Call("CSS.setEffectivePropertyValueForNode", args, nil)
}

/*
	Modifies the keyframe rule key text.
*/
func SetKeyframeKey(c protocol.Caller, args SetKeyframeKeyArgs) (*SetKeyframeKeyVal, error) {
	var val = &SetKeyframeKeyVal{}
	return val, c.Call("CSS.setKeyframeKey", args, val)
}

/*
	Modifies the rule selector.
*/
func SetMediaText(c protocol.Caller, args SetMediaTextArgs) (*SetMediaTextVal, error) {
	var val = &SetMediaTextVal{}
	return val, c.Call("CSS.setMediaText", args, val)
}

/*
	Modifies the rule selector.
*/
func SetRuleSelector(c protocol.Caller, args SetRuleSelectorArgs) (*SetRuleSelectorVal, error) {
	var val = &SetRuleSelectorVal{}
	return val, c.Call("CSS.setRuleSelector", args, val)
}

/*
	Sets the new stylesheet text.
*/
func SetStyleSheetText(c protocol.Caller, args SetStyleSheetTextArgs) (*SetStyleSheetTextVal, error) {
	var val = &SetStyleSheetTextVal{}
	return val, c.Call("CSS.setStyleSheetText", args, val)
}

/*
	Applies specified style edits one after another in the given order.
*/
func SetStyleTexts(c protocol.Caller, args SetStyleTextsArgs) (*SetStyleTextsVal, error) {
	var val = &SetStyleTextsVal{}
	return val, c.Call("CSS.setStyleTexts", args, val)
}

/*
	Enables the selector recording.
*/
func StartRuleUsageTracking(c protocol.Caller) error {
	return c.Call("CSS.startRuleUsageTracking", nil, nil)
}

/*
	Stop tracking rule usage and return the list of rules that were used since last call to
`takeCoverageDelta` (or since start of coverage instrumentation)
*/
func StopRuleUsageTracking(c protocol.Caller) (*StopRuleUsageTrackingVal, error) {
	var val = &StopRuleUsageTrackingVal{}
	return val, c.Call("CSS.stopRuleUsageTracking", nil, val)
}

/*
	Obtain list of rules that became used since last call to this method (or since start of coverage
instrumentation)
*/
func TakeCoverageDelta(c protocol.Caller) (*TakeCoverageDeltaVal, error) {
	var val = &TakeCoverageDeltaVal{}
	return val, c.Call("CSS.takeCoverageDelta", nil, val)
}

/*
	Enables/disables rendering of local CSS fonts (enabled by default).
*/
func SetLocalFontsEnabled(c protocol.Caller, args SetLocalFontsEnabledArgs) error {
	return c.Call("CSS.setLocalFontsEnabled", args, nil)
}
