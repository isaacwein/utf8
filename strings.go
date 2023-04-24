package tools

func FilterNonUTF8(input string) string {
	result := make([]rune, 0, len(input))
	for _, r := range input {
		if isDisplayableInTerminal(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

func isDisplayableInTerminal(r rune) bool {
	return (r >= 0x20 && r <= 0x7E) || // Basic Latin (printable)
		(r >= 0xA0 && r <= 0xFF) || // Latin-1 Supplement (printable)
		(r >= 0x100 && r <= 0x10FFFF) // All other valid Unicode code points
}
