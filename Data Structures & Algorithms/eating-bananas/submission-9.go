func minEatingSpeed(piles []int, h int) int {
	l,r := 1, 0
	hm := 0

	for _, p := range piles {
		r = max(r,p)
	}

	for l <= r {
		mid := l + (r-l)/2

		if eatMaxBanana(piles, mid, h) {
			hm = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return hm
}

func eatMaxBanana(piles []int, speed int, h int) bool {
	hUse := 0

	for i:= 0 ; i< len(piles);i++ {
		hUse += (piles[i] + speed - 1) / (speed)

		if hUse > h {
			return false
		}
	}

	return hUse <= h
}
