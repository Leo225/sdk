package bytes

import "testing"

func TestByteSize_String(t *testing.T) {
	var b ByteSize
	b = 1 << (10 * 5)
	result := b.String()
	t.Logf("ByteSize byte: %f, result: %s", b, result)
}
