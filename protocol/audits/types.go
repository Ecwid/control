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
	RequestId network.RequestId `json:"requestId,omitempty"`
	Url       string            `json:"url"`
}

/*
Information about the frame affected by an inspector issue.
*/
type AffectedFrame struct {
	FrameId common.FrameId `json:"frameId"`
}

/*
 */
type CookieExclusionReason string

/*
 */
type CookieWarningReason string

/*
 */
type CookieOperation string

/*
Represents the category of insight that a cookie issue falls under.
*/
type InsightType string

/*
Information about the suggested solution to a cookie issue.
*/
type CookieIssueInsight struct {
	Type          InsightType `json:"type"`
	TableEntryUrl string      `json:"tableEntryUrl,omitempty"`
}

/*
	This information is currently necessary, as the front-end has a difficult

time finding a specific cookie. With this, we can convey specific error
information without the cookie.
*/
type CookieIssueDetails struct {
	Cookie                 *AffectedCookie         `json:"cookie,omitempty"`
	RawCookieLine          string                  `json:"rawCookieLine,omitempty"`
	CookieWarningReasons   []CookieWarningReason   `json:"cookieWarningReasons"`
	CookieExclusionReasons []CookieExclusionReason `json:"cookieExclusionReasons"`
	Operation              CookieOperation         `json:"operation"`
	SiteForCookies         string                  `json:"siteForCookies,omitempty"`
	CookieUrl              string                  `json:"cookieUrl,omitempty"`
	Request                *AffectedRequest        `json:"request,omitempty"`
	Insight                *CookieIssueInsight     `json:"insight,omitempty"`
}

/*
 */
type PerformanceIssueType string

