func searchMatrix(matrix [][]int, target int) bool {
	for i:= 0 ;i< len(matrix);i++ {
		if search(matrix[i],target) {
			return true
		}
	}

	return false
}

func search(nums []int, target int) bool {
	l,r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return true
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

