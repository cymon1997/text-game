package main

import (
	"github.com/cymon1997/text-game/engine/game"
	"github.com/cymon1997/text-game/provider"
)

func main() {
	world := provider.GetWorld(1)
	game.Play(world)
}
