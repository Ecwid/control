package preload

import (
	"github.com/ecwid/control/protocol/common"
	"github.com/ecwid/control/protocol/network"
)

/*
Upsert. Currently, it is only emitted when a rule set added.
*/
type RuleSetUpdated struct {
	RuleSet *RuleSet `json:"ruleSet"`
}

/*
 */
type RuleSetRemoved struct {
	Id RuleSetId `json:"id"`
}

/*
Fired when a preload enabled state is updated.
*/
type PreloadEnabledStateUpdated struct {
	DisabledByPreference                        bool `json:"disabledByPreference"`
	DisabledByDataSaver                         bool `json:"disabledByDataSaver"`
	DisabledByBatterySaver                      bool `json:"disabledByBatterySaver"`
	DisabledByHoldbackPrefetchSpeculationRules  bool `json:"disabledByHoldbackPrefetchSpeculationRules"`
	DisabledByHoldbackPrerenderSpeculationRules bool `json:"disabledByHoldbackPrerenderSpeculationRules"`
}

/*
Fired when a prefetch attempt is updated.
*/
type PrefetchStatusUpdated struct {
	Key               *PreloadingAttemptKey `json:"key"`
	PipelineId        PreloadPipelineId     `json:"pipelineId"`
	InitiatingFrameId common.FrameId        `json:"initiatingFrameId"`
	PrefetchUrl       string                `json:"prefetchUrl"`
	Status            PreloadingStatus      `json:"status"`
	PrefetchStatus    PrefetchStatus        `json:"prefetchStatus"`
	RequestId         network.RequestId     `json:"requestId"`
}

/*
Fired when a prerender attempt is updated.
*/
type PrerenderStatusUpdated struct {
	Key                     *PreloadingAttemptKey         `json:"key"`
	PipelineId              PreloadPipelineId             `json:"pipelineId"`
	Status                  PreloadingStatus              `json:"status"`
	PrerenderStatus         PrerenderFinalStatus          `json:"prerenderStatus,omitempty"`
	DisallowedMojoInterface string                        `json:"disallowedMojoInterface,omitempty"`
	MismatchedHeaders       []*PrerenderMismatchedHeaders `json:"mismatchedHeaders,omitempty"`
}

/*
Send a list of sources for all preloading attempts in a document.
*/
type PreloadingAttemptSourcesUpdated struct {
	LoaderId                 network.LoaderId           `json:"loaderId"`
	PreloadingAttemptSources []*PreloadingAttemptSource `json:"preloadingAttemptSources"`
}
