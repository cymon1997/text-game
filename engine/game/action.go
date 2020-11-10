package game

import (
	"fmt"
	"github.com/cymon1997/text-game/engine/logger"
	"github.com/cymon1997/text-game/entity"
)

func GenerateModeOptions(world *entity.World) []entity.Option {
	return []entity.Option{
		{
			Name:  "Look around",
			Value: entity.ModeExamine,
		},
		{
			Name:  "Move",
			Value: entity.ModeMove,
		},
	}
}

func GenerateExamineOptions(world *entity.World, loc string) []entity.Option {
	options := make([]entity.Option, 0)
	here := world.Rooms[loc]
	for _, actionKey := range here.AvailableTasks {
		if action := world.Tasks[actionKey]; action.IsShown {
			options = append(options, entity.Option{
				Name:  action.Name,
				Value: actionKey,
			})
		}
	}
	options = append(options, entity.Option{
		Name: "[Task list]",
	})
	return options
}

func GenerateMoveOptions(world *entity.World, loc string) []entity.Option {
	options := make([]entity.Option, 0)
	here := world.Rooms[loc]
	if here.In != "" {
		options = append(options, entity.Option{
			Name:  fmt.Sprintf("Go inside [%s]", world.Rooms[here.In].Name),
			Value: here.In,
		})
	}
	if here.Left != "" {
		options = append(options, entity.Option{
			Name:  fmt.Sprintf("Go to left [%s]", world.Rooms[here.Left].Name),
			Value: here.Left,
		})
	}
	if here.Right != "" {
		options = append(options, entity.Option{
			Name:  fmt.Sprintf("Go to right [%s]", world.Rooms[here.Right].Name),
			Value: here.Right,
		})
	}
	if here.Out != "" {
		options = append(options, entity.Option{
			Name:  fmt.Sprintf("Go outside [%s]", world.Rooms[here.Out].Name),
			Value: here.Out,
		})
	}
	options = append(options, entity.Option{
		Name: "[Task list]",
	})
	return options
}

func Examine(world *entity.World, actionKey string) {
	action := world.Tasks[actionKey]
	logger.Printf("You %s\n", action.Name)
	if action.IsAble {
		logger.Println(action.Desc)
		applyAction(world, actionKey)
	} else {
		logger.Println(action.DescLocked)
	}
}

func MoveTo(world *entity.World, loc string) {
	if world.Rooms[loc].IsAble {
		logger.Printf("You move to %s\n", world.Rooms[loc].Name)
		world.Position = loc
	} else {
		logger.Println(world.Rooms[loc].DescLocked)
	}
}

func applyAction(world *entity.World, actionKey string) {
	action := world.Tasks[actionKey]
	// Apply task actions
	for _, nextKey := range action.TaskAction.Enable {
		world.Tasks[nextKey].IsAble = true
	}
	for _, nextKey := range action.TaskAction.Disable {
		world.Tasks[nextKey].IsAble = false
	}
	for _, nextKey := range action.TaskAction.Show {
		world.Tasks[nextKey].IsShown = true
	}
	for _, nextKey := range action.TaskAction.Hide {
		world.Tasks[nextKey].IsShown = false
	}
	// Apply room actions
	for _, nextKey := range action.RoomAction.Enable {
		world.Rooms[nextKey].IsAble = true
	}
	for _, nextKey := range action.RoomAction.Disable {
		world.Rooms[nextKey].IsAble = false
	}
	for _, nextKey := range action.RoomAction.Show {
		world.Rooms[nextKey].IsShown = true
	}
	for _, nextKey := range action.RoomAction.Hide {
		world.Rooms[nextKey].IsShown = false
	}
}
