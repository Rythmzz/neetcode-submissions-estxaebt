func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := []int{}

	for _, n:= range nums1 {
		nums = append(nums,n)
	}

	for _, n:= range nums2 {
		nums = append(nums,n)
	}

	sort.Ints(nums)
	n := (len(nums)-1)/2

	if len(nums) % 2 != 0 {
		return float64(nums[n])
	} else {
		m:= n+1
		return (float64(nums[n])+float64(nums[m]))/float64(2)
	}


}
