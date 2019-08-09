package list

import (
	"github.com/ktr0731/go-fuzzyfinder"
)

// FuzzyPrompt implements a fuzzy interactive picker
type FuzzyPrompt struct{}

// TODO: Support displaying message

// Execute runs the list prompt
func (f FuzzyPrompt) Execute(_ string, os OptionSet) (int, error) {
	logger.InfoMsg("executing fuzzy prompt")
	lines := make([]bool, len(os))
	return fuzzyfinder.Find(lines, func(i int) string {
		return os[i].String()
	})
}
