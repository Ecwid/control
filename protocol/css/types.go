package css

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
)

/*

 */
type StyleSheetId string

/*
	Stylesheet type: "injected" for stylesheets injected via extension, "user-agent" for user-agent
stylesheets, "inspector" for stylesheets created by the inspector (i.e. those holding the "via
inspector" rules), "regular" for regular stylesheets.
*/
type StyleSheetOrigin string

/*
	CSS rule collection for a single pseudo style.
*/
type PseudoElementMatches struct {
	PseudoType dom.PseudoType `json:"pseudoType"`
	Matches    []*RuleMatch   `json:"matches"`
}

/*
	Inherited CSS rule collection from ancestor node.
*/
type InheritedStyleEntry struct {
	InlineStyle     *CSSStyle    `json:"inlineStyle,omitempty"`
	MatchedCSSRules []*RuleMatch `json:"matchedCSSRules"`
}

/*
	Match data for a CSS rule.
*/
type RuleMatch struct {
	Rule              *CSSRule `json:"rule"`
	MatchingSelectors []int    `json:"matchingSelectors"`
}

/*
	Data for a simple selector (these are delimited by commas in a selector list).
*/
type Value struct {
	Text  string       `json:"text"`
	Range *SourceRange `json:"range,omitempty"`
}

/*
	Selector list data.
*/
type SelectorList struct {
	Selectors []*Value `json:"selectors"`
	Text      string   `json:"text"`
}

/*
	CSS stylesheet metainformation.
*/
type CSSStyleSheetHeader struct {
	StyleSheetId  StyleSheetId      `json:"styleSheetId"`
	FrameId       common.FrameId    `json:"frameId"`
	SourceURL     string            `json:"sourceURL"`
	SourceMapURL  string            `json:"sourceMapURL,omitempty"`
	Origin        StyleSheetOrigin  `json:"origin"`
	Title         string            `json:"title"`
	OwnerNode     dom.BackendNodeId `json:"ownerNode,omitempty"`
	Disabled      bool              `json:"disabled"`
	HasSourceURL  bool              `json:"hasSourceURL,omitempty"`
	IsInline      bool              `json:"isInline"`
	IsMutable     bool              `json:"isMutable"`
	IsConstructed bool              `json:"isConstructed"`
	StartLine     float64           `json:"startLine"`
	StartColumn   float64           `json:"startColumn"`
	Length        float64           `json:"length"`
	EndLine       float64           `json:"endLine"`
	EndColumn     float64           `json:"endColumn"`
}

/*
	CSS rule representation.
*/
type CSSRule struct {
	StyleSheetId StyleSheetId     `json:"styleSheetId,omitempty"`
	SelectorList *SelectorList    `json:"selectorList"`
	Origin       StyleSheetOrigin `json:"origin"`
	Style        *CSSStyle        `json:"style"`
	Media        []*CSSMedia      `json:"media,omitempty"`
}

/*
	CSS coverage information.
*/
type RuleUsage struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
	StartOffset  float64      `json:"startOffset"`
	EndOffset    float64      `json:"endOffset"`
	Used         bool         `json:"used"`
}

/*
	Text range within a resource. All numbers are zero-based.
*/
type SourceRange struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn"`
	EndLine     int `json:"endLine"`
	EndColumn   int `json:"endColumn"`
}

/*

 */
type ShorthandEntry struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Important bool   `json:"important,omitempty"`
}

/*

 */
type CSSComputedStyleProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
	CSS style representation.
*/
type CSSStyle struct {
	StyleSheetId     StyleSheetId      `json:"styleSheetId,omitempty"`
	CssProperties    []*CSSProperty    `json:"cssProperties"`
	ShorthandEntries []*ShorthandEntry `json:"shorthandEntries"`
	CssText          string            `json:"cssText,omitempty"`
	Range            *SourceRange      `json:"range,omitempty"`
}

