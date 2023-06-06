package challenge

func ArrayOfArrayProducts(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	res := make([]int, len(arr))
	for i, _ := range arr {
		res[i] = sumAll(arr, i)
	}

	return res
}

func sumAll(arr []int, idx int) int {
	sum := 1
	for i, v := range arr {
		if i != idx {
			sum *= v
		}
	}
	return sum
}
