package slice

import "testing"

func TestProduct(t *testing.T) {
	attrs := [][]interface{}{
		{
			"Man's",
			"Women's",
		},
		{
			"Sweat shirt",
			"T shirt",
		},
		{
			"S code",
			"L code",
		},
		{
			"Blue",
			"Red",
		},
	}

	result := Product(attrs...)
	for _, v := range result {
		for _, v1 := range v {
			t.Logf("%v ", v1)
		}
	}
}
