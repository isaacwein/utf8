package tools

import (
	"testing"
)

func TestFilterNonUTF8(t *testing.T) {
	t.Log("TestFilterNonUTF8")
	input := "Hello, 世界! Gophér\x80" // The \x80 is an invalid UTF8 character
	t.Log("Input string:", input)

	filtered := FilterNonUTF8(input)
	t.Log("Filtered string:", filtered)
}
