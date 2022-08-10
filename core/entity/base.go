package entity

type Status struct {
	IsLocked bool `json:"is_locked"`
	IsHidden bool `json:"is_hidden"`
}

type Update struct {
	Key      string `json:"key"`
	IsLocked bool   `json:"is_locked"`
	IsHidden bool   `json:"is_hidden"`
}
