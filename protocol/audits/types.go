package audits

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/network"
	"github.com/ecwid/control/protocol/runtime"
)

/*
	Information about a cookie that is affected by an inspector issue.
*/
type AffectedCookie struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Domain string `json:"domain"`
}

/*
	Information about a request that is affected by an inspector issue.
*/
type AffectedRequest struct {
	RequestId network.RequestId `json:"requestId"`
	Url       string            `json:"url,omitempty"`
}

/*
	Information about the frame affected by an inspector issue.
*/
type AffectedFrame struct {
	FrameId common.FrameId `json:"frameId"`
}

/*

 */
type SameSiteCookieExclusionReason string

/*

 */
type SameSiteCookieWarningReason string

/*

 */
type SameSiteCookieOperation string

/*
	This information is currently necessary, as the front-end has a difficult
time finding a specific cookie. With this, we can convey specific error
information without the cookie.
*/
type SameSiteCookieIssueDetails struct {
	Cookie                 *AffectedCookie                 `json:"cookie"`
	CookieWarningReasons   []SameSiteCookieWarningReason   `json:"cookieWarningReasons"`
	CookieExclusionReasons []SameSiteCookieExclusionReason `json:"cookieExclusionReasons"`
	Operation              SameSiteCookieOperation         `json:"operation"`
	SiteForCookies         string                          `json:"siteForCookies,omitempty"`
	CookieUrl              string                          `json:"cookieUrl,omitempty"`
	Request                *AffectedRequest                `json:"request,omitempty"`
}

/*

 */
type MixedContentResolutionStatus string

/*

 */
type MixedContentResourceType string

/*

 */
type MixedContentIssueDetails struct {
	ResourceType     MixedContentResourceType     `json:"resourceType,omitempty"`
	ResolutionStatus MixedContentResolutionStatus `json:"resolutionStatus"`
	InsecureURL      string                       `json:"insecureURL"`
	MainResourceURL  string                       `json:"mainResourceURL"`
	Request          *AffectedRequest             `json:"request,omitempty"`
	Frame            *AffectedFrame               `json:"frame,omitempty"`
}

/*
	Enum indicating the reason a response has been blocked. These reasons are
refinements of the net error BLOCKED_BY_RESPONSE.
*/
type BlockedByResponseReason string

/*
	Details for a request that has been blocked with the BLOCKED_BY_RESPONSE
code. Currently only used for COEP/COOP, but may be extended to include
some CSP errors in the future.
*/
type BlockedByResponseIssueDetails struct {
	Request      *AffectedRequest        `json:"request"`
	ParentFrame  *AffectedFrame          `json:"parentFrame,omitempty"`
	BlockedFrame *AffectedFrame          `json:"blockedFrame,omitempty"`
	Reason       BlockedByResponseReason `json:"reason"`
}

/*

 */
type HeavyAdResolutionStatus string

/*

 */
type HeavyAdReason string

/*

 */
type HeavyAdIssueDetails struct {
	Resolution HeavyAdResolutionStatus `json:"resolution"`
	Reason     HeavyAdReason           `json:"reason"`
	Frame      *AffectedFrame          `json:"frame"`
}

/*

 */
type ContentSecurityPolicyViolationType string

/*

 */
type SourceCodeLocation struct {
	ScriptId     runtime.ScriptId `json:"scriptId,omitempty"`
	Url          string           `json:"url"`
	LineNumber   int              `json:"lineNumber"`
	ColumnNumber int              `json:"columnNumber"`
}

/*

 */
type ContentSecurityPolicyIssueDetails struct {
	BlockedURL                         string                             `json:"blockedURL,omitempty"`
	ViolatedDirective                  string                             `json:"violatedDirective"`
	IsReportOnly                       bool                               `json:"isReportOnly"`
	ContentSecurityPolicyViolationType ContentSecurityPolicyViolationType `json:"contentSecurityPolicyViolationType"`
	FrameAncestor                      *AffectedFrame                     `json:"frameAncestor,omitempty"`
	SourceCodeLocation                 *SourceCodeLocation                `json:"sourceCodeLocation,omitempty"`
	ViolatingNodeId                    dom.BackendNodeId                  `json:"violatingNodeId,omitempty"`
}

/*

 */
type SharedArrayBufferIssueType string

