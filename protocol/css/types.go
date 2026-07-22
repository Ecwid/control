package css

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
)

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
	PseudoType       dom.PseudoType `json:"pseudoType"`
	PseudoIdentifier string         `json:"pseudoIdentifier,omitempty"`
	Matches          []*RuleMatch   `json:"matches"`
}

/*
CSS style coming from animations with the name of the animation.
*/
type CSSAnimationStyle struct {
	Name  string    `json:"name,omitempty"`
	Style *CSSStyle `json:"style"`
}

/*
Inherited CSS rule collection from ancestor node.
*/
type InheritedStyleEntry struct {
	InlineStyle     *CSSStyle    `json:"inlineStyle,omitempty"`
	MatchedCSSRules []*RuleMatch `json:"matchedCSSRules"`
}

/*
Inherited CSS style collection for animated styles from ancestor node.
*/
type InheritedAnimatedStyleEntry struct {
	AnimationStyles  []*CSSAnimationStyle `json:"animationStyles,omitempty"`
	TransitionsStyle *CSSStyle            `json:"transitionsStyle,omitempty"`
}

/*
Inherited pseudo element matches from pseudos of an ancestor node.
*/
type InheritedPseudoElementMatches struct {
	PseudoElements []*PseudoElementMatches `json:"pseudoElements"`
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
	Text        string       `json:"text"`
	Range       *SourceRange `json:"range,omitempty"`
	Specificity *Specificity `json:"specificity,omitempty"`
}

/*
	Specificity:

https://drafts.csswg.org/selectors/#specificity-rules
*/
type Specificity struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
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
	StyleSheetId  dom.StyleSheetId  `json:"styleSheetId"`
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
	LoadingFailed bool              `json:"loadingFailed,omitempty"`
}

/*
CSS rule representation.
*/
type CSSRule struct {
	StyleSheetId          dom.StyleSheetId     `json:"styleSheetId,omitempty"`
	SelectorList          *SelectorList        `json:"selectorList"`
	NestingSelectors      []string             `json:"nestingSelectors,omitempty"`
	Origin                StyleSheetOrigin     `json:"origin"`
	Style                 *CSSStyle            `json:"style"`
	OriginTreeScopeNodeId dom.BackendNodeId    `json:"originTreeScopeNodeId,omitempty"`
	Media                 []*CSSMedia          `json:"media,omitempty"`
	ContainerQueries      []*CSSContainerQuery `json:"containerQueries,omitempty"`
	Supports              []*CSSSupports       `json:"supports,omitempty"`
	Layers                []*CSSLayer          `json:"layers,omitempty"`
	Scopes                []*CSSScope          `json:"scopes,omitempty"`
	RuleTypes             []CSSRuleType        `json:"ruleTypes,omitempty"`
	StartingStyles        []*CSSStartingStyle  `json:"startingStyles,omitempty"`
	Navigations           []*CSSNavigation     `json:"navigations,omitempty"`
}

/*
	Enum indicating the type of a CSS rule, used to represent the order of a style rule's ancestors.

This list only contains rule types that are collected during the ancestor rule collection.
*/
type CSSRuleType string

/*
CSS coverage information.
*/
type RuleUsage struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	StartOffset  float64          `json:"startOffset"`
	EndOffset    float64          `json:"endOffset"`
	Used         bool             `json:"used"`
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
 */
type ComputedStyleExtraFields struct {
	IsAppearanceBase bool `json:"isAppearanceBase"`
}

/*
CSS style representation.
*/
type CSSStyle struct {
	StyleSheetId     dom.StyleSheetId  `json:"styleSheetId,omitempty"`
	CssProperties    []*CSSProperty    `json:"cssProperties"`
	ShorthandEntries []*ShorthandEntry `json:"shorthandEntries"`
	CssText          string            `json:"cssText,omitempty"`
	Range            *SourceRange      `json:"range,omitempty"`
}

/*
CSS property declaration data.
*/
type CSSProperty struct {
	Name               string         `json:"name"`
	Value              string         `json:"value"`
	Important          bool           `json:"important,omitempty"`
	Implicit           bool           `json:"implicit,omitempty"`
	Text               string         `json:"text,omitempty"`
	ParsedOk           bool           `json:"parsedOk,omitempty"`
	Disabled           bool           `json:"disabled,omitempty"`
	Range              *SourceRange   `json:"range,omitempty"`
	LonghandProperties []*CSSProperty `json:"longhandProperties,omitempty"`
}

