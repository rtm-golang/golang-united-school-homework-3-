package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var keys []int
	for k, _ := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for k := range keys {
		result = append(result, input[k])
	}
	return
}
