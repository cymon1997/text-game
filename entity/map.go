package entity

const (
	PathMap = "src/map"
)

type Direction struct {
	Left  string `json:"left"`
	Right string `json:"right"`
	In    string `json:"in"`
	Out   string `json:"out"`
}

type Description struct {
	Desc       string `json:"desc"`
	DescLocked string `json:"desc_locked"`
	ActionDesc string `json:"action_desc"`
	MoveDesc   string `json:"move_desc"`
}

type Action struct {
	Name           string   `json:"name"`
	IsAble         bool     `json:"is_able"`
	IsShown        bool     `json:"is_shown"`
	EnableActions  []string `json:"enable_actions"`
	DisableActions []string `json:"disable_actions"`
	ShowActions    []string `json:"show_actions"`
	HideActions    []string `json:"hide_actions"`
	Description
}

type Room struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
	IsAble  bool     `json:"is_able"`
	IsShown bool     `json:"is_shown"`
	Description
	Direction
}

type World struct {
	Rooms      map[string]*Room   `json:"rooms"`
	Actions    map[string]*Action `json:"actions"`
	Position   string             `json:"position"`
	ActionMode string
}
