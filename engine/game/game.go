package game

import (
	"github.com/cymon1997/text-game/core/entity"
	"github.com/cymon1997/text-game/engine"
	"github.com/cymon1997/text-game/engine/logger"
)

const (
	ModeAction   string = "MODE_ACTION"
	ModeMovement string = "MODE_MOVEMENT"
)

var (
	isNeedLocationInfo = true
)

func Play(world *entity.World) {
	engine.Clear()
	if _, ok := world.Rooms[world.CurrentPosition]; ok {
		//logger.Printf("You are in %s\n", room.Name)
		//logger.Println(room.Descriptions.General)
		isNeedAction := WaitForInput(world)
		for isNeedAction {
			//logger.Printf("You are in %s\n", room.Name)
			//logger.Println(room.Descriptions.General)
			isNeedAction = WaitForInput(world)
		}
	} else {
		logger.Println("Error: your saved file is corrupted.")
	}
}

func WaitForInput(world *entity.World) bool {
	if currentRoom, ok := world.Rooms[world.CurrentPosition]; ok {
		if isNeedLocationInfo {
			logger.Printf("You are in %s.\n", currentRoom.Name)
			logger.Println(currentRoom.Descriptions.General)
			isNeedLocationInfo = false
		}
		var options []entity.Option
		switch world.CurrentMode {
		case ModeAction:
			logger.Println(currentRoom.Descriptions.Action)
			logger.Println("What should I do?")
			options = GenerateActionOptions(world, world.CurrentPosition)
		case ModeMovement:
			logger.Println(currentRoom.Descriptions.Movement)
			logger.Println("Where should I go?")
			options = GenerateMovementOptions(world, world.CurrentPosition)
		default:
			logger.Println("What should I do?")
			options = GenerateModeOptions(world)
		}
		engine.PromptOptions(options)
		// TODO: Implement more flexible inputs
		input, err := engine.InputInt("Your action: ")
		engine.Clear()
		for err != nil || input < 1 || input > len(options) {
			logger.Println("I don't understand what you mean...")
			isNeedLocationInfo = true
			return true
		}
		return Execute(world, options[input-1].Value)
	}
	logger.Println("Error: unknown error")
	return false
}

func Execute(world *entity.World, key string) bool {
	if key == "" {
		world.CurrentMode = ""
		isNeedLocationInfo = true
	}
	switch world.CurrentMode {
	case ModeAction:
		DoAction(world, entity.ActionKey(key))
	case ModeMovement:
		DoMovement(world, entity.RoomKey(key))
		isNeedLocationInfo = true
		world.CurrentMode = ""
	default:
		world.CurrentMode = key
	}
	return true
}
