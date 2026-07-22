package pwa

import (
	"github.com/ecwid/control/protocol/target"
)

/*
	The following types are the replica of

https://crsrc.org/c/chrome/browser/web_applications/proto/web_app_os_integration_state.proto;drc=9910d3be894c8f142c977ba1023f30a656bc13fc;l=67
*/
type FileHandlerAccept struct {
	MediaType      string   `json:"mediaType"`
	FileExtensions []string `json:"fileExtensions"`
}

/*
 */
type FileHandler struct {
	Action      string               `json:"action"`
	Accepts     []*FileHandlerAccept `json:"accepts"`
	DisplayName string               `json:"displayName"`
}

/*
If user prefers opening the app in browser or an app window.
*/
type DisplayMode string

type GetOsAppStateArgs struct {
	ManifestId string `json:"manifestId"`
}

type GetOsAppStateVal struct {
	BadgeCount   int            `json:"badgeCount"`
	FileHandlers []*FileHandler `json:"fileHandlers"`
}

type InstallArgs struct {
	ManifestId            string `json:"manifestId"`
	InstallUrlOrBundleUrl string `json:"installUrlOrBundleUrl,omitempty"`
}

type UninstallArgs struct {
	ManifestId string `json:"manifestId"`
}

type LaunchArgs struct {
	ManifestId string `json:"manifestId"`
	Url        string `json:"url,omitempty"`
}

type LaunchVal struct {
	TargetId target.TargetID `json:"targetId"`
}

type LaunchFilesInAppArgs struct {
	ManifestId string   `json:"manifestId"`
	Files      []string `json:"files"`
}

type LaunchFilesInAppVal struct {
	TargetIds []target.TargetID `json:"targetIds"`
}

type OpenCurrentPageInAppArgs struct {
	ManifestId string `json:"manifestId"`
}

type ChangeAppUserSettingsArgs struct {
	ManifestId    string      `json:"manifestId"`
	LinkCapturing bool        `json:"linkCapturing,omitempty"`
	DisplayMode   DisplayMode `json:"displayMode,omitempty"`
}
