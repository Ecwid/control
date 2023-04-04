package security

import (
	"github.com/ecwid/control/protocol/common"
)

/*
An internal certificate ID value.
*/
type CertificateId int

/*
	A description of mixed content (HTTP resources on HTTPS pages), as defined by

https://www.w3.org/TR/mixed-content/#categories
*/
type MixedContentType string

/*
The security level of a page or resource.
*/
type SecurityState string

/*
Details about the security state of the page certificate.
*/
type CertificateSecurityState struct {
	Protocol                    string                `json:"protocol"`
	KeyExchange                 string                `json:"keyExchange"`
	KeyExchangeGroup            string                `json:"keyExchangeGroup,omitempty"`
	Cipher                      string                `json:"cipher"`
	Mac                         string                `json:"mac,omitempty"`
	Certificate                 []string              `json:"certificate"`
	SubjectName                 string                `json:"subjectName"`
	Issuer                      string                `json:"issuer"`
	ValidFrom                   common.TimeSinceEpoch `json:"validFrom"`
	ValidTo                     common.TimeSinceEpoch `json:"validTo"`
	CertificateNetworkError     string                `json:"certificateNetworkError,omitempty"`
	CertificateHasWeakSignature bool                  `json:"certificateHasWeakSignature"`
	CertificateHasSha1Signature bool                  `json:"certificateHasSha1Signature"`
	ModernSSL                   bool                  `json:"modernSSL"`
	ObsoleteSslProtocol         bool                  `json:"obsoleteSslProtocol"`
	ObsoleteSslKeyExchange      bool                  `json:"obsoleteSslKeyExchange"`
	ObsoleteSslCipher           bool                  `json:"obsoleteSslCipher"`
	ObsoleteSslSignature        bool                  `json:"obsoleteSslSignature"`
}

/*
 */
type SafetyTipStatus string

/*
 */
type SafetyTipInfo struct {
	SafetyTipStatus SafetyTipStatus `json:"safetyTipStatus"`
	SafeUrl         string          `json:"safeUrl,omitempty"`
}

/*
Security state information about the page.
*/
type VisibleSecurityState struct {
	SecurityState            SecurityState             `json:"securityState"`
	CertificateSecurityState *CertificateSecurityState `json:"certificateSecurityState,omitempty"`
	SafetyTipInfo            *SafetyTipInfo            `json:"safetyTipInfo,omitempty"`
	SecurityStateIssueIds    []string                  `json:"securityStateIssueIds"`
}

/*
An explanation of an factor contributing to the security state.
*/
type SecurityStateExplanation struct {
	SecurityState    SecurityState    `json:"securityState"`
	Title            string           `json:"title"`
	Summary          string           `json:"summary"`
	Description      string           `json:"description"`
	MixedContentType MixedContentType `json:"mixedContentType"`
	Certificate      []string         `json:"certificate"`
	Recommendations  []string         `json:"recommendations,omitempty"`
}

/*
	The action to take when a certificate error occurs. continue will continue processing the

request and cancel will cancel the request.
*/
type CertificateErrorAction string

type SetIgnoreCertificateErrorsArgs struct {
	Ignore bool `json:"ignore"`
}
