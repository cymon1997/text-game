package game

import (
	"github.com/cymon1997/text-game/engine"
	"github.com/cymon1997/text-game/engine/logger"
	"github.com/cymon1997/text-game/entity"
)

var (
	modeChange = false
)

func Play(world *entity.World) {
	logger.Printf("You are in %s\n", world.Rooms[world.Position].Name)
	logger.Println(world.Rooms[world.Position].Desc)
	isNeedAction := WaitForInput(world)
	for isNeedAction {
		//logger.Printf("You are in %s\n", world.Rooms[world.Position].Name)
		//logger.Println(world.Rooms[world.Position].Desc)
		isNeedAction = WaitForInput(world)
	}
}

func WaitForInput(world *entity.World) bool {
	var options []entity.Option
	switch world.Mode {
	case entity.ModeExamine:
		if modeChange {
			logger.Println(world.Rooms[world.Position].ActionDesc)
			modeChange = false
		}
		logger.Println("What should I do?")
		options = GenerateExamineOptions(world, world.Position)
	case entity.ModeMove:
		if modeChange {
			logger.Println(world.Rooms[world.Position].MoveDesc)
			modeChange = false
		}
		logger.Println("Where should I go?")
		options = GenerateMoveOptions(world, world.Position)
	default:
		options = GenerateModeOptions(world)
	}
	engine.PromptOptions(options)
	input, err := engine.InputInt("Your action: ")
	engine.Clear()
	for err != nil || input < 1 || input > len(options) {
		logger.Println("I don't understand what you mean...")
		return true
	}
	return Execute(world, options[input-1].Value)
}

func Execute(world *entity.World, key string) bool {
	if key == "" {
		world.Mode = ""
	}
	switch world.Mode {
	case entity.ModeExamine:
		Examine(world, key)
	case entity.ModeMove:
		MoveTo(world, key)
	default:
		world.Mode = key
		modeChange = true
	}
	return true
}