/*
Details for a performance issue.
*/
type PerformanceIssueDetails struct {
	PerformanceIssueType PerformanceIssueType `json:"performanceIssueType"`
	SourceCodeLocation   *SourceCodeLocation  `json:"sourceCodeLocation,omitempty"`
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

transferred to a context that is not cross-origin isolated.
*/
type SharedArrayBufferIssueDetails struct {
	SourceCodeLocation *SourceCodeLocation        `json:"sourceCodeLocation"`
	IsWarning          bool                       `json:"isWarning"`
	Type               SharedArrayBufferIssueType `json:"type"`
}

/*
	Details for a CORS related issue, e.g. a warning or error related to

CORS RFC1918 enforcement.
*/
type CorsIssueDetails struct {
	CorsErrorStatus        *network.CorsErrorStatus     `json:"corsErrorStatus"`
	IsWarning              bool                         `json:"isWarning"`
	Request                *AffectedRequest             `json:"request"`
	Location               *SourceCodeLocation          `json:"location,omitempty"`
	InitiatorOrigin        string                       `json:"initiatorOrigin,omitempty"`
	ResourceIPAddressSpace network.IPAddressSpace       `json:"resourceIPAddressSpace,omitempty"`
	ClientSecurityState    *network.ClientSecurityState `json:"clientSecurityState,omitempty"`
}

/*
 */
type AttributionReportingIssueType string

/*
 */
type SharedDictionaryError string

/*
 */
type SRIMessageSignatureError string

/*
 */
type UnencodedDigestError string

/*
 */
type ConnectionAllowlistError string

/*
	Details for issues around "Attribution Reporting API" usage.

Explainer: https://github.com/WICG/attribution-reporting-api
*/
type AttributionReportingIssueDetails struct {
	ViolationType    AttributionReportingIssueType `json:"violationType"`
	Request          *AffectedRequest              `json:"request,omitempty"`
	ViolatingNodeId  dom.BackendNodeId             `json:"violatingNodeId,omitempty"`
	InvalidParameter string                        `json:"invalidParameter,omitempty"`
}

/*
	Details for issues about documents in Quirks Mode

or Limited Quirks Mode that affects page layouting.
*/
type QuirksModeIssueDetails struct {
	IsLimitedQuirksMode bool              `json:"isLimitedQuirksMode"`
	DocumentNodeId      dom.BackendNodeId `json:"documentNodeId"`
	Url                 string            `json:"url"`
	FrameId             common.FrameId    `json:"frameId"`
	LoaderId            network.LoaderId  `json:"loaderId"`
}

/*
 */
type SharedDictionaryIssueDetails struct {
	SharedDictionaryError SharedDictionaryError `json:"sharedDictionaryError"`
	Request               *AffectedRequest      `json:"request"`
}

/*
 */
type SRIMessageSignatureIssueDetails struct {
	Error               SRIMessageSignatureError `json:"error"`
	SignatureBase       string                   `json:"signatureBase"`
	IntegrityAssertions []string                 `json:"integrityAssertions"`
	Request             *AffectedRequest         `json:"request"`
}

/*
 */
type UnencodedDigestIssueDetails struct {
	Error   UnencodedDigestError `json:"error"`
	Request *AffectedRequest     `json:"request"`
}

/*
 */
type ConnectionAllowlistIssueDetails struct {
	Error   ConnectionAllowlistError `json:"error"`
	Request *AffectedRequest         `json:"request"`
}

/*
 */
type GenericIssueErrorType string

/*
Depending on the concrete errorType, different properties are set.
*/
type GenericIssueDetails struct {
	ErrorType              GenericIssueErrorType `json:"errorType"`
	FrameId                common.FrameId        `json:"frameId,omitempty"`
	ViolatingNodeId        dom.BackendNodeId     `json:"violatingNodeId,omitempty"`
	ViolatingNodeAttribute string                `json:"violatingNodeAttribute,omitempty"`
	Request                *AffectedRequest      `json:"request,omitempty"`
}

/*
	This issue tracks information needed to print a deprecation message.

https://source.chromium.org/chromium/chromium/src/+/main:third_party/blink/renderer/core/frame/third_party/blink/renderer/core/frame/deprecation/README.md
*/
type DeprecationIssueDetails struct {
	AffectedFrame      *AffectedFrame      `json:"affectedFrame,omitempty"`
	SourceCodeLocation *SourceCodeLocation `json:"sourceCodeLocation"`
	Type               string              `json:"type"`
}

/*
	This issue warns about sites in the redirect chain of a finished navigation

that may be flagged as trackers and have their state cleared if they don't
receive a user interaction. Note that in this context 'site' means eTLD+1.
For example, if the URL `https://example.test:80/bounce` was in the
redirect chain, the site reported would be `example.test`.
*/
type BounceTrackingIssueDetails struct {
	TrackingSites []string `json:"trackingSites"`
}

/*
	This issue warns about third-party sites that are accessing cookies on the

current page, and have been permitted due to having a global metadata grant.
Note that in this context 'site' means eTLD+1. For example, if the URL
`https://example.test:80/web_page` was accessing cookies, the site reported
would be `example.test`.
*/
type CookieDeprecationMetadataIssueDetails struct {
	AllowedSites     []string        `json:"allowedSites"`
	OptOutPercentage float64         `json:"optOutPercentage"`
	IsOptOutTopLevel bool            `json:"isOptOutTopLevel"`
	Operation        CookieOperation `json:"operation"`
}

/*
 */
type ClientHintIssueReason string

/*
 */
type FederatedAuthRequestIssueDetails struct {
	FederatedAuthRequestIssueReason FederatedAuthRequestIssueReason `json:"federatedAuthRequestIssueReason"`
}

/*
	Represents the failure reason when a federated authentication reason fails.

Should be updated alongside RequestIdTokenStatus in
third_party/blink/public/mojom/devtools/inspector_issue.mojom to include
all cases except for success.
*/
type FederatedAuthRequestIssueReason string

/*
 */
type FederatedAuthUserInfoRequestIssueDetails struct {
	FederatedAuthUserInfoRequestIssueReason FederatedAuthUserInfoRequestIssueReason `json:"federatedAuthUserInfoRequestIssueReason"`
}

/*
	Represents the failure reason when a getUserInfo() call fails.

Should be updated alongside FederatedAuthUserInfoRequestResult in
third_party/blink/public/mojom/devtools/inspector_issue.mojom.
*/
type FederatedAuthUserInfoRequestIssueReason string

/*
 */
type EmailVerificationRequestIssueDetails struct {
	EmailVerificationRequestIssueReason EmailVerificationRequestIssueReason `json:"emailVerificationRequestIssueReason"`
}

/*
	Represents the failure reason when an email verification request fails.

Should be updated alongside EmailVerificationRequestResult in
third_party/blink/public/mojom/devtools/inspector_issue.mojom.
*/
type EmailVerificationRequestIssueReason string

/*
	This issue tracks client hints related issues. It's used to deprecate old

features, encourage the use of new ones, and provide general guidance.
*/
type ClientHintIssueDetails struct {
	SourceCodeLocation    *SourceCodeLocation   `json:"sourceCodeLocation"`
	ClientHintIssueReason ClientHintIssueReason `json:"clientHintIssueReason"`
}

/*
 */
type FailedRequestInfo struct {
	Url            string            `json:"url"`
	FailureMessage string            `json:"failureMessage"`
	RequestId      network.RequestId `json:"requestId,omitempty"`
}

/*
 */
type PartitioningBlobURLInfo string

/*
 */
type PartitioningBlobURLIssueDetails struct {
	Url                     string                  `json:"url"`
	PartitioningBlobURLInfo PartitioningBlobURLInfo `json:"partitioningBlobURLInfo"`
}

/*
 */
type ElementAccessibilityIssueReason string

/*
This issue warns about errors in the select or summary element content model.
*/
type ElementAccessibilityIssueDetails struct {
	NodeId                          dom.BackendNodeId               `json:"nodeId"`
	ElementAccessibilityIssueReason ElementAccessibilityIssueReason `json:"elementAccessibilityIssueReason"`
	HasDisallowedAttributes         bool                            `json:"hasDisallowedAttributes"`
}

/*
 */
type StyleSheetLoadingIssueReason string

/*
This issue warns when a referenced stylesheet couldn't be loaded.
*/
type StylesheetLoadingIssueDetails struct {
	SourceCodeLocation           *SourceCodeLocation          `json:"sourceCodeLocation"`
	StyleSheetLoadingIssueReason StyleSheetLoadingIssueReason `json:"styleSheetLoadingIssueReason"`
	FailedRequestInfo            *FailedRequestInfo           `json:"failedRequestInfo,omitempty"`
}

/*
 */
type PropertyRuleIssueReason string

/*
	This issue warns about errors in property rules that lead to property

registrations being ignored.
*/
type PropertyRuleIssueDetails struct {
	SourceCodeLocation      *SourceCodeLocation     `json:"sourceCodeLocation"`
	PropertyRuleIssueReason PropertyRuleIssueReason `json:"propertyRuleIssueReason"`
	PropertyValue           string                  `json:"propertyValue,omitempty"`
}

/*
 */
type UserReidentificationIssueType string

/*
	This issue warns about uses of APIs that may be considered misuse to

re-identify users.
*/
type UserReidentificationIssueDetails struct {
	Type               UserReidentificationIssueType `json:"type"`
	Request            *AffectedRequest              `json:"request,omitempty"`
	SourceCodeLocation *SourceCodeLocation           `json:"sourceCodeLocation,omitempty"`
}

/*
 */
type PermissionElementIssueType string

/*
This issue warns about improper usage of the <permission> element.
*/
type PermissionElementIssueDetails struct {
	IssueType              PermissionElementIssueType `json:"issueType"`
	Type                   string                     `json:"type,omitempty"`
	NodeId                 dom.BackendNodeId          `json:"nodeId,omitempty"`
	IsWarning              bool                       `json:"isWarning,omitempty"`
	PermissionName         string                     `json:"permissionName,omitempty"`
	OccluderNodeInfo       string                     `json:"occluderNodeInfo,omitempty"`
	OccluderParentNodeInfo string                     `json:"occluderParentNodeInfo,omitempty"`
	DisableReason          string                     `json:"disableReason,omitempty"`
}

/*
	The issue warns about blocked calls to privacy sensitive APIs via the

Selective Permissions Intervention.
*/
type SelectivePermissionsInterventionIssueDetails struct {
	ApiName    string              `json:"apiName"`
	AdAncestry *network.AdAncestry `json:"adAncestry"`
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`
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
	CookieIssueDetails                           *CookieIssueDetails                           `json:"cookieIssueDetails,omitempty"`
	MixedContentIssueDetails                     *MixedContentIssueDetails                     `json:"mixedContentIssueDetails,omitempty"`
	BlockedByResponseIssueDetails                *BlockedByResponseIssueDetails                `json:"blockedByResponseIssueDetails,omitempty"`
	HeavyAdIssueDetails                          *HeavyAdIssueDetails                          `json:"heavyAdIssueDetails,omitempty"`
	ContentSecurityPolicyIssueDetails            *ContentSecurityPolicyIssueDetails            `json:"contentSecurityPolicyIssueDetails,omitempty"`
	SharedArrayBufferIssueDetails                *SharedArrayBufferIssueDetails                `json:"sharedArrayBufferIssueDetails,omitempty"`
	CorsIssueDetails                             *CorsIssueDetails                             `json:"corsIssueDetails,omitempty"`
	AttributionReportingIssueDetails             *AttributionReportingIssueDetails             `json:"attributionReportingIssueDetails,omitempty"`
	QuirksModeIssueDetails                       *QuirksModeIssueDetails                       `json:"quirksModeIssueDetails,omitempty"`
	PartitioningBlobURLIssueDetails              *PartitioningBlobURLIssueDetails              `json:"partitioningBlobURLIssueDetails,omitempty"`
	GenericIssueDetails                          *GenericIssueDetails                          `json:"genericIssueDetails,omitempty"`
	DeprecationIssueDetails                      *DeprecationIssueDetails                      `json:"deprecationIssueDetails,omitempty"`
	ClientHintIssueDetails                       *ClientHintIssueDetails                       `json:"clientHintIssueDetails,omitempty"`
	FederatedAuthRequestIssueDetails             *FederatedAuthRequestIssueDetails             `json:"federatedAuthRequestIssueDetails,omitempty"`
	BounceTrackingIssueDetails                   *BounceTrackingIssueDetails                   `json:"bounceTrackingIssueDetails,omitempty"`
	CookieDeprecationMetadataIssueDetails        *CookieDeprecationMetadataIssueDetails        `json:"cookieDeprecationMetadataIssueDetails,omitempty"`
	StylesheetLoadingIssueDetails                *StylesheetLoadingIssueDetails                `json:"stylesheetLoadingIssueDetails,omitempty"`
	PropertyRuleIssueDetails                     *PropertyRuleIssueDetails                     `json:"propertyRuleIssueDetails,omitempty"`
	FederatedAuthUserInfoRequestIssueDetails     *FederatedAuthUserInfoRequestIssueDetails     `json:"federatedAuthUserInfoRequestIssueDetails,omitempty"`
	SharedDictionaryIssueDetails                 *SharedDictionaryIssueDetails                 `json:"sharedDictionaryIssueDetails,omitempty"`
	ElementAccessibilityIssueDetails             *ElementAccessibilityIssueDetails             `json:"elementAccessibilityIssueDetails,omitempty"`
	SriMessageSignatureIssueDetails              *SRIMessageSignatureIssueDetails              `json:"sriMessageSignatureIssueDetails,omitempty"`
	UnencodedDigestIssueDetails                  *UnencodedDigestIssueDetails                  `json:"unencodedDigestIssueDetails,omitempty"`
	ConnectionAllowlistIssueDetails              *ConnectionAllowlistIssueDetails              `json:"connectionAllowlistIssueDetails,omitempty"`
	UserReidentificationIssueDetails             *UserReidentificationIssueDetails             `json:"userReidentificationIssueDetails,omitempty"`
	PermissionElementIssueDetails                *PermissionElementIssueDetails                `json:"permissionElementIssueDetails,omitempty"`
	PerformanceIssueDetails                      *PerformanceIssueDetails                      `json:"performanceIssueDetails,omitempty"`
	SelectivePermissionsInterventionIssueDetails *SelectivePermissionsInterventionIssueDetails `json:"selectivePermissionsInterventionIssueDetails,omitempty"`
	EmailVerificationRequestIssueDetails         *EmailVerificationRequestIssueDetails         `json:"emailVerificationRequestIssueDetails,omitempty"`
}

/*
	A unique id for a DevTools inspector issue. Allows other entities (e.g.

exceptions, CDP message, console messages, etc.) to reference an issue.
*/
type IssueId string

/*
An inspector issue reported from the back-end.
*/
type InspectorIssue struct {
	Code    InspectorIssueCode     `json:"code"`
	Details *InspectorIssueDetails `json:"details"`
	IssueId IssueId                `json:"issueId,omitempty"`
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

type CheckFormsIssuesVal struct {
	FormIssues []*GenericIssueDetails `json:"formIssues"`
}
