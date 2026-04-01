	func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
		nums := []int{}

		for _,n1:= range nums1{
			nums = append(nums,n1)
		}

		for _,n2:= range nums2{
			nums = append(nums,n2)
		}

		sort.Ints(nums)

	if len(nums) % 2 != 0 {
		num := nums[(len(nums)-1)/2]
		return float64(num)
	} else {
		num := float64((len(nums)-1)/2)
		mid1 := math.Ceil(num)
		mid2 := mid1+1
		fmt.Println(mid1)
		fmt.Println(mid2)
		return float64(nums[int(mid1)]+nums[int(mid2)])/float64(2)
	}
	}
