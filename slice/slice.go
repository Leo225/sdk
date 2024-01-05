package slice

func Product(sets ...[]interface{}) [][]interface{} {
	f := func(i int) int {
		return len(sets[i])
	}

	var products [][]interface{}
	for index := make([]int, len(sets)); index[0] < f(0); nextIndex(index, f) {
		var r []interface{}
		for i, k := range index {
			r = append(r, sets[i][k])
		}
		products = append(products, r)
	}
	return products
}

func nextIndex(index []int, f func(i int) int) {
	for i := len(index) - 1; i >= 0; i-- {
		index[i]++
		if i == 0 || index[i] < f(i) {
			return
		}
		index[i] = 0
	}
}