/*
CSS media rule descriptor.
*/
type CSSMedia struct {
	Text         string           `json:"text"`
	Source       string           `json:"source"`
	SourceURL    string           `json:"sourceURL,omitempty"`
	Range        *SourceRange     `json:"range,omitempty"`
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
	MediaList    []*MediaQuery    `json:"mediaList,omitempty"`
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
CSS container query rule descriptor.
*/
type CSSContainerQuery struct {
	Range              *SourceRange     `json:"range,omitempty"`
	StyleSheetId       dom.StyleSheetId `json:"styleSheetId,omitempty"`
	Name               string           `json:"name,omitempty"`
	PhysicalAxes       dom.PhysicalAxes `json:"physicalAxes,omitempty"`
	LogicalAxes        dom.LogicalAxes  `json:"logicalAxes,omitempty"`
	QueriesScrollState bool             `json:"queriesScrollState,omitempty"`
	QueriesAnchored    bool             `json:"queriesAnchored,omitempty"`
	ConditionText      string           `json:"conditionText"`
}

/*
CSS Supports at-rule descriptor.
*/
type CSSSupports struct {
	Text         string           `json:"text"`
	Active       bool             `json:"active"`
	Range        *SourceRange     `json:"range,omitempty"`
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
}

/*
CSS Navigation at-rule descriptor.
*/
type CSSNavigation struct {
	Text         string           `json:"text"`
	Active       bool             `json:"active,omitempty"`
	Range        *SourceRange     `json:"range,omitempty"`
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
}

/*
CSS Scope at-rule descriptor.
*/
type CSSScope struct {
	Text         string           `json:"text"`
	Range        *SourceRange     `json:"range,omitempty"`
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
}

/*
CSS Layer at-rule descriptor.
*/
type CSSLayer struct {
	Text         string           `json:"text"`
	Range        *SourceRange     `json:"range,omitempty"`
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
}

/*
CSS Starting Style at-rule descriptor.
*/
type CSSStartingStyle struct {
	Range        *SourceRange     `json:"range,omitempty"`
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
}

/*
CSS Layer data.
*/
type CSSLayerData struct {
	Name      string          `json:"name"`
	SubLayers []*CSSLayerData `json:"subLayers,omitempty"`
	Order     float64         `json:"order"`
}

/*
Information about amount of glyphs that were rendered with given font.
*/
type PlatformFontUsage struct {
	FamilyName     string  `json:"familyName"`
	PostScriptName string  `json:"postScriptName"`
	IsCustomFont   bool    `json:"isCustomFont"`
	GlyphCount     float64 `json:"glyphCount"`
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
	FontDisplay        string               `json:"fontDisplay"`
	UnicodeRange       string               `json:"unicodeRange"`
	Src                string               `json:"src"`
	PlatformFontFamily string               `json:"platformFontFamily"`
	FontVariationAxes  []*FontVariationAxis `json:"fontVariationAxes,omitempty"`
}

/*
CSS try rule representation.
*/
type CSSTryRule struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
	Origin       StyleSheetOrigin `json:"origin"`
	Style        *CSSStyle        `json:"style"`
}

/*
CSS @position-try rule representation.
*/
type CSSPositionTryRule struct {
	Name         *Value           `json:"name"`
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
	Origin       StyleSheetOrigin `json:"origin"`
	Style        *CSSStyle        `json:"style"`
	Active       bool             `json:"active"`
}

/*
CSS keyframes rule representation.
*/
type CSSKeyframesRule struct {
	AnimationName *Value             `json:"animationName"`
	Keyframes     []*CSSKeyframeRule `json:"keyframes"`
}

/*
Representation of a custom property registration through CSS.registerProperty
*/
type CSSPropertyRegistration struct {
	PropertyName string `json:"propertyName"`
	InitialValue *Value `json:"initialValue,omitempty"`
	Inherits     bool   `json:"inherits"`
	Syntax       string `json:"syntax"`
}

/*
CSS generic @rule representation.
*/
type CSSAtRule struct {
	Type         string           `json:"type"`
	Subsection   string           `json:"subsection,omitempty"`
	Name         *Value           `json:"name,omitempty"`
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
	Origin       StyleSheetOrigin `json:"origin"`
	Style        *CSSStyle        `json:"style"`
}

