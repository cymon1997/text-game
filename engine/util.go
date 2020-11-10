package engine

import (
	"fmt"
	"github.com/cymon1997/text-game/engine/logger"
	"github.com/cymon1997/text-game/engine/provider"
	"github.com/cymon1997/text-game/entity"
	"strconv"
	"strings"
)

func PromptOptions(opt []entity.Option) {
	for i, v := range opt {
		logger.Printf("%d) %s\n", i+1, v.Name)
	}
}

func InputInt(prompt string) (input int, err error) {
	logger.Print(prompt)
	raw, _ := provider.GetInputReader().ReadString('\n')
	return strconv.Atoi(strings.ReplaceAll(raw, "\n", ""))
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}