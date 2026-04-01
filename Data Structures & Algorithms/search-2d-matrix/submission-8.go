func searchMatrix(matrix [][]int, target int) bool {
	for i:= 0;i < len(matrix);i++ {
		if searchBinary(matrix[i],target) {
			return true
		}
	}

	return false
}

func searchBinary(arr []int, target int) bool{
	l, r := 0, len(arr)-1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == target {
			return true
		}

		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}


 	}

	return false
}
