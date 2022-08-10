package game

import (
	"fmt"

	"github.com/cymon1997/text-game/core/entity"
	"github.com/cymon1997/text-game/engine/logger"
)

func GenerateModeOptions(_ *entity.World) []entity.Option {
	return []entity.Option{
		{
			Name:  "Look around",
			Value: ModeAction,
		},
		{
			Name:  "Move",
			Value: ModeMovement,
		},
	}
}

func GenerateActionOptions(world *entity.World, loc entity.RoomKey) []entity.Option {
	var options []entity.Option
	pos := world.Rooms[loc]
	for _, key := range pos.Actions {
		if action, ok := world.Actions[key]; ok && !action.Status.IsHidden {
			options = append(options, entity.Option{
				Name:  action.Name,
				Value: string(key),
			})
		}
	}
	options = append(options, entity.Option{
		Name: "[Back]",
	})
	return options
}

func GenerateMovementOptions(world *entity.World, loc entity.RoomKey) []entity.Option {
	var options []entity.Option
	pos := world.Rooms[loc]
	for _, key := range pos.Movements {
		if room, ok := world.Rooms[key]; ok {
			options = append(options, entity.Option{
				// TODO: Move this desc to room's movements
				Name:  fmt.Sprint("Go to ", room.Name),
				Value: string(key),
			})
		}
	}
	options = append(options, entity.Option{
		Name: "[Back]",
	})
	return options
}

func DoAction(world *entity.World, key entity.ActionKey) {
	if action, ok := world.Actions[key]; ok {
		//logger.Printf("You %s\n", action.Name)
		if !action.Status.IsLocked {
			logger.Println(action.Descriptions.General)
			applyAction(world, key)
		} else {
			logger.Println(action.Descriptions.Locked)
		}
	}
}

func applyAction(world *entity.World, key entity.ActionKey) {
	if action, ok := world.Actions[key]; ok {
		for _, update := range action.Updates.Room {
			if roomUpdate, roomOK := world.Rooms[entity.RoomKey(update.Key)]; roomOK {
				roomUpdate.Status.IsLocked = update.IsLocked
				roomUpdate.Status.IsHidden = update.IsHidden
			}
		}
		for _, update := range action.Updates.Action {
			if actionUpdate, actionOK := world.Actions[entity.ActionKey(update.Key)]; actionOK {
				actionUpdate.Status.IsLocked = update.IsLocked
				actionUpdate.Status.IsHidden = update.IsHidden
			}
		}
	}
}

func DoMovement(world *entity.World, target entity.RoomKey) {
	if room, ok := world.Rooms[target]; ok {
		if !room.Status.IsLocked {
			//logger.Printf("You move to %s\n", room.Name)
			world.CurrentPosition = target
		} else {
			logger.Println(room.Descriptions.Locked)
		}
	}
}