/*
	CSS property declaration data.
*/
type CSSProperty struct {
	Name      string       `json:"name"`
	Value     string       `json:"value"`
	Important bool         `json:"important,omitempty"`
	Implicit  bool         `json:"implicit,omitempty"`
	Text      string       `json:"text,omitempty"`
	ParsedOk  bool         `json:"parsedOk,omitempty"`
	Disabled  bool         `json:"disabled,omitempty"`
	Range     *SourceRange `json:"range,omitempty"`
}

/*
	CSS media rule descriptor.
*/
type CSSMedia struct {
	Text         string        `json:"text"`
	Source       string        `json:"source"`
	SourceURL    string        `json:"sourceURL,omitempty"`
	Range        *SourceRange  `json:"range,omitempty"`
	StyleSheetId StyleSheetId  `json:"styleSheetId,omitempty"`
	MediaList    []*MediaQuery `json:"mediaList,omitempty"`
}

/*
	Media query descriptor.
*/
type MediaQuery struct {
	Expressions []*MediaQueryExpression `json:"expressions"`
	Active      bool                    `json:"active"`
}

/*
	Media query expression descriptor.
*/
type MediaQueryExpression struct {
	Value          float64      `json:"value"`
	Unit           string       `json:"unit"`
	Feature        string       `json:"feature"`
	ValueRange     *SourceRange `json:"valueRange,omitempty"`
	ComputedLength float64      `json:"computedLength,omitempty"`
}

/*
	Information about amount of glyphs that were rendered with given font.
*/
type PlatformFontUsage struct {
	FamilyName   string  `json:"familyName"`
	IsCustomFont bool    `json:"isCustomFont"`
	GlyphCount   float64 `json:"glyphCount"`
}

/*
	Information about font variation axes for variable fonts
*/
type FontVariationAxis struct {
	Tag          string  `json:"tag"`
	Name         string  `json:"name"`
	MinValue     float64 `json:"minValue"`
	MaxValue     float64 `json:"maxValue"`
	DefaultValue float64 `json:"defaultValue"`
}

/*
	Properties of a web font: https://www.w3.org/TR/2008/REC-CSS2-20080411/fonts.html#font-descriptions
and additional information such as platformFontFamily and fontVariationAxes.
*/
type FontFace struct {
	FontFamily         string               `json:"fontFamily"`
	FontStyle          string               `json:"fontStyle"`
	FontVariant        string               `json:"fontVariant"`
	FontWeight         string               `json:"fontWeight"`
	FontStretch        string               `json:"fontStretch"`
	UnicodeRange       string               `json:"unicodeRange"`
	Src                string               `json:"src"`
	PlatformFontFamily string               `json:"platformFontFamily"`
	FontVariationAxes  []*FontVariationAxis `json:"fontVariationAxes,omitempty"`
}

/*
	CSS keyframes rule representation.
*/
type CSSKeyframesRule struct {
	AnimationName *Value             `json:"animationName"`
	Keyframes     []*CSSKeyframeRule `json:"keyframes"`
}

/*
	CSS keyframe rule representation.
*/
type CSSKeyframeRule struct {
	StyleSheetId StyleSheetId     `json:"styleSheetId,omitempty"`
	Origin       StyleSheetOrigin `json:"origin"`
	KeyText      *Value           `json:"keyText"`
	Style        *CSSStyle        `json:"style"`
}

/*
	A descriptor of operation to mutate style declaration text.
*/
type StyleDeclarationEdit struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Text         string       `json:"text"`
}

type AddRuleArgs struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
	RuleText     string       `json:"ruleText"`
	Location     *SourceRange `json:"location"`
}

type AddRuleVal struct {
	Rule *CSSRule `json:"rule"`
}

type CollectClassNamesArgs struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
}

type CollectClassNamesVal struct {
	ClassNames []string `json:"classNames"`
}

type CreateStyleSheetArgs struct {
	FrameId common.FrameId `json:"frameId"`
}

