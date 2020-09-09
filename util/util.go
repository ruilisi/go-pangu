package util

import "strings"

func Contains(s []string, target string) bool {
	for _, element := range s {
		if strings.EqualFold(element, target) {
			return true
		}
	}
	return false
}
