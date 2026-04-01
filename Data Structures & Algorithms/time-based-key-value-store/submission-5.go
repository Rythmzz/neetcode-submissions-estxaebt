type TimeMap struct {
	store map[string][]Pair
}

type Pair struct {
	val string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap {
		store: map[string][]Pair{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key],Pair{val: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.store[key]
	result := ""

	if len(arr) == 0 {
		return ""
	}

	l,r := 0, len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid].timestamp == timestamp {
			return arr[mid].val
		}

		if arr[mid].timestamp < timestamp {
			result = arr[mid].val
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return result
}
