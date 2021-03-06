package list

import (
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/nsf/termbox-go"
)

// FuzzyPrompt implements a fuzzy interactive picker
type FuzzyPrompt struct{}

// Execute runs the list prompt
func (f FuzzyPrompt) Execute(msg string, os OptionSet) (int, error) {
	logger.InfoMsg("executing fuzzy prompt")
	lines := make([]bool, len(os))

	handler := func(i int) string {
		return os[i].String()
	}

	return fuzzyfinder.Find(
		lines,
		handler,
		fuzzyfinder.WithPromptString(msg+": "),
		fuzzyfinder.WithPromptColor(termbox.ColorDefault),
	)
}