/*
	Details for a issue arising from an SAB being instantiated in, or
transfered to a context that is not cross-origin isolated.
*/
type SharedArrayBufferIssueDetails struct {
	SourceCodeLocation *SourceCodeLocation        `json:"sourceCodeLocation"`
	IsWarning          bool                       `json:"isWarning"`
	Type               SharedArrayBufferIssueType `json:"type"`
}

/*

 */
type TwaQualityEnforcementViolationType string

/*

 */
type TrustedWebActivityIssueDetails struct {
	Url            string                             `json:"url"`
	ViolationType  TwaQualityEnforcementViolationType `json:"violationType"`
	HttpStatusCode int                                `json:"httpStatusCode,omitempty"`
	PackageName    string                             `json:"packageName,omitempty"`
	Signature      string                             `json:"signature,omitempty"`
}

/*

 */
type LowTextContrastIssueDetails struct {
	ViolatingNodeId       dom.BackendNodeId `json:"violatingNodeId"`
	ViolatingNodeSelector string            `json:"violatingNodeSelector"`
	ContrastRatio         float64           `json:"contrastRatio"`
	ThresholdAA           float64           `json:"thresholdAA"`
	ThresholdAAA          float64           `json:"thresholdAAA"`
	FontSize              string            `json:"fontSize"`
	FontWeight            string            `json:"fontWeight"`
}

/*
	Details for a CORS related issue, e.g. a warning or error related to
CORS RFC1918 enforcement.
*/
type CorsIssueDetails struct {
	CorsErrorStatus        *network.CorsErrorStatus     `json:"corsErrorStatus"`
	IsWarning              bool                         `json:"isWarning"`
	Request                *AffectedRequest             `json:"request"`
	InitiatorOrigin        string                       `json:"initiatorOrigin,omitempty"`
	ResourceIPAddressSpace network.IPAddressSpace       `json:"resourceIPAddressSpace,omitempty"`
	ClientSecurityState    *network.ClientSecurityState `json:"clientSecurityState,omitempty"`
}

/*
	A unique identifier for the type of issue. Each type may use one of the
optional fields in InspectorIssueDetails to convey more specific
information about the kind of issue.
*/
type InspectorIssueCode string

/*
	This struct holds a list of optional fields with additional information
specific to the kind of issue. When adding a new issue code, please also
add a new optional field to this type.
*/
type InspectorIssueDetails struct {
	SameSiteCookieIssueDetails        *SameSiteCookieIssueDetails        `json:"sameSiteCookieIssueDetails,omitempty"`
	MixedContentIssueDetails          *MixedContentIssueDetails          `json:"mixedContentIssueDetails,omitempty"`
	BlockedByResponseIssueDetails     *BlockedByResponseIssueDetails     `json:"blockedByResponseIssueDetails,omitempty"`
	HeavyAdIssueDetails               *HeavyAdIssueDetails               `json:"heavyAdIssueDetails,omitempty"`
	ContentSecurityPolicyIssueDetails *ContentSecurityPolicyIssueDetails `json:"contentSecurityPolicyIssueDetails,omitempty"`
	SharedArrayBufferIssueDetails     *SharedArrayBufferIssueDetails     `json:"sharedArrayBufferIssueDetails,omitempty"`
	TwaQualityEnforcementDetails      *TrustedWebActivityIssueDetails    `json:"twaQualityEnforcementDetails,omitempty"`
	LowTextContrastIssueDetails       *LowTextContrastIssueDetails       `json:"lowTextContrastIssueDetails,omitempty"`
	CorsIssueDetails                  *CorsIssueDetails                  `json:"corsIssueDetails,omitempty"`
}

/*
	An inspector issue reported from the back-end.
*/
type InspectorIssue struct {
	Code    InspectorIssueCode     `json:"code"`
	Details *InspectorIssueDetails `json:"details"`
}

type GetEncodedResponseArgs struct {
	RequestId network.RequestId `json:"requestId"`
	Encoding  string            `json:"encoding"`
	Quality   float64           `json:"quality,omitempty"`
	SizeOnly  bool              `json:"sizeOnly,omitempty"`
}

type GetEncodedResponseVal struct {
	Body         []byte `json:"body,omitempty"`
	OriginalSize int    `json:"originalSize"`
	EncodedSize  int    `json:"encodedSize"`
}

type CheckContrastArgs struct {
	ReportAAA bool `json:"reportAAA,omitempty"`
}
