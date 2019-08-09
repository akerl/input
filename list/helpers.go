package list

import (
	"fmt"
	"sort"
)

// WithInputString returns either the pre-provided selection or the result of a user prompt
func WithInputString(p Prompt, options []string, val, msg string) (string, error) {
	logger.InfoMsg("running WithInputString helper")
	if val != "" {
		logger.DebugMsgf("value provided: %s", val)
		for _, item := range options {
			if item == val {
				return val, nil
			}
		}
		return "", fmt.Errorf("user provided selection not found: %s", val)
	}
	if len(options) == 1 {
		logger.DebugMsgf("only one option provided: %s", options[0])
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
	logger.InfoMsg("running WithInputSlice helper")
	switch len(list) {
	case 0:
		logger.DebugMsg("empty input list provided")
		return WithInputString(p, options, "", msg)
	case 1:
		logger.DebugMsg("single-item input list provided")
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
