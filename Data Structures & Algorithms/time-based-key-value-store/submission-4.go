type TimeMap struct {
	Store map[string][]Pair
}

type Pair struct {
	Value string
	TimeStamp int
}

func Constructor() TimeMap {
	return TimeMap{
		Store: map[string][]Pair{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	pair := Pair{
		Value: value,
		TimeStamp: timestamp,
	}

	this.Store[key] = append(this.Store[key],pair)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.Store[key]
	if len(arr) == 0 {
		return ""
	}

	l,r := 0, len(arr)-1
	result := ""

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid].TimeStamp < timestamp {
			result = arr[mid].Value
			l = mid + 1
		} else if arr[mid].TimeStamp > timestamp {
			r = mid - 1
		} else {
			return arr[mid].Value
		}
	}

	return result
}
