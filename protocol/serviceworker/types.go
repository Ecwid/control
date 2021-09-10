package serviceworker

import (
	"github.com/ecwid/control/protocol/target"
)

/*

 */
type RegistrationID string

/*
	ServiceWorker registration.
*/
type ServiceWorkerRegistration struct {
	RegistrationId RegistrationID `json:"registrationId"`
	ScopeURL       string         `json:"scopeURL"`
	IsDeleted      bool           `json:"isDeleted"`
}

/*

 */
type ServiceWorkerVersionRunningStatus string

/*

 */
type ServiceWorkerVersionStatus string

/*
	ServiceWorker version.
*/
type ServiceWorkerVersion struct {
	VersionId          string                            `json:"versionId"`
	RegistrationId     RegistrationID                    `json:"registrationId"`
	ScriptURL          string                            `json:"scriptURL"`
	RunningStatus      ServiceWorkerVersionRunningStatus `json:"runningStatus"`
	Status             ServiceWorkerVersionStatus        `json:"status"`
	ScriptLastModified float64                           `json:"scriptLastModified,omitempty"`
	ScriptResponseTime float64                           `json:"scriptResponseTime,omitempty"`
	ControlledClients  []target.TargetID                 `json:"controlledClients,omitempty"`
	TargetId           target.TargetID                   `json:"targetId,omitempty"`
}

/*
	ServiceWorker error message.
*/
type ServiceWorkerErrorMessage struct {
	ErrorMessage   string         `json:"errorMessage"`
	RegistrationId RegistrationID `json:"registrationId"`
	VersionId      string         `json:"versionId"`
	SourceURL      string         `json:"sourceURL"`
	LineNumber     int            `json:"lineNumber"`
	ColumnNumber   int            `json:"columnNumber"`
}

type DeliverPushMessageArgs struct {
	Origin         string         `json:"origin"`
	RegistrationId RegistrationID `json:"registrationId"`
	Data           string         `json:"data"`
}

type DispatchSyncEventArgs struct {
	Origin         string         `json:"origin"`
	RegistrationId RegistrationID `json:"registrationId"`
	Tag            string         `json:"tag"`
	LastChance     bool           `json:"lastChance"`
}

type DispatchPeriodicSyncEventArgs struct {
	Origin         string         `json:"origin"`
	RegistrationId RegistrationID `json:"registrationId"`
	Tag            string         `json:"tag"`
}

type InspectWorkerArgs struct {
	VersionId string `json:"versionId"`
}

type SetForceUpdateOnPageLoadArgs struct {
	ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"`
}

type SkipWaitingArgs struct {
	ScopeURL string `json:"scopeURL"`
}

type StartWorkerArgs struct {
	ScopeURL string `json:"scopeURL"`
}

type StopWorkerArgs struct {
	VersionId string `json:"versionId"`
}

type UnregisterArgs struct {
	ScopeURL string `json:"scopeURL"`
}

type UpdateRegistrationArgs struct {
	ScopeURL string `json:"scopeURL"`
}
