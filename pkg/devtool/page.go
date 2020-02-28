package devtool

// LifecycleEventType type of LifecycleEvent
type LifecycleEventType string

// DownloadBehavior deny, allow, default
type DownloadBehavior string

// DownloadBehavior
const (
	DownloadBehaviorDeny    DownloadBehavior = "deny"
	DownloadBehaviorAllow   DownloadBehavior = "allow"
	DownloadBehaviorDefault DownloadBehavior = "default"
)

// LifecycleEventType
const (
	DOMContentLoaded              LifecycleEventType = "DOMContentLoaded"
	FirstContentfulPaint          LifecycleEventType = "firstContentfulPaint"
	FirstMeaningfulPaint          LifecycleEventType = "firstMeaningfulPaint"
	FirstMeaningfulPaintCandidate LifecycleEventType = "firstMeaningfulPaintCandidate"
	FirstPaint                    LifecycleEventType = "firstPaint"
	FirstTextPaint                LifecycleEventType = "firstTextPaint"
	Init                          LifecycleEventType = "init"
	Load                          LifecycleEventType = "load"
	NetworkAlmostIdle             LifecycleEventType = "networkAlmostIdle"
	NetworkIdle                   LifecycleEventType = "networkIdle"
)

// NavigationResult https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigate
type NavigationResult struct {
	FrameID   string `json:"frameId"`
	LoaderID  string `json:"loaderId"`
	ErrorText string `json:"errorText"`
}

// NavigationEntry https://chromedevtools.github.io/devtools-protocol/tot/Page#type-NavigationEntry
type NavigationEntry struct {
	ID             int64  `json:"id"`
	URL            string `json:"url"`
	UserTypedURL   string `json:"userTypedURL"`
	Title          string `json:"title"`
	TransitionType string `json:"transitionType"`
}

// LayoutViewport https://chromedevtools.github.io/devtools-protocol/tot/Page#type-LayoutViewport
type LayoutViewport struct {
	PageX        int64 `json:"pageX"`
	PageY        int64 `json:"pageY"`
	ClientWidth  int64 `json:"clientWidth"`
	ClientHeight int64 `json:"clientHeight"`
}

// Viewport https://chromedevtools.github.io/devtools-protocol/tot/Page#type-Viewport
type Viewport struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Scale  float64 `json:"scale"`
}

// VisualViewport https://chromedevtools.github.io/devtools-protocol/tot/Page#type-VisualViewport
type VisualViewport struct {
	OffsetX      float64 `json:"offsetX"`
	OffsetY      float64 `json:"offsetY"`
	PageX        float64 `json:"pageX"`
	PageY        float64 `json:"pageY"`
	ClientWidth  float64 `json:"clientWidth"`
	ClientHeight float64 `json:"clientHeight"`
	Scale        float64 `json:"scale"`
	Zoom         float64 `json:"zoom"`
}

// LayoutMetrics https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getLayoutMetrics
type LayoutMetrics struct {
	LayoutViewport *LayoutViewport `json:"layoutViewport"`
	VisualViewport *VisualViewport `json:"visualViewport"`
	ContentSize    *Rect           `json:"contentSize"`
}

// NavigationHistory https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getNavigationHistory
type NavigationHistory struct {
	CurrentIndex int64              `json:"currentIndex"`
	Entries      []*NavigationEntry `json:"entries"`
}

// LifecycleEvent https://chromedevtools.github.io/devtools-protocol/tot/Page#event-lifecycleEvent
type LifecycleEvent struct {
	FrameID   string  `json:"frameId"`
	LoaderID  string  `json:"loaderId"`
	Name      string  `json:"name"`
	Timestamp float64 `json:"timestamp"`
}

// Frame https://chromedevtools.github.io/devtools-protocol/tot/Page#type-Frame
type Frame struct {
	ID             string `json:"id"`
	ParentID       string `json:"parentId"`
	LoaderID       string `json:"loaderId"`
	Name           string `json:"name"`
	URL            string `json:"url"`
	SecurityOrigin string `json:"securityOrigin"`
	MimeType       string `json:"mimeTypeurl"`
	UnreachableURL string `json:"unreachableUrl"`
}

// FrameTreeResult https://chromedevtools.github.io/devtools-protocol/tot/Page#method-getFrameTree
type FrameTreeResult struct {
	FrameTree *FrameTree `json:"frameTree"`
}

// FrameTree https://chromedevtools.github.io/devtools-protocol/tot/Page#type-FrameTree
type FrameTree struct {
	Frame       *Frame       `json:"frame"`
	ChildFrames []*FrameTree `json:"childFrames"`
}

// Look look for frame with ID
func (f FrameTree) Look(ID string) *Frame {
	if f.Frame.ID == ID {
		return f.Frame
	}
	for _, c := range f.ChildFrames {
		if cf := c.Look(ID); cf != nil {
			return cf
		}
	}
	return nil
}
