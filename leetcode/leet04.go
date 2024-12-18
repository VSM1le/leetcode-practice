package leetcode

func SortArray(list []int) []int {
	var numIndex int
	var smallestValue int
	var sortedValue []int
	for len(list) > 0 {
		numIndex = 0
		smallestValue = list[0]
		for index, value := range list {
			if value < smallestValue {
				smallestValue = value
				numIndex = index
			}
		}
		list = append(list[0:numIndex], list[numIndex+1:]...)
		sortedValue = append(sortedValue, smallestValue)
	}
	return sortedValue
}

func FindMedianSortedArrays() float64 {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	mergeNums := append(nums1, nums2...)
	sortedNums := SortArray(mergeNums)
	lenNums := len(sortedNums) / 2
	if len(sortedNums)%2 != 0 {
		return float64(sortedNums[lenNums])
	}
	sumValue := sortedNums[lenNums] + sortedNums[lenNums-1]
	return float64(sumValue) / 2
}
