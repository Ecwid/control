package page

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/debugger"
	"github.com/ecwid/control/protocol/io"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/runtime"
)

/*
Unique frame identifier.
*/
type FrameId string

/*
Indicates whether a frame has been identified as an ad.
*/
type AdFrameType string

/*
 */
type AdFrameExplanation string

/*
Indicates whether a frame has been identified as an ad and why.
*/
type AdFrameStatus struct {
	AdFrameType  AdFrameType          `json:"adFrameType"`
	Explanations []AdFrameExplanation `json:"explanations,omitempty"`
}

/*
	Identifies the bottom-most script which caused the frame to be labelled

as an ad.
*/
type AdScriptId struct {
	ScriptId   runtime.ScriptId         `json:"scriptId"`
	DebuggerId runtime.UniqueDebuggerId `json:"debuggerId"`
}

/*
Indicates whether the frame is a secure context and why it is the case.
*/
type SecureContextType string

/*
Indicates whether the frame is cross-origin isolated and why it is the case.
*/
type CrossOriginIsolatedContextType string

/*
 */
type GatedAPIFeatures string

/*
	All Permissions Policy features. This enum should match the one defined

in third_party/blink/renderer/core/permissions_policy/permissions_policy_features.json5.
*/
type PermissionsPolicyFeature string

/*
Reason for a permissions policy feature to be disabled.
*/
type PermissionsPolicyBlockReason string

/*
 */
type PermissionsPolicyBlockLocator struct {
	FrameId     common.FrameId               `json:"frameId"`
	BlockReason PermissionsPolicyBlockReason `json:"blockReason"`
}

/*
 */
type PermissionsPolicyFeatureState struct {
	Feature PermissionsPolicyFeature       `json:"feature"`
	Allowed bool                           `json:"allowed"`
	Locator *PermissionsPolicyBlockLocator `json:"locator,omitempty"`
}

/*
	Origin Trial(https://www.chromium.org/blink/origin-trials) support.

Status for an Origin Trial token.
*/
type OriginTrialTokenStatus string

/*
Status for an Origin Trial.
*/
type OriginTrialStatus string

/*
 */
type OriginTrialUsageRestriction string

/*
 */
type OriginTrialToken struct {
	Origin           string                      `json:"origin"`
	MatchSubDomains  bool                        `json:"matchSubDomains"`
	TrialName        string                      `json:"trialName"`
	ExpiryTime       common.TimeSinceEpoch       `json:"expiryTime"`
	IsThirdParty     bool                        `json:"isThirdParty"`
	UsageRestriction OriginTrialUsageRestriction `json:"usageRestriction"`
}

/*
 */
type OriginTrialTokenWithStatus struct {
	RawTokenText string                 `json:"rawTokenText"`
	ParsedToken  *OriginTrialToken      `json:"parsedToken,omitempty"`
	Status       OriginTrialTokenStatus `json:"status"`
}

/*
 */
type OriginTrial struct {
	TrialName        string                        `json:"trialName"`
	Status           OriginTrialStatus             `json:"status"`
	TokensWithStatus []*OriginTrialTokenWithStatus `json:"tokensWithStatus"`
}

/*
Information about the Frame on the page.
*/
type Frame struct {
	Id                             common.FrameId                 `json:"id"`
	ParentId                       common.FrameId                 `json:"parentId,omitempty"`
	LoaderId                       network.LoaderId               `json:"loaderId"`
	Name                           string                         `json:"name,omitempty"`
	Url                            string                         `json:"url"`
	UrlFragment                    string                         `json:"urlFragment,omitempty"`
	DomainAndRegistry              string                         `json:"domainAndRegistry"`
	SecurityOrigin                 string                         `json:"securityOrigin"`
	MimeType                       string                         `json:"mimeType"`
	UnreachableUrl                 string                         `json:"unreachableUrl,omitempty"`
	AdFrameStatus                  *AdFrameStatus                 `json:"adFrameStatus,omitempty"`
	SecureContextType              SecureContextType              `json:"secureContextType"`
	CrossOriginIsolatedContextType CrossOriginIsolatedContextType `json:"crossOriginIsolatedContextType"`
	GatedAPIFeatures               []GatedAPIFeatures             `json:"gatedAPIFeatures"`
}

/*
Information about the Resource on the page.
*/
type FrameResource struct {
	Url          string                `json:"url"`
	Type         network.ResourceType  `json:"type"`
	MimeType     string                `json:"mimeType"`
	LastModified common.TimeSinceEpoch `json:"lastModified,omitempty"`
	ContentSize  float64               `json:"contentSize,omitempty"`
	Failed       bool                  `json:"failed,omitempty"`
	Canceled     bool                  `json:"canceled,omitempty"`
}

