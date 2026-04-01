func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := []int{}
	for i:= 0 ;i< len(nums1);i++ {
		nums = append(nums,nums1[i])
	}

	for i:= 0 ;i< len(nums2);i++ {
		nums = append(nums,nums2[i])
	}

	sort.Ints(nums)
	n := nums[(len(nums)-1)/2]

	if len(nums) % 2 != 0 {
		return float64(n)
	} else {
		m := nums[((len(nums)-1)/2)+1]

		return float64(m+n)/2
	}
}
