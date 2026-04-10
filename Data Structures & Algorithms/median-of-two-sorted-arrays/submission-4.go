func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := []int{}

	for _, n:= range nums1 {
		nums= append(nums,n)
	}

	for _, n:= range nums2 {
		nums= append(nums,n)
	}

	sort.Ints(nums)

	n := nums[(len(nums)-1)/2]
	if len(nums) % 2 != 0 {
		return float64(n)
	} else {
		m := nums[((len(nums)-1)/2)+1]
		return float64(n+m)/2
	}
}
