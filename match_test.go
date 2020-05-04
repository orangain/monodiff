package monodiff

import "testing"

type singleInputTest struct {
	expected  bool
	pattern   string
	inputPath string
}

func TestEmptyPatternMatch(t *testing.T) {
	result := matchPattern("src", []string{})
	if result != false {
		t.Fatal("should not match empty paths")
	}
}

func TestSinglePatternMatch(t *testing.T) {
	tests := []singleInputTest{
		{true, "src", "src/index.js"},
		{true, "src", "src"},
		{false, "src", "sr"},
		{false, "src", "src.txt"},
		{false, "src", "src2/index.js"},
		{false, "src", "old_src/index.js"},
		{false, "src", "dist/index.js"},
		{false, "src", "README.md"},
		{true, "package-lock.json", "package-lock.json"},
		{true, "package-lock.json", "package-lock.json/foo"},
		{false, "package-lock.json", "package-lock.json5"},
		{true, "package*.json", "package-lock.json"},
		{true, "package*.json", "package.json"},
		{false, "package*.json", "package/config.json"},
		{true, "**/build.gradle.kts", "libs/build.gradle.kts"},
		{true, "**/build.gradle.kts", "libs/profile/build.gradle.kts"},
		{false, "**/build.gradle.kts", "build.gradle.kts"},          // Pattern "**/" requires "/"
		{true, "libs/**/build.gradle.kts", "libs/build.gradle.kts"}, // Pattern "/**/" does not require "//"
		{true, "libs/**/build.gradle.kts", "libs/profile/build.gradle.kts"},
		{true, "libs/**/build.gradle.kts", "libs/profile/util/build.gradle.kts"},
		{true, "*", "foo.txt"},
		{true, "*", "test/foo.txt"}, // Pattern "*" matches any directory and files under the directory. So it is equivalent to "**"
		{true, "*.bak", "foo.bak"},
		{true, "*.bak", "foo.bak/bar.txt"},
		{false, "*.bak", "foo/bar.bak"},
		{true, "index*", "index"},
		{true, "index*", "index.js"},
		{true, "index*", "index.d/foo"},
		{true, "**", "foo.txt"},
		{true, "**", "test/foo.txt"},
		{true, "**.bak", "foo.bak"},
		{true, "**.bak", "foo.bak/bar.txt"},
		{true, "**.bak", "foo/bar.bak"},
		{true, "index**", "index"},
		{true, "index**", "index.js"},
		{true, "index**", "index.d/foo"},
	}

	for _, test := range tests {
		result := matchPattern(test.pattern, []string{test.inputPath})
		if result != test.expected {
			if test.expected {
				t.Fatalf("expected pattern %#v should match input %#v but not matched", test.pattern, test.inputPath)
			} else {
				t.Fatalf("expected pattern %#v should not match input %#v but matched", test.pattern, test.inputPath)
			}
		}
	}
}
