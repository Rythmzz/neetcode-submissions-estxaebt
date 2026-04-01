func carFleet(target int, position []int, speed []int) int {
	cars := make([][2]int,len(position))
	for i:= 0 ; i < len(position);i++ {
		cars[i][0] = position[i]
		cars[i][1] = speed[i]
	}

	sort.Slice(cars,func(i,j int) bool{
		return cars[i][0] > cars[j][0]
	})

	fleet := 1
	maxTime := (float64(target) - float64(cars[0][0])) / float64(cars[0][1])
	// 3, 3
	for i:= 1 ;i < len(cars);i++{
		time := (float64(target) - float64(cars[i][0])) / float64(cars[i][1])
		if maxTime < time {
			fleet++
			maxTime = time
		}  
	}

	return fleet

}
