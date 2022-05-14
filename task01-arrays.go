package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32 = 0
	var count int = 0
	for _, v := range input {
		if v > 0 {
			count += 1
		}
		sum += v
	}
	return sum / float32(count)
}
