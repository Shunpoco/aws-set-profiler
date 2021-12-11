package parser

import "testing"

func TestParseConfig(t *testing.T) {
	items, err := ParseConfig("./test.txt")
	if err != nil {
		t.Errorf("ParserConfig failed: %v\n", err)
	}

	expected := []string{
		"profile1",
		"profile3",
	}

	if len(items) != len(expected) {
		t.Errorf("Expected length of items is %d, but found %d.\n", len(expected), len(items))
	}

	for i := 0; i < len(expected); i++ {
		if items[i] != expected[i] {
			t.Errorf("Expected value of index %d is %s, but found %s.\n", i, expected[i], items[i])
		}
	}
}
