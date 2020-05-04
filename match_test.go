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
		{false, "src", "dist/index.js"},
		{false, "src", "README.md"},
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
