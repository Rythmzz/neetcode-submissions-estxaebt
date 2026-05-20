type TimeMap struct {
	data map[string][]Pair
}

type Pair struct {
	val string
	ts int
}

func Constructor() TimeMap {
	return TimeMap{data: map[string][]Pair{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key],Pair{val:value, ts: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	store := this.data[key]
	value := ""
	l,r := 0, len(store)-1
	for l <= r {
		mid := l + (r-l)/2

		if store[mid].ts == timestamp {
			return store[mid].val
		}

		if store[mid].ts < timestamp {
			value = store[mid].val
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return value

}