/*
CSS property at-rule representation.
*/
type CSSPropertyRule struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
	Origin       StyleSheetOrigin `json:"origin"`
	PropertyName *Value           `json:"propertyName"`
	Style        *CSSStyle        `json:"style"`
}

/*
CSS function argument representation.
*/
type CSSFunctionParameter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

/*
CSS function conditional block representation.
*/
type CSSFunctionConditionNode struct {
	Media            *CSSMedia          `json:"media,omitempty"`
	ContainerQueries *CSSContainerQuery `json:"containerQueries,omitempty"`
	Supports         *CSSSupports       `json:"supports,omitempty"`
	Navigation       *CSSNavigation     `json:"navigation,omitempty"`
	Children         []*CSSFunctionNode `json:"children"`
	ConditionText    string             `json:"conditionText"`
}

/*
Section of the body of a CSS function rule.
*/
type CSSFunctionNode struct {
	Condition *CSSFunctionConditionNode `json:"condition,omitempty"`
	Style     *CSSStyle                 `json:"style,omitempty"`
}

/*
CSS function at-rule representation.
*/
type CSSFunctionRule struct {
	Name                  *Value                  `json:"name"`
	StyleSheetId          dom.StyleSheetId        `json:"styleSheetId,omitempty"`
	Origin                StyleSheetOrigin        `json:"origin"`
	Parameters            []*CSSFunctionParameter `json:"parameters"`
	Children              []*CSSFunctionNode      `json:"children"`
	OriginTreeScopeNodeId dom.BackendNodeId       `json:"originTreeScopeNodeId,omitempty"`
}

/*
CSS keyframe rule representation.
*/
type CSSKeyframeRule struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId,omitempty"`
	Origin       StyleSheetOrigin `json:"origin"`
	KeyText      *Value           `json:"keyText"`
	Style        *CSSStyle        `json:"style"`
}

/*
A descriptor of operation to mutate style declaration text.
*/
type StyleDeclarationEdit struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	Text         string           `json:"text"`
}

type AddRuleArgs struct {
	StyleSheetId                    dom.StyleSheetId `json:"styleSheetId"`
	RuleText                        string           `json:"ruleText"`
	Location                        *SourceRange     `json:"location"`
	NodeForPropertySyntaxValidation dom.NodeId       `json:"nodeForPropertySyntaxValidation,omitempty"`
}

type AddRuleVal struct {
	Rule *CSSRule `json:"rule"`
}

type CollectClassNamesArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
}

type CollectClassNamesVal struct {
	ClassNames []string `json:"classNames"`
}

type CreateStyleSheetArgs struct {
	FrameId common.FrameId `json:"frameId"`
	Force   bool           `json:"force,omitempty"`
}

type CreateStyleSheetVal struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
}

type ForcePseudoStateArgs struct {
	NodeId              dom.NodeId `json:"nodeId"`
	ForcedPseudoClasses []string   `json:"forcedPseudoClasses"`
}

type ForceStartingStyleArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
	Forced bool       `json:"forced"`
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
	ExtraFields   *ComputedStyleExtraFields   `json:"extraFields"`
}

type ResolveValuesArgs struct {
	Values           []string       `json:"values"`
	NodeId           dom.NodeId     `json:"nodeId"`
	PropertyName     string         `json:"propertyName,omitempty"`
	PseudoType       dom.PseudoType `json:"pseudoType,omitempty"`
	PseudoIdentifier string         `json:"pseudoIdentifier,omitempty"`
}

type ResolveValuesVal struct {
	Results []string `json:"results"`
}

type GetLonghandPropertiesArgs struct {
	ShorthandName string `json:"shorthandName"`
	Value         string `json:"value"`
}

type GetLonghandPropertiesVal struct {
	LonghandProperties []*CSSProperty `json:"longhandProperties"`
}

type GetInlineStylesForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetInlineStylesForNodeVal struct {
	InlineStyle     *CSSStyle `json:"inlineStyle,omitempty"`
	AttributesStyle *CSSStyle `json:"attributesStyle,omitempty"`
}

type GetAnimatedStylesForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetAnimatedStylesForNodeVal struct {
	AnimationStyles  []*CSSAnimationStyle           `json:"animationStyles,omitempty"`
	TransitionsStyle *CSSStyle                      `json:"transitionsStyle,omitempty"`
	Inherited        []*InheritedAnimatedStyleEntry `json:"inherited,omitempty"`
}

type GetMatchedStylesForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetMatchedStylesForNodeVal struct {
	InlineStyle                 *CSSStyle                        `json:"inlineStyle,omitempty"`
	AttributesStyle             *CSSStyle                        `json:"attributesStyle,omitempty"`
	MatchedCSSRules             []*RuleMatch                     `json:"matchedCSSRules,omitempty"`
	PseudoElements              []*PseudoElementMatches          `json:"pseudoElements,omitempty"`
	Inherited                   []*InheritedStyleEntry           `json:"inherited,omitempty"`
	InheritedPseudoElements     []*InheritedPseudoElementMatches `json:"inheritedPseudoElements,omitempty"`
	CssKeyframesRules           []*CSSKeyframesRule              `json:"cssKeyframesRules,omitempty"`
	CssPositionTryRules         []*CSSPositionTryRule            `json:"cssPositionTryRules,omitempty"`
	ActivePositionFallbackIndex int                              `json:"activePositionFallbackIndex,omitempty"`
	CssPropertyRules            []*CSSPropertyRule               `json:"cssPropertyRules,omitempty"`
	CssPropertyRegistrations    []*CSSPropertyRegistration       `json:"cssPropertyRegistrations,omitempty"`
	CssAtRules                  []*CSSAtRule                     `json:"cssAtRules,omitempty"`
	ParentLayoutNodeId          dom.NodeId                       `json:"parentLayoutNodeId,omitempty"`
	CssFunctionRules            []*CSSFunctionRule               `json:"cssFunctionRules,omitempty"`
}

type GetEnvironmentVariablesVal struct {
	EnvironmentVariables any `json:"environmentVariables"`
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
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
}

type GetStyleSheetTextVal struct {
	Text string `json:"text"`
}

type GetLayersForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId"`
}

type GetLayersForNodeVal struct {
	RootLayer *CSSLayerData `json:"rootLayer"`
}

type GetLocationForSelectorArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	SelectorText string           `json:"selectorText"`
}

type GetLocationForSelectorVal struct {
	Ranges []*SourceRange `json:"ranges"`
}

type TrackComputedStyleUpdatesForNodeArgs struct {
	NodeId dom.NodeId `json:"nodeId,omitempty"`
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

type SetPropertyRulePropertyNameArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	PropertyName string           `json:"propertyName"`
}

type SetPropertyRulePropertyNameVal struct {
	PropertyName *Value `json:"propertyName"`
}

type SetKeyframeKeyArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	KeyText      string           `json:"keyText"`
}

type SetKeyframeKeyVal struct {
	KeyText *Value `json:"keyText"`
}

type SetMediaTextArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	Text         string           `json:"text"`
}

type SetMediaTextVal struct {
	Media *CSSMedia `json:"media"`
}

type SetContainerQueryConditionTextArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	Text         string           `json:"text"`
}

type SetContainerQueryConditionTextVal struct {
	ContainerQuery *CSSContainerQuery `json:"containerQuery"`
}

type SetSupportsTextArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	Text         string           `json:"text"`
}

type SetSupportsTextVal struct {
	Supports *CSSSupports `json:"supports"`
}

type SetNavigationTextArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	Text         string           `json:"text"`
}

type SetNavigationTextVal struct {
	Navigation *CSSNavigation `json:"navigation"`
}

type SetScopeTextArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	Text         string           `json:"text"`
}

type SetScopeTextVal struct {
	Scope *CSSScope `json:"scope"`
}

type SetRuleSelectorArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Range        *SourceRange     `json:"range"`
	Selector     string           `json:"selector"`
}

type SetRuleSelectorVal struct {
	SelectorList *SelectorList `json:"selectorList"`
}

type SetStyleSheetTextArgs struct {
	StyleSheetId dom.StyleSheetId `json:"styleSheetId"`
	Text         string           `json:"text"`
}

type SetStyleSheetTextVal struct {
	SourceMapURL string `json:"sourceMapURL,omitempty"`
}

type SetStyleTextsArgs struct {
	Edits                           []*StyleDeclarationEdit `json:"edits"`
	NodeForPropertySyntaxValidation dom.NodeId              `json:"nodeForPropertySyntaxValidation,omitempty"`
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
