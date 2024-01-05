package strings

import "bytes"

func PadLeft(left, s string, length int) string {
	buf := bytes.NewBuffer(nil)
	for i := 0; i < length; i++ {
		buf.WriteString(left)
	}
	buf.WriteString(s)
	return buf.String()
}

func PadRight(right, s string, length int) string {
	buf := bytes.NewBufferString(s)
	for i := 0; i < length; i++ {
		buf.WriteString(right)
	}
	return buf.String()
}