type CreateStyleSheetVal struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
}

type ForcePseudoStateArgs struct {
	NodeId              dom.NodeId `json:"nodeId"`
	ForcedPseudoClasses []string   `json:"forcedPseudoClasses"`
}

type GetBackgroundColorsArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetBackgroundColorsVal struct {
	BackgroundColors   []string `json:"backgroundColors,omitempty"`
	ComputedFontSize   string   `json:"computedFontSize,omitempty"`
	ComputedFontWeight string   `json:"computedFontWeight,omitempty"`
}

type GetComputedStyleForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetComputedStyleForNodeVal struct {
	ComputedStyle []*CSSComputedStyleProperty `json:"computedStyle"`
}

type GetInlineStylesForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetInlineStylesForNodeVal struct {
	InlineStyle     *CSSStyle `json:"inlineStyle,omitempty"`
	AttributesStyle *CSSStyle `json:"attributesStyle,omitempty"`
}

type GetMatchedStylesForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetMatchedStylesForNodeVal struct {
	InlineStyle       *CSSStyle               `json:"inlineStyle,omitempty"`
	AttributesStyle   *CSSStyle               `json:"attributesStyle,omitempty"`
	MatchedCSSRules   []*RuleMatch            `json:"matchedCSSRules,omitempty"`
	PseudoElements    []*PseudoElementMatches `json:"pseudoElements,omitempty"`
	Inherited         []*InheritedStyleEntry  `json:"inherited,omitempty"`
	CssKeyframesRules []*CSSKeyframesRule     `json:"cssKeyframesRules,omitempty"`
}

type GetMediaQueriesVal struct {
	Medias []*CSSMedia `json:"medias"`
}

type GetPlatformFontsForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetPlatformFontsForNodeVal struct {
	Fonts []*PlatformFontUsage `json:"fonts"`
}

type GetStyleSheetTextArgs struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
}

type GetStyleSheetTextVal struct {
	Text string `json:"text"`
}

type TrackComputedStyleUpdatesArgs struct {
	PropertiesToTrack []*CSSComputedStyleProperty `json:"propertiesToTrack"`
}

type TakeComputedStyleUpdatesVal struct {
	NodeIds []dom.NodeId `json:"nodeIds"`
}

type SetEffectivePropertyValueForNodeArgs struct {
	NodeId       dom.NodeId `json:"nodeId"`
	PropertyName string     `json:"propertyName"`
	Value        string     `json:"value"`
}

type SetKeyframeKeyArgs struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	KeyText      string       `json:"keyText"`
}

type SetKeyframeKeyVal struct {
	KeyText *Value `json:"keyText"`
}

type SetMediaTextArgs struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Text         string       `json:"text"`
}

type SetMediaTextVal struct {
	Media *CSSMedia `json:"media"`
}

type SetRuleSelectorArgs struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange `json:"range"`
	Selector     string       `json:"selector"`
}

type SetRuleSelectorVal struct {
	SelectorList *SelectorList `json:"selectorList"`
}

type SetStyleSheetTextArgs struct {
	StyleSheetId StyleSheetId `json:"styleSheetId"`
	Text         string       `json:"text"`
}

type SetStyleSheetTextVal struct {
	SourceMapURL string `json:"sourceMapURL,omitempty"`
}

type SetStyleTextsArgs struct {
	Edits []*StyleDeclarationEdit `json:"edits"`
}

type SetStyleTextsVal struct {
	Styles []*CSSStyle `json:"styles"`
}

type StopRuleUsageTrackingVal struct {
	RuleUsage []*RuleUsage `json:"ruleUsage"`
}

type TakeCoverageDeltaVal struct {
	Coverage  []*RuleUsage `json:"coverage"`
	Timestamp float64      `json:"timestamp"`
}

type SetLocalFontsEnabledArgs struct {
	Enabled bool `json:"enabled"`
}
