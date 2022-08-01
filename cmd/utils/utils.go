package utils

import "regexp"

func IsArgsValids(args []string) bool {
	if len(args) != 1 {
		return false
	}

	checkArgShape := regexp.MustCompile(`^[0-9]+\.[0-9]+(?:\.[0-9]+)?$`)
	return checkArgShape.MatchString(args[0])
}
