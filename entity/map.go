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
	ActionDesc string `json:"desc_examine"`
	MoveDesc   string `json:"desc_move"`
}

type Action struct {
	Enable  []string `json:"enable"`
	Disable []string `json:"disable"`
	Show    []string `json:"show"`
	Hide    []string `json:"hide"`
}

type Task struct {
	Name       string `json:"name"`
	IsAble     bool   `json:"is_able"`
	IsShown    bool   `json:"is_shown"`
	TaskAction Action `json:"task_action"`
	RoomAction Action `json:"room_action"`
	Description
}

type Room struct {
	Name           string   `json:"name"`
	IsAble         bool     `json:"is_able"`
	IsShown        bool     `json:"is_shown"`
	AvailableTasks []string `json:"tasks"`
	Description
	Direction
}

type World struct {
	Rooms    map[string]*Room `json:"rooms"`
	Tasks    map[string]*Task `json:"tasks"`
	Position string           `json:"position"`
	Mode     string
}
