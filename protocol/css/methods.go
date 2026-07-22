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
Ensures that the given node is in its starting-style state.
*/
func ForceStartingStyle(c protocol.Caller, args ForceStartingStyleArgs) error {
	return c.Call("CSS.forceStartingStyle", args, nil)
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
	Resolve the specified values in the context of the provided element.

For example, a value of '1em' is evaluated according to the computed
'font-size' of the element and a value 'calc(1px + 2px)' will be
resolved to '3px'.
If the `propertyName` was specified the `values` are resolved as if
they were property's declaration. If a value cannot be parsed according
to the provided property syntax, the value is parsed using combined
syntax as if null `propertyName` was provided. If the value cannot be
resolved even then, return the provided value without any changes.
Note: this function currently does not resolve CSS random() function,
it returns unmodified random() function parts.`
*/
func ResolveValues(c protocol.Caller, args ResolveValuesArgs) (*ResolveValuesVal, error) {
	var val = &ResolveValuesVal{}
	return val, c.Call("CSS.resolveValues", args, val)
}

/*
 */
func GetLonghandProperties(c protocol.Caller, args GetLonghandPropertiesArgs) (*GetLonghandPropertiesVal, error) {
	var val = &GetLonghandPropertiesVal{}
	return val, c.Call("CSS.getLonghandProperties", args, val)
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
	Returns the styles coming from animations & transitions

including the animation & transition styles coming from inheritance chain.
*/
func GetAnimatedStylesForNode(c protocol.Caller, args GetAnimatedStylesForNodeArgs) (*GetAnimatedStylesForNodeVal, error) {
	var val = &GetAnimatedStylesForNodeVal{}
	return val, c.Call("CSS.getAnimatedStylesForNode", args, val)
}

/*
Returns requested styles for a DOM node identified by `nodeId`.
*/
func GetMatchedStylesForNode(c protocol.Caller, args GetMatchedStylesForNodeArgs) (*GetMatchedStylesForNodeVal, error) {
	var val = &GetMatchedStylesForNodeVal{}
	return val, c.Call("CSS.getMatchedStylesForNode", args, val)
}

/*
Returns the values of the default UA-defined environment variables used in env()
*/
func GetEnvironmentVariables(c protocol.Caller) (*GetEnvironmentVariablesVal, error) {
	var val = &GetEnvironmentVariablesVal{}
	return val, c.Call("CSS.getEnvironmentVariables", nil, val)
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
	Returns all layers parsed by the rendering engine for the tree scope of a node.

Given a DOM element identified by nodeId, getLayersForNode returns the root
layer for the nearest ancestor document or shadow root. The layer root contains
the full layer tree for the tree scope and their ordering.
*/
func GetLayersForNode(c protocol.Caller, args GetLayersForNodeArgs) (*GetLayersForNodeVal, error) {
	var val = &GetLayersForNodeVal{}
	return val, c.Call("CSS.getLayersForNode", args, val)
}

/*
	Given a CSS selector text and a style sheet ID, getLocationForSelector

returns an array of locations of the CSS selector in the style sheet.
*/
func GetLocationForSelector(c protocol.Caller, args GetLocationForSelectorArgs) (*GetLocationForSelectorVal, error) {
	var val = &GetLocationForSelectorVal{}
	return val, c.Call("CSS.getLocationForSelector", args, val)
}

/*
	Starts tracking the given node for the computed style updates

and whenever the computed style is updated for node, it queues
a `computedStyleUpdated` event with throttling.
There can only be 1 node tracked for computed style updates
so passing a new node id removes tracking from the previous node.
Pass `undefined` to disable tracking.
*/
func TrackComputedStyleUpdatesForNode(c protocol.Caller, args TrackComputedStyleUpdatesForNodeArgs) error {
	return c.Call("CSS.trackComputedStyleUpdatesForNode", args, nil)
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
Modifies the property rule property name.
*/
func SetPropertyRulePropertyName(c protocol.Caller, args SetPropertyRulePropertyNameArgs) (*SetPropertyRulePropertyNameVal, error) {
	var val = &SetPropertyRulePropertyNameVal{}
	return val, c.Call("CSS.setPropertyRulePropertyName", args, val)
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
 */
func SetContainerQueryConditionText(c protocol.Caller, args SetContainerQueryConditionTextArgs) (*SetContainerQueryConditionTextVal, error) {
	var val = &SetContainerQueryConditionTextVal{}
	return val, c.Call("CSS.setContainerQueryConditionText", args, val)
}

/*
Modifies the expression of a supports at-rule.
*/
func SetSupportsText(c protocol.Caller, args SetSupportsTextArgs) (*SetSupportsTextVal, error) {
	var val = &SetSupportsTextVal{}
	return val, c.Call("CSS.setSupportsText", args, val)
}

/*
Modifies the expression of a navigation at-rule.
*/
func SetNavigationText(c protocol.Caller, args SetNavigationTextArgs) (*SetNavigationTextVal, error) {
	var val = &SetNavigationTextVal{}
	return val, c.Call("CSS.setNavigationText", args, val)
}

/*
Modifies the expression of a scope at-rule.
*/
func SetScopeText(c protocol.Caller, args SetScopeTextArgs) (*SetScopeTextVal, error) {
	var val = &SetScopeTextVal{}
	return val, c.Call("CSS.setScopeText", args, val)
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

`takeCoverageDelta` (or since start of coverage instrumentation).
*/
func StopRuleUsageTracking(c protocol.Caller) (*StopRuleUsageTrackingVal, error) {
	var val = &StopRuleUsageTrackingVal{}
	return val, c.Call("CSS.stopRuleUsageTracking", nil, val)
}

/*
	Obtain list of rules that became used since last call to this method (or since start of coverage

instrumentation).
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
