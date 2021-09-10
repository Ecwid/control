package tethering

type BindArgs struct {
	Port int `json:"port"`
}

type UnbindArgs struct {
	Port int `json:"port"`
}
