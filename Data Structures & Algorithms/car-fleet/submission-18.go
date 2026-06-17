func carFleet(target int, position []int, speed []int) int {
	cars := make([][2]int,len(position))

	for i:= 0; i < len(cars);i++ {
		cars[i][0] = position[i]
		cars[i][1] = speed[i]
	}

	sort.Slice(cars,func (i,j int) bool {
		return cars[i][0] > cars[j][0]
	})
	fleet := 1
	maxD := (float64(target) - float64(cars[0][0]))/ float64(cars[0][1])
	for i:= 1; i< len(cars) ;i++ {
		d := (float64(target) - float64(cars[i][0]))/ float64(cars[i][1])
		if maxD < d {
			maxD = d
			fleet++
		}
	}

	return fleet
}
