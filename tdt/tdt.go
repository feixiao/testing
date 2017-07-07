package tdt

func Add(a, b int) int {
	return a + b
}

func Min(nums []int) int {
	var rst int
	if nums != nil {
		rst = nums[0]
	} else {
		return 0
	}
	for i := range nums {
		if rst > nums[i] {
			rst = nums[i]
		}
	}
	return rst
}

func Max(nums []int) int {
	var rst int
	if nums != nil {
		rst = nums[0]
	} else {
		return 0
	}
	for i := range nums {
		if rst < nums[i] {
			rst = nums[i]
		}
	}
	return rst
}
