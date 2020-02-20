package list

import (
	"os"

	"github.com/dixonwille/wmenu"
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

	actionFunc := func(opts []wmenu.Opt) error {
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
