package world

import (
	"encoding/json"
	"github.com/cymon1997/text-game/entity"
	"io/ioutil"
)

func Render(path string) (*entity.World, error) {
	plan, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var world entity.World
	err = json.Unmarshal(plan, &world)
	if err != nil {
		return nil, err
	}
	return &world, nil
}
