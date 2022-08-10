package entity

type ActionKey string

type Actions []ActionKey

type ActionDescriptions struct {
	General string `json:"general"`
	Locked  string `json:"locked"`
}

type ActionUpdate struct {
	Action []Update `json:"actions"`
	Room   []Update `json:"rooms"`
}

type Action struct {
	Name         string             `json:"name"`
	Descriptions ActionDescriptions `json:"descriptions"`
	Updates      ActionUpdate       `json:"updates"`
	Status       Status             `json:"status"`
}
