package provider

import (
	"bufio"
	"os"
	"sync"
)

var (
	inputReader *bufio.Reader
	syncInputReader sync.Once
)

func GetInputReader() *bufio.Reader {
	syncInputReader.Do(func() {
		inputReader = bufio.NewReader(os.Stdin)
	})
	return inputReader
}