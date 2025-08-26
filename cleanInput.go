package main

import (
	"strings"
)

func CleanInput(text string) []string {
	if len(text) == 0 {
        return []string{""}
    }
	lowerCased := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowerCased)
	return strings.Fields(trimmed)
}