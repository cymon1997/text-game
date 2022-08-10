package provider

import (
	"fmt"
	"sync"

	"github.com/cymon1997/text-game/core/entity"
	engine "github.com/cymon1997/text-game/engine/world"
)

var (
	world     *entity.World
	syncWorld sync.Once
)

func GetWorld(scene int) *entity.World {
	syncWorld.Do(func() {
		world, _ = engine.Render(fmt.Sprintf("%s/scene_%d.json", entity.PathMap, scene))
	})
	return world
}
