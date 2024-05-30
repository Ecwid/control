package key

type Definition struct {
	KeyCode      int
	ShiftKeyCode int
	Key          string
	ShiftKey     string
	Code         string
	Text         string
	ShiftText    string
	Location     int
}

const (
	Control rune = iota + 255
	Abort
	Help
	Backspace
	Tab
	Enter
	ShiftLeft
	ShiftRight
	ControlLeft
	ControlRight
	AltLeft
	AltRight
	Pause
	CapsLock
	Escape
	Convert
	NonConvert
	Space
	PageUp
	PageDown
	End
	Home
	ArrowLeft
	ArrowUp
	ArrowRight
	ArrowDown
	Select
	Open
	PrintScreen
	Insert
	Delete
	MetaLeft
	MetaRight
	ContextMenu
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24
	NumLock
	ScrollLock
	AudioVolumeMute
	AudioVolumeDown
	AudioVolumeUp
	MediaTrackNext
	MediaTrackPrevious
	MediaStop
	MediaPlayPause
	AltGraph
	Props
	Cancel
	Clear
	Shift
	Alt
	Accept
	ModeChange
	Print
	Execute
	Meta
	Attn
	CrSel
	ExSel
	EraseEof
	Play
	ZoomOut
	SoftLeft
	SoftRight
	Camera
	Call
	EndCall
	VolumeDown
	VolumeUp
)

