func searchMatrix(matrix [][]int, target int) bool {
	for i:= 0 ;i< len(matrix);i++ {
		if search(matrix[i],target) {
			return true
		}
	}

	return false
}

func search(list []int, target int) bool {
	l,r := 0, len(list)-1

	for l <= r {
		mid := l + (r-l)/2

		if list[mid] == target {
			return true
		}

		if list[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