/*
Information about the Frame hierarchy along with their cached resources.
*/
type FrameResourceTree struct {
	Frame       *Frame               `json:"frame"`
	ChildFrames []*FrameResourceTree `json:"childFrames,omitempty"`
	Resources   []*FrameResource     `json:"resources"`
}

/*
Information about the Frame hierarchy.
*/
type FrameTree struct {
	Frame       *Frame       `json:"frame"`
	ChildFrames []*FrameTree `json:"childFrames,omitempty"`
}

/*
Unique script identifier.
*/
type ScriptIdentifier string

/*
Transition type.
*/
type TransitionType string

/*
Navigation history entry.
*/
type NavigationEntry struct {
	Id             int            `json:"id"`
	Url            string         `json:"url"`
	UserTypedURL   string         `json:"userTypedURL"`
	Title          string         `json:"title"`
	TransitionType TransitionType `json:"transitionType"`
}

/*
Screencast frame metadata.
*/
type ScreencastFrameMetadata struct {
	OffsetTop       float64               `json:"offsetTop"`
	PageScaleFactor float64               `json:"pageScaleFactor"`
	DeviceWidth     float64               `json:"deviceWidth"`
	DeviceHeight    float64               `json:"deviceHeight"`
	ScrollOffsetX   float64               `json:"scrollOffsetX"`
	ScrollOffsetY   float64               `json:"scrollOffsetY"`
	Timestamp       common.TimeSinceEpoch `json:"timestamp,omitempty"`
}

/*
Javascript dialog type.
*/
type DialogType string

/*
Error while paring app manifest.
*/
type AppManifestError struct {
	Message  string `json:"message"`
	Critical int    `json:"critical"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`
}

/*
Parsed app manifest properties.
*/
type AppManifestParsedProperties struct {
	Scope string `json:"scope"`
}

/*
Layout viewport position and dimensions.
*/
type LayoutViewport struct {
	PageX        int `json:"pageX"`
	PageY        int `json:"pageY"`
	ClientWidth  int `json:"clientWidth"`
	ClientHeight int `json:"clientHeight"`
}

/*
Visual viewport position, dimensions, and scale.
*/
type VisualViewport struct {
	OffsetX      float64 `json:"offsetX"`
	OffsetY      float64 `json:"offsetY"`
	PageX        float64 `json:"pageX"`
	PageY        float64 `json:"pageY"`
	ClientWidth  float64 `json:"clientWidth"`
	ClientHeight float64 `json:"clientHeight"`
	Scale        float64 `json:"scale"`
	Zoom         float64 `json:"zoom,omitempty"`
}

/*
Viewport for capturing screenshot.
*/
type Viewport struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Scale  float64 `json:"scale"`
}

/*
Generic font families collection.
*/
type FontFamilies struct {
	Standard  string `json:"standard,omitempty"`
	Fixed     string `json:"fixed,omitempty"`
	Serif     string `json:"serif,omitempty"`
	SansSerif string `json:"sansSerif,omitempty"`
	Cursive   string `json:"cursive,omitempty"`
	Fantasy   string `json:"fantasy,omitempty"`
	Math      string `json:"math,omitempty"`
}

/*
Font families collection for a script.
*/
type ScriptFontFamilies struct {
	Script       string        `json:"script"`
	FontFamilies *FontFamilies `json:"fontFamilies"`
}

/*
Default font sizes.
*/
type FontSizes struct {
	Standard int `json:"standard,omitempty"`
	Fixed    int `json:"fixed,omitempty"`
}

/*
 */
type ClientNavigationReason string

/*
 */
type ClientNavigationDisposition string

/*
 */
type InstallabilityErrorArgument struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
The installability error
*/
type InstallabilityError struct {
	ErrorId        string                         `json:"errorId"`
	ErrorArguments []*InstallabilityErrorArgument `json:"errorArguments"`
}

/*
The referring-policy used for the navigation.
*/
type ReferrerPolicy string

/*
Per-script compilation cache parameters for `Page.produceCompilationCache`
*/
type CompilationCacheParams struct {
	Url   string `json:"url"`
	Eager bool   `json:"eager,omitempty"`
}

/*
The type of a frameNavigated event.
*/
type NavigationType string

