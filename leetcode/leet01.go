package leetcode

func TwoSum() []int {
	target := 9
	nums := []int{2, 7, 11, 15}
	mapIndex := make(map[int]int)
	for i, v := range nums {
		mapIndex[v] = i
	}
	for i, v := range nums {
		anotherValue := target - v
		if _, ok := mapIndex[anotherValue]; ok && mapIndex[anotherValue] != i {
			return []int{i, mapIndex[anotherValue]}
		}
	}
	return nil
}
