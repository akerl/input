package list

import (
	"sort"
	"strings"
)

// Option defines a single option for a list prompt
type Option struct {
	Name     string
	Metadata map[string]string
}

// Types provides a map of prompt types by name
var Types = map[string]func() Prompt{
	"":      func() Prompt { return WmenuPrompt{} },
	"wmenu": func() Prompt { return WmenuPrompt{} },
	"fuzzy": func() Prompt { return FuzzyPrompt{} },
}

func Default() Prompt {
	return Types[""]()
}

// OptionSet defines a list of options for a list prompt
type OptionSet []Option

// Prompt defines an object which can execute list prompts
type Prompt interface {
	Execute(string, OptionSet) (int, error)
}

// String condenses the Name and Metadata into a single flat string
func (o Option) String() string {
	var str strings.Builder
	str.WriteString(o.Name)

	keys := make([]string, len(o.Metadata))
	index := 0
	for k := range o.Metadata {
		keys[index] = k
	}
	sort.Strings(keys)

	for _, k := range keys {
		str.WriteString(" ")
		str.WriteString(k)
		str.WriteString(":")
		str.WriteString(o.Metadata[k])
	}
	return str.String()
}