/*
List of not restored reasons for back-forward cache.
*/
type BackForwardCacheNotRestoredReason string

/*
Types of not restored reasons for back-forward cache.
*/
type BackForwardCacheNotRestoredReasonType string

/*
 */
type BackForwardCacheNotRestoredExplanation struct {
	Type    BackForwardCacheNotRestoredReasonType `json:"type"`
	Reason  BackForwardCacheNotRestoredReason     `json:"reason"`
	Context string                                `json:"context,omitempty"`
}

/*
 */
type BackForwardCacheNotRestoredExplanationTree struct {
	Url          string                                        `json:"url"`
	Explanations []*BackForwardCacheNotRestoredExplanation     `json:"explanations"`
	Children     []*BackForwardCacheNotRestoredExplanationTree `json:"children"`
}

/*
List of FinalStatus reasons for Prerender2.
*/
type PrerenderFinalStatus string

type AddScriptToEvaluateOnNewDocumentArgs struct {
	Source                string `json:"source"`
	WorldName             string `json:"worldName,omitempty"`
	IncludeCommandLineAPI bool   `json:"includeCommandLineAPI,omitempty"`
}

type AddScriptToEvaluateOnNewDocumentVal struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

type CaptureScreenshotArgs struct {
	Format                string    `json:"format,omitempty"`
	Quality               int       `json:"quality,omitempty"`
	Clip                  *Viewport `json:"clip,omitempty"`
	FromSurface           bool      `json:"fromSurface,omitempty"`
	CaptureBeyondViewport bool      `json:"captureBeyondViewport,omitempty"`
	OptimizeForSpeed      bool      `json:"optimizeForSpeed,omitempty"`
}

type CaptureScreenshotVal struct {
	Data []byte `json:"data"`
}

type CaptureSnapshotArgs struct {
	Format string `json:"format,omitempty"`
}

type CaptureSnapshotVal struct {
	Data string `json:"data"`
}

type CreateIsolatedWorldArgs struct {
	FrameId             common.FrameId `json:"frameId"`
	WorldName           string         `json:"worldName,omitempty"`
	GrantUniveralAccess bool           `json:"grantUniveralAccess,omitempty"`
}

type CreateIsolatedWorldVal struct {
	ExecutionContextId runtime.ExecutionContextId `json:"executionContextId"`
}

type GetAppManifestVal struct {
	Url    string                       `json:"url"`
	Errors []*AppManifestError          `json:"errors"`
	Data   string                       `json:"data,omitempty"`
	Parsed *AppManifestParsedProperties `json:"parsed,omitempty"`
}

type GetInstallabilityErrorsVal struct {
	InstallabilityErrors []*InstallabilityError `json:"installabilityErrors"`
}

type GetManifestIconsVal struct {
	PrimaryIcon []byte `json:"primaryIcon,omitempty"`
}

type GetAppIdVal struct {
	AppId         string `json:"appId,omitempty"`
	RecommendedId string `json:"recommendedId,omitempty"`
}

type GetAdScriptIdArgs struct {
	FrameId common.FrameId `json:"frameId"`
}

type GetAdScriptIdVal struct {
	AdScriptId *AdScriptId `json:"adScriptId,omitempty"`
}

type GetFrameTreeVal struct {
	FrameTree *FrameTree `json:"frameTree"`
}

type GetLayoutMetricsVal struct {
	CssLayoutViewport *LayoutViewport `json:"cssLayoutViewport"`
	CssVisualViewport *VisualViewport `json:"cssVisualViewport"`
	CssContentSize    *common.Rect    `json:"cssContentSize"`
}

type GetNavigationHistoryVal struct {
	CurrentIndex int                `json:"currentIndex"`
	Entries      []*NavigationEntry `json:"entries"`
}

type GetResourceContentArgs struct {
	FrameId common.FrameId `json:"frameId"`
	Url     string         `json:"url"`
}

type GetResourceContentVal struct {
	Content       string `json:"content"`
	Base64Encoded bool   `json:"base64Encoded"`
}

type GetResourceTreeVal struct {
	FrameTree *FrameResourceTree `json:"frameTree"`
}

type HandleJavaScriptDialogArgs struct {
	Accept     bool   `json:"accept"`
	PromptText string `json:"promptText,omitempty"`
}

type NavigateArgs struct {
	Url            string         `json:"url"`
	Referrer       string         `json:"referrer,omitempty"`
	TransitionType TransitionType `json:"transitionType,omitempty"`
	FrameId        common.FrameId `json:"frameId,omitempty"`
	ReferrerPolicy ReferrerPolicy `json:"referrerPolicy,omitempty"`
}

