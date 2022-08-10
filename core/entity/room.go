package entity

type RoomKey string

type RoomDescriptions struct {
	General  string `json:"general"`
	Locked   string `json:"locked"`
	Action   string `json:"action"`
	Movement string `json:"movement"`
}

type Room struct {
	Name         string           `json:"name"`
	Descriptions RoomDescriptions `json:"descriptions"`
	Actions      Actions          `json:"actions"`
	Movements    Movements        `json:"movements"`
	Status       Status           `json:"status"`
}