var Keys = map[rune]Definition{
	Abort:              {KeyCode: 3, Code: "Abort", Key: "Cancel"},
	Help:               {KeyCode: 6, Code: "Help", Key: "Help"},
	Backspace:          {KeyCode: 8, Code: "Backspace", Key: "Backspace"},
	Tab:                {KeyCode: 9, Code: "Tab", Key: "Tab"},
	Enter:              {KeyCode: 13, Code: "Enter", Key: "Enter", Text: "\r"},
	ShiftLeft:          {KeyCode: 16, Code: "ShiftLeft", Key: "Shift", Location: 1},
	ShiftRight:         {KeyCode: 16, Code: "ShiftRight", Key: "Shift", Location: 2},
	ControlLeft:        {KeyCode: 17, Code: "ControlLeft", Key: "Control", Location: 1},
	ControlRight:       {KeyCode: 17, Code: "ControlRight", Key: "Control", Location: 2},
	AltLeft:            {KeyCode: 18, Code: "AltLeft", Key: "Alt", Location: 1},
	AltRight:           {KeyCode: 18, Code: "AltRight", Key: "Alt", Location: 2},
	Pause:              {KeyCode: 19, Code: "Pause", Key: "Pause"},
	CapsLock:           {KeyCode: 20, Code: "CapsLock", Key: "CapsLock"},
	Escape:             {KeyCode: 27, Code: "Escape", Key: "Escape"},
	Convert:            {KeyCode: 28, Code: "Convert", Key: "Convert"},
	NonConvert:         {KeyCode: 29, Code: "NonConvert", Key: "NonConvert"},
	Space:              {KeyCode: 32, Code: "Space", Key: " "},
	PageUp:             {KeyCode: 33, Code: "PageUp", Key: "PageUp"},
	PageDown:           {KeyCode: 34, Code: "PageDown", Key: "PageDown"},
	End:                {KeyCode: 35, Code: "End", Key: "End"},
	Home:               {KeyCode: 36, Code: "Home", Key: "Home"},
	ArrowLeft:          {KeyCode: 37, Code: "ArrowLeft", Key: "ArrowLeft"},
	ArrowUp:            {KeyCode: 38, Code: "ArrowUp", Key: "ArrowUp"},
	ArrowRight:         {KeyCode: 39, Code: "ArrowRight", Key: "ArrowRight"},
	ArrowDown:          {KeyCode: 40, Code: "ArrowDown", Key: "ArrowDown"},
	Select:             {KeyCode: 41, Code: "Select", Key: "Select"},
	Open:               {KeyCode: 43, Code: "Open", Key: "Execute"},
	PrintScreen:        {KeyCode: 44, Code: "PrintScreen", Key: "PrintScreen"},
	Insert:             {KeyCode: 45, Code: "Insert", Key: "Insert"},
	Delete:             {KeyCode: 46, Code: "Delete", Key: "Delete"},
	MetaLeft:           {KeyCode: 91, Code: "MetaLeft", Key: "Meta", Location: 1},
	MetaRight:          {KeyCode: 92, Code: "MetaRight", Key: "Meta", Location: 2},
	ContextMenu:        {KeyCode: 93, Code: "ContextMenu", Key: "ContextMenu"},
	F1:                 {KeyCode: 112, Code: "F1", Key: "F1"},
	F2:                 {KeyCode: 113, Code: "F2", Key: "F2"},
	F3:                 {KeyCode: 114, Code: "F3", Key: "F3"},
	F4:                 {KeyCode: 115, Code: "F4", Key: "F4"},
	F5:                 {KeyCode: 116, Code: "F5", Key: "F5"},
	F6:                 {KeyCode: 117, Code: "F6", Key: "F6"},
	F7:                 {KeyCode: 118, Code: "F7", Key: "F7"},
	F8:                 {KeyCode: 119, Code: "F8", Key: "F8"},
	F9:                 {KeyCode: 120, Code: "F9", Key: "F9"},
	F10:                {KeyCode: 121, Code: "F10", Key: "F10"},
	F11:                {KeyCode: 122, Code: "F11", Key: "F11"},
	F12:                {KeyCode: 123, Code: "F12", Key: "F12"},
	F13:                {KeyCode: 124, Code: "F13", Key: "F13"},
	F14:                {KeyCode: 125, Code: "F14", Key: "F14"},
	F15:                {KeyCode: 126, Code: "F15", Key: "F15"},
	F16:                {KeyCode: 127, Code: "F16", Key: "F16"},
	F17:                {KeyCode: 128, Code: "F17", Key: "F17"},
	F18:                {KeyCode: 129, Code: "F18", Key: "F18"},
	F19:                {KeyCode: 130, Code: "F19", Key: "F19"},
	F20:                {KeyCode: 131, Code: "F20", Key: "F20"},
	F21:                {KeyCode: 132, Code: "F21", Key: "F21"},
	F22:                {KeyCode: 133, Code: "F22", Key: "F22"},
	F23:                {KeyCode: 134, Code: "F23", Key: "F23"},
	F24:                {KeyCode: 135, Code: "F24", Key: "F24"},
	NumLock:            {KeyCode: 144, Code: "NumLock", Key: "NumLock"},
	ScrollLock:         {KeyCode: 145, Code: "ScrollLock", Key: "ScrollLock"},
	AudioVolumeMute:    {KeyCode: 173, Code: "AudioVolumeMute", Key: "AudioVolumeMute"},
	AudioVolumeDown:    {KeyCode: 174, Code: "AudioVolumeDown", Key: "AudioVolumeDown"},
	AudioVolumeUp:      {KeyCode: 175, Code: "AudioVolumeUp", Key: "AudioVolumeUp"},
	MediaTrackNext:     {KeyCode: 176, Code: "MediaTrackNext", Key: "MediaTrackNext"},
	MediaTrackPrevious: {KeyCode: 177, Code: "MediaTrackPrevious", Key: "MediaTrackPrevious"},
	MediaStop:          {KeyCode: 178, Code: "MediaStop", Key: "MediaStop"},
	MediaPlayPause:     {KeyCode: 179, Code: "MediaPlayPause", Key: "MediaPlayPause"},
	AltGraph:           {KeyCode: 225, Code: "AltGraph", Key: "AltGraph"},
	Props:              {KeyCode: 247, Code: "Props", Key: "CrSel"},
	Cancel:             {KeyCode: 3, Key: "Cancel", Code: "Abort"},
	Clear:              {KeyCode: 12, Key: "Clear", Code: "Numpad5", Location: 3},
	Shift:              {KeyCode: 16, Key: "Shift", Code: "ShiftLeft", Location: 1},
	Alt:                {KeyCode: 18, Key: "Alt", Code: "AltLeft", Location: 1},
	Accept:             {KeyCode: 30, Key: "Accept"},
	ModeChange:         {KeyCode: 31, Key: "ModeChange"},
	Print:              {KeyCode: 42, Key: "Print"},
	Execute:            {KeyCode: 43, Key: "Execute", Code: "Open"},
	Meta:               {KeyCode: 91, Key: "Meta", Code: "MetaLeft", Location: 1},
	Attn:               {KeyCode: 246, Key: "Attn"},
	CrSel:              {KeyCode: 247, Key: "CrSel", Code: "Props"},
	ExSel:              {KeyCode: 248, Key: "ExSel"},
	EraseEof:           {KeyCode: 249, Key: "EraseEof"},
	Play:               {KeyCode: 250, Key: "Play"},
	ZoomOut:            {KeyCode: 251, Key: "ZoomOut"},
	Camera:             {KeyCode: 44, Key: "Camera", Code: "Camera", Location: 4},
	EndCall:            {KeyCode: 95, Key: "EndCall", Code: "EndCall", Location: 4},
	VolumeDown:         {KeyCode: 182, Key: "VolumeDown", Code: "VolumeDown", Location: 4},
	VolumeUp:           {KeyCode: 183, Key: "VolumeUp", Code: "VolumeUp", Location: 4},
	Control:            {KeyCode: 17, Key: "Control", Code: "ControlLeft", Location: 1},
	'0':                {KeyCode: 48, Key: "0", Code: "Digit0"},
	'1':                {KeyCode: 49, Key: "1", Code: "Digit1"},
	'2':                {KeyCode: 50, Key: "2", Code: "Digit2"},
	'3':                {KeyCode: 51, Key: "3", Code: "Digit3"},
	'4':                {KeyCode: 52, Key: "4", Code: "Digit4"},
	'5':                {KeyCode: 53, Key: "5", Code: "Digit5"},
	'6':                {KeyCode: 54, Key: "6", Code: "Digit6"},
	'7':                {KeyCode: 55, Key: "7", Code: "Digit7"},
	'8':                {KeyCode: 56, Key: "8", Code: "Digit8"},
	'9':                {KeyCode: 57, Key: "9", Code: "Digit9"},
	'\r':               {KeyCode: 13, Code: "Enter", Key: "Enter", Text: "\r"},
	'\n':               {KeyCode: 13, Code: "Enter", Key: "Enter", Text: "\r"},
	' ':                {KeyCode: 32, Key: " ", Code: "Space"},
	'a':                {KeyCode: 65, Key: "a", Code: "KeyA"},
	'b':                {KeyCode: 66, Key: "b", Code: "KeyB"},
	'c':                {KeyCode: 67, Key: "c", Code: "KeyC"},
	'd':                {KeyCode: 68, Key: "d", Code: "KeyD"},
	'e':                {KeyCode: 69, Key: "e", Code: "KeyE"},
	'f':                {KeyCode: 70, Key: "f", Code: "KeyF"},
	'g':                {KeyCode: 71, Key: "g", Code: "KeyG"},
	'h':                {KeyCode: 72, Key: "h", Code: "KeyH"},
	'i':                {KeyCode: 73, Key: "i", Code: "KeyI"},
	'j':                {KeyCode: 74, Key: "j", Code: "KeyJ"},
	'k':                {KeyCode: 75, Key: "k", Code: "KeyK"},
	'l':                {KeyCode: 76, Key: "l", Code: "KeyL"},
	'm':                {KeyCode: 77, Key: "m", Code: "KeyM"},
	'n':                {KeyCode: 78, Key: "n", Code: "KeyN"},
	'o':                {KeyCode: 79, Key: "o", Code: "KeyO"},
	'p':                {KeyCode: 80, Key: "p", Code: "KeyP"},
	'q':                {KeyCode: 81, Key: "q", Code: "KeyQ"},
	'r':                {KeyCode: 82, Key: "r", Code: "KeyR"},
	's':                {KeyCode: 83, Key: "s", Code: "KeyS"},
	't':                {KeyCode: 84, Key: "t", Code: "KeyT"},
	'u':                {KeyCode: 85, Key: "u", Code: "KeyU"},
	'v':                {KeyCode: 86, Key: "v", Code: "KeyV"},
	'w':                {KeyCode: 87, Key: "w", Code: "KeyW"},
	'x':                {KeyCode: 88, Key: "x", Code: "KeyX"},
	'y':                {KeyCode: 89, Key: "y", Code: "KeyY"},
	'z':                {KeyCode: 90, Key: "z", Code: "KeyZ"},
	'*':                {KeyCode: 106, Key: "*", Code: "NumpadMultiply", Location: 3},
	'+':                {KeyCode: 107, Key: "+", Code: "NumpadAdd", Location: 3},
	'-':                {KeyCode: 109, Key: "-", Code: "NumpadSubtract", Location: 3},
	'/':                {KeyCode: 111, Key: "/", Code: "NumpadDivide", Location: 3},
	';':                {KeyCode: 186, Key: ";", Code: "Semicolon"},
	'=':                {KeyCode: 187, Key: "=", Code: "Equal"},
	',':                {KeyCode: 188, Key: ",", Code: "Comma"},
	'.':                {KeyCode: 190, Key: ".", Code: "Period"},
	'`':                {KeyCode: 192, Key: "`", Code: "Backquote"},
	'[':                {KeyCode: 219, Key: "[", Code: "BracketLeft"},
	'\\':               {KeyCode: 220, Key: "\\", Code: "Backslash"},
	']':                {KeyCode: 221, Key: "]", Code: "BracketRight"},
	'\'':               {KeyCode: 222, Key: "'", Code: "Quote"},
	')':                {KeyCode: 48, Key: ")", Code: "Digit0"},
	'!':                {KeyCode: 49, Key: "!", Code: "Digit1"},
	'@':                {KeyCode: 50, Key: "@", Code: "Digit2"},
	'#':                {KeyCode: 51, Key: "#", Code: "Digit3"},
	'$':                {KeyCode: 52, Key: "$", Code: "Digit4"},
	'%':                {KeyCode: 53, Key: "%", Code: "Digit5"},
	'^':                {KeyCode: 54, Key: "^", Code: "Digit6"},
	'&':                {KeyCode: 55, Key: "&", Code: "Digit7"},
	'(':                {KeyCode: 57, Key: "(", Code: "Digit9"},
	'A':                {KeyCode: 65, Key: "A", Code: "KeyA"},
	'B':                {KeyCode: 66, Key: "B", Code: "KeyB"},
	'C':                {KeyCode: 67, Key: "C", Code: "KeyC"},
	'D':                {KeyCode: 68, Key: "D", Code: "KeyD"},
	'E':                {KeyCode: 69, Key: "E", Code: "KeyE"},
	'F':                {KeyCode: 70, Key: "F", Code: "KeyF"},
	'G':                {KeyCode: 71, Key: "G", Code: "KeyG"},
	'H':                {KeyCode: 72, Key: "H", Code: "KeyH"},
	'I':                {KeyCode: 73, Key: "I", Code: "KeyI"},
	'J':                {KeyCode: 74, Key: "J", Code: "KeyJ"},
	'K':                {KeyCode: 75, Key: "K", Code: "KeyK"},
	'L':                {KeyCode: 76, Key: "L", Code: "KeyL"},
	'M':                {KeyCode: 77, Key: "M", Code: "KeyM"},
	'N':                {KeyCode: 78, Key: "N", Code: "KeyN"},
	'O':                {KeyCode: 79, Key: "O", Code: "KeyO"},
	'P':                {KeyCode: 80, Key: "P", Code: "KeyP"},
	'Q':                {KeyCode: 81, Key: "Q", Code: "KeyQ"},
	'R':                {KeyCode: 82, Key: "R", Code: "KeyR"},
	'S':                {KeyCode: 83, Key: "S", Code: "KeyS"},
	'T':                {KeyCode: 84, Key: "T", Code: "KeyT"},
	'U':                {KeyCode: 85, Key: "U", Code: "KeyU"},
	'V':                {KeyCode: 86, Key: "V", Code: "KeyV"},
	'W':                {KeyCode: 87, Key: "W", Code: "KeyW"},
	'X':                {KeyCode: 88, Key: "X", Code: "KeyX"},
	'Y':                {KeyCode: 89, Key: "Y", Code: "KeyY"},
	'Z':                {KeyCode: 90, Key: "Z", Code: "KeyZ"},
	':':                {KeyCode: 186, Key: ":", Code: "Semicolon"},
	'<':                {KeyCode: 188, Key: "<", Code: "Comma"},
	'_':                {KeyCode: 189, Key: "_", Code: "Minus"},
	'>':                {KeyCode: 190, Key: ">", Code: "Period"},
	'?':                {KeyCode: 191, Key: "?", Code: "Slash"},
	'~':                {KeyCode: 192, Key: "~", Code: "Backquote"},
	'{':                {KeyCode: 219, Key: "{", Code: "BracketLeft"},
	'|':                {KeyCode: 220, Key: "|", Code: "Backslash"},
	'}':                {KeyCode: 221, Key: "}", Code: "BracketRight"},
	'"':                {KeyCode: 222, Key: "\"", Code: "Quote"},
}
