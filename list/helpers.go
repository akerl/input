package list

import (
	"fmt"
	"sort"
)

// WithInputString returns either the pre-provided selection or the result of a user prompt
func WithInputString(p Prompt, options []string, val, msg string) (string, error) {
	if val != "" {
		for _, item := range options {
			if item == val {
				return val, nil
			}
		}
		return "", fmt.Errorf("user provided selection not found: %s", val)
	}
	if len(options) == 1 {
		return options[0], nil
	}

	sort.Strings(options)
	os := make(OptionSet, len(options))
	for index, option := range options {
		os[index] = Option{Name: option}
	}
	index, err := p.Execute(msg, os)
	if err != nil {
		return "", err
	}
	return options[index], nil
}

// WithInputSlice filters options based on a provided slice and then performs a
// WithInputString prompt
func WithInputSlice(p Prompt, options, list []string, msg string) (string, error) {
	switch len(list) {
	case 0:
		return WithInputString(p, options, "", msg)
	case 1:
		return WithInputString(p, options, list[0], msg)
	default:
		var matchList []string
		for _, item := range list {
			for _, option := range options {
				if option == item {
					matchList = append(matchList, item)
					break
				}
			}
		}
		if len(matchList) == 0 {
			return "", fmt.Errorf("no matching options found")
		}
		return WithInputString(p, matchList, "", msg)
	}
}
