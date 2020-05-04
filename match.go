package monodiff

import (
	"github.com/gobwas/glob"
)

func matchPattern(globPattern string, paths []string) bool {
	sep := '/'
	globFile := glob.MustCompile(globPattern, sep)      // Exact match
	globDir := glob.MustCompile(globPattern+"/**", sep) // Prefix match

	for _, path := range paths {
		if globFile.Match(path) || globDir.Match(path) {
			return true
		}
	}
	return false
}
