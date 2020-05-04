package monodiff

import "strings"

func matchPattern(globPattern string, paths []string) bool {
	for _, path := range paths {
		if path == globPattern || strings.HasPrefix(path, globPattern+"/") {
			return true
		}
	}
	return false
}
