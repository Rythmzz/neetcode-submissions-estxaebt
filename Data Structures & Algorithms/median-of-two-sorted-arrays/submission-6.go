func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := []int{}

	for _, n := range nums1 {
		nums = append(nums,n)
	}

	for _, n := range nums2 {
		nums = append(nums,n)
	}

	sort.Ints(nums)

	if len(nums) % 2 != 0 {
		return float64(nums[(len(nums)-1)/2])
	} else {
		x1 := float64(nums[(len(nums)-1)/2])
		x2 := float64(nums[((len(nums)-1)/2)+1])

		return (x1+x2)/2
	}
}
