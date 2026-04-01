type TimeMap struct {
	StoreVal map[string][]string
	StoreTimeStamp map[string][]int
}

func Constructor() TimeMap {
	return TimeMap{
		StoreVal: map[string][]string{},
		StoreTimeStamp: map[string][]int{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if len(this.StoreVal[key]) != 0 {

		arr := this.StoreVal[key]
		for i:= 0;i< len(arr);i++ {
			if arr[i] == value {
				return
			}
		}

		this.StoreVal[key] = append(this.StoreVal[key],value)
	} else {
		this.StoreVal[key] = []string{value}
	}

	if len(this.StoreTimeStamp[key]) != 0 {
		this.StoreTimeStamp[key] = append(this.StoreTimeStamp[key],timestamp)
	} else {
		this.StoreTimeStamp[key] = []int{timestamp}
	}

}

func (this *TimeMap) Get(key string, timestamp int) string {
	arrVal := this.StoreVal[key]
	arrTime := this.StoreTimeStamp[key]
	if len(arrVal) != 0 {
		l,r := 0, len(arrTime)-1
		minVal := ""
		for l <= r {
			mid := l + (r-l)/2

			if arrTime[mid] < timestamp {
				minVal = arrVal[mid]
				l = mid + 1
			} else if arrTime[mid] > timestamp {
				r = mid - 1
			} else {
				return arrVal[mid]
			}
		}

		fmt.Printf("l: %d, r: %d",l,r)

		return minVal

		// for i:= 0;i< len(arrTime);i++{
		// 	if arrTime[i] == timestamp {
		// 		return arrVal[i]
		// 	}

		// 	if arrTime[i] > indexMaxTime && indexMaxTime < timestamp {
		// 		indexMaxTime = i
		// 	}
		// }
		
		
	}

	return ""
}
