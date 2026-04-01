func searchMatrix(matrix [][]int, target int) bool {
	for i:= 0 ;i< len(matrix);i++ {
		arr := []int{}
		for j:= 0 ;j< len(matrix[i]);j++ {
			arr = append(arr,matrix[i][j])
		}
		if searchBinary(arr,target) {
			return true
		}
	}

	return false
	
	// return searchBinary(arr,target)

}

func searchBinary(arr []int, target int) bool {
	mid := len(arr)/2
	start := 0
	end := len(arr)
	fmt.Println(arr)
	fmt.Println(arr[mid])
	if target > arr[mid] {
		start = mid
	} else if target < arr[mid] {
		end = mid
	} else {
		return true
	}

	for start < end {
		if arr[start] == target {
			return true
		}
		start++
	}

	return false
}