type NavigateVal struct {
	FrameId   common.FrameId   `json:"frameId"`
	LoaderId  network.LoaderId `json:"loaderId,omitempty"`
	ErrorText string           `json:"errorText,omitempty"`
}

type NavigateToHistoryEntryArgs struct {
	EntryId int `json:"entryId"`
}

type PrintToPDFArgs struct {
	Landscape           bool    `json:"landscape,omitempty"`
	DisplayHeaderFooter bool    `json:"displayHeaderFooter,omitempty"`
	PrintBackground     bool    `json:"printBackground,omitempty"`
	Scale               float64 `json:"scale,omitempty"`
	PaperWidth          float64 `json:"paperWidth,omitempty"`
	PaperHeight         float64 `json:"paperHeight,omitempty"`
	MarginTop           float64 `json:"marginTop,omitempty"`
	MarginBottom        float64 `json:"marginBottom,omitempty"`
	MarginLeft          float64 `json:"marginLeft,omitempty"`
	MarginRight         float64 `json:"marginRight,omitempty"`
	PageRanges          string  `json:"pageRanges,omitempty"`
	HeaderTemplate      string  `json:"headerTemplate,omitempty"`
	FooterTemplate      string  `json:"footerTemplate,omitempty"`
	PreferCSSPageSize   bool    `json:"preferCSSPageSize,omitempty"`
	TransferMode        string  `json:"transferMode,omitempty"`
}

type PrintToPDFVal struct {
	Data   []byte          `json:"data"`
	Stream io.StreamHandle `json:"stream,omitempty"`
}

type ReloadArgs struct {
	IgnoreCache            bool   `json:"ignoreCache,omitempty"`
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"`
}

type RemoveScriptToEvaluateOnNewDocumentArgs struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

type ScreencastFrameAckArgs struct {
	SessionId int `json:"sessionId"`
}

type SearchInResourceArgs struct {
	FrameId       common.FrameId `json:"frameId"`
	Url           string         `json:"url"`
	Query         string         `json:"query"`
	CaseSensitive bool           `json:"caseSensitive,omitempty"`
	IsRegex       bool           `json:"isRegex,omitempty"`
}

type SearchInResourceVal struct {
	Result []*debugger.SearchMatch `json:"result"`
}

type SetAdBlockingEnabledArgs struct {
	Enabled bool `json:"enabled"`
}

type SetBypassCSPArgs struct {
	Enabled bool `json:"enabled"`
}

type GetPermissionsPolicyStateArgs struct {
	FrameId common.FrameId `json:"frameId"`
}

type GetPermissionsPolicyStateVal struct {
	States []*PermissionsPolicyFeatureState `json:"states"`
}

type GetOriginTrialsArgs struct {
	FrameId common.FrameId `json:"frameId"`
}

type GetOriginTrialsVal struct {
	OriginTrials []*OriginTrial `json:"originTrials"`
}

type SetFontFamiliesArgs struct {
	FontFamilies *FontFamilies         `json:"fontFamilies"`
	ForScripts   []*ScriptFontFamilies `json:"forScripts,omitempty"`
}

type SetFontSizesArgs struct {
	FontSizes *FontSizes `json:"fontSizes"`
}

type SetDocumentContentArgs struct {
	FrameId common.FrameId `json:"frameId"`
	Html    string         `json:"html"`
}

type SetLifecycleEventsEnabledArgs struct {
	Enabled bool `json:"enabled"`
}

type StartScreencastArgs struct {
	Format        string `json:"format,omitempty"`
	Quality       int    `json:"quality,omitempty"`
	MaxWidth      int    `json:"maxWidth,omitempty"`
	MaxHeight     int    `json:"maxHeight,omitempty"`
	EveryNthFrame int    `json:"everyNthFrame,omitempty"`
}

type SetWebLifecycleStateArgs struct {
	State string `json:"state"`
}

type ProduceCompilationCacheArgs struct {
	Scripts []*CompilationCacheParams `json:"scripts"`
}

type AddCompilationCacheArgs struct {
	Url  string `json:"url"`
	Data []byte `json:"data"`
}

type SetSPCTransactionModeArgs struct {
	Mode string `json:"mode"`
}

type GenerateTestReportArgs struct {
	Message string `json:"message"`
	Group   string `json:"group,omitempty"`
}

type SetInterceptFileChooserDialogArgs struct {
	Enabled bool `json:"enabled"`
}
