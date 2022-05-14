package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	reversed := make([]int64, len(input))
	for i, v := range input {
		reversed[len(reversed)-i-1] = v
	}
	return reversed
}
