package entity

const (
	PathMap = "src/map"
)

type World struct {
	Rooms           map[RoomKey]*Room     `json:"rooms"`
	Actions         map[ActionKey]*Action `json:"actions"`
	CurrentPosition RoomKey               `json:"current_position"`
	CurrentMode     string
}
