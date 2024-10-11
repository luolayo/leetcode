package code

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for i := 1; i < len(nums); i++ {
		if nums[sum] != nums[i] {
			sum++
			nums[sum] = nums[i]
		}
	}
	return sum + 1
}
