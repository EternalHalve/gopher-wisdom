package quotes

import (
	"fmt"
	"strings"
)

func Alienify(quote string) string {
	replacements := map[string]string{
		"the":   "ze",
		"human": "carbon-unit",
		"love":  "bio-resonance",
		"world": "sector-G",
		"life":  "process-v0.1",
	}

	words := strings.Fields(quote)
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		if val, ok := replacements[lowerWord]; ok {
			words[i] = val
		}
	}

	translated := strings.Join(words, " ")

	r := strings.NewReplacer(
		"s", "z", "S", "Z",
		"a", "v", "A", "V",
		"o", "0", "O", "0",
		"e", "ë", "E", "Ë",
	)

	result := r.Replace(translated)

	return fmt.Sprintf("⋏ [%s] ⋎", result)
}
