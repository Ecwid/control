package preload

import (
	"github.com/ecwid/control/protocol/dom"
	"github.com/ecwid/control/protocol/network"
)

/*
Unique id
*/
type RuleSetId string

/*
Corresponds to SpeculationRuleSet
*/
type RuleSet struct {
	Id            RuleSetId         `json:"id"`
	LoaderId      network.LoaderId  `json:"loaderId"`
	SourceText    string            `json:"sourceText"`
	BackendNodeId dom.BackendNodeId `json:"backendNodeId,omitempty"`
	Url           string            `json:"url,omitempty"`
	RequestId     network.RequestId `json:"requestId,omitempty"`
	ErrorType     RuleSetErrorType  `json:"errorType,omitempty"`
	Tag           string            `json:"tag,omitempty"`
}

/*
 */
type RuleSetErrorType string

/*
	The type of preloading attempted. It corresponds to

mojom::SpeculationAction (although PrefetchWithSubresources is omitted as it
isn't being used by clients).
*/
type SpeculationAction string

/*
	Corresponds to mojom::SpeculationTargetHint.

See https://github.com/WICG/nav-speculation/blob/main/triggers.md#window-name-targeting-hints
*/
type SpeculationTargetHint string

/*
	A key that identifies a preloading attempt.

The url used is the url specified by the trigger (i.e. the initial URL), and
not the final url that is navigated to. For example, prerendering allows
same-origin main frame navigations during the attempt, but the attempt is
still keyed with the initial URL.
*/
type PreloadingAttemptKey struct {
	LoaderId       network.LoaderId      `json:"loaderId"`
	Action         SpeculationAction     `json:"action"`
	Url            string                `json:"url"`
	FormSubmission bool                  `json:"formSubmission,omitempty"`
	TargetHint     SpeculationTargetHint `json:"targetHint,omitempty"`
}

/*
	Lists sources for a preloading attempt, specifically the ids of rule sets

that had a speculation rule that triggered the attempt, and the
BackendNodeIds of <a href> or <area href> elements that triggered the
attempt (in the case of attempts triggered by a document rule). It is
possible for multiple rule sets and links to trigger a single attempt.
*/
type PreloadingAttemptSource struct {
	Key        *PreloadingAttemptKey `json:"key"`
	RuleSetIds []RuleSetId           `json:"ruleSetIds"`
	NodeIds    []dom.BackendNodeId   `json:"nodeIds"`
}

/*
	Chrome manages different types of preloads together using a

concept of preloading pipeline. For example, if a site uses a
SpeculationRules for prerender, Chrome first starts a prefetch and
then upgrades it to prerender.

CDP events for them are emitted separately but they share
`PreloadPipelineId`.
*/
type PreloadPipelineId string

/*
List of FinalStatus reasons for Prerender2.
*/
type PrerenderFinalStatus string

/*
	Preloading status values, see also PreloadingTriggeringOutcome. This

status is shared by prefetchStatusUpdated and prerenderStatusUpdated.
*/
type PreloadingStatus string

/*
	TODO(https://crbug.com/1384419): revisit the list of PrefetchStatus and

filter out the ones that aren't necessary to the developers.
*/
type PrefetchStatus string

/*
Information of headers to be displayed when the header mismatch occurred.
*/
type PrerenderMismatchedHeaders struct {
	HeaderName      string `json:"headerName"`
	InitialValue    string `json:"initialValue,omitempty"`
	ActivationValue string `json:"activationValue,omitempty"`
}
