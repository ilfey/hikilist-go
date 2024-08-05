package inflection

import "bytes"

func Snake(camel string) string {
	var buf bytes.Buffer

	for _, char := range camel {
		if char <= 'Z' && char >= 'A' {
			if buf.Len() > 0 {
				buf.WriteRune('_')
			}

			buf.WriteRune(char + 32) // 32 is the shift to lowercase
		} else {
			buf.WriteRune(char)
		}
	}

	return buf.String()
}
