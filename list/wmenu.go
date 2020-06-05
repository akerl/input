package list

import (
	"fmt"
	"os"

	"github.com/dixonwille/wmenu/v5"
)

// WmenuPrompt implements a wmenu-based prompt
type WmenuPrompt struct{}

// Execute runs the list prompt
func (w WmenuPrompt) Execute(msg string, optSet OptionSet) (int, error) {
	logger.InfoMsg("executing wmenu prompt")
	c := make(chan int, 1)

	menu := wmenu.NewMenu(msg + ":")
	menu.ChangeReaderWriter(os.Stdin, os.Stderr, os.Stderr)
	menu.LoopOnInvalid()
	menu.InitialIndex(0)

	actionFunc := func(opts []wmenu.Opt) error {
		if opts[0].ID == -1 {
			return fmt.Errorf("no response provided")
		}
		c <- opts[0].ID
		return nil
	}
	menu.Action(actionFunc)

	for _, option := range optSet {
		menu.Option(option.String(), nil, false, nil)
	}

	if err := menu.Run(); err != nil {
		return -1, err
	}

	return <-c, nil
}
