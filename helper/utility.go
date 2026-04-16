package helper

import "strings"

func GenerateSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}