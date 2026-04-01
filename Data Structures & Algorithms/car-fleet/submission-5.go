func carFleet(target int, position []int, speed []int) int {
	// [4,1,0,7]
	// [3,4.5,10,3]
	// 7,4 => 1 fleet
	// 1
	var stack []float32
	for i:= 0; i< len(position);i++{
		for j:= i+1; j <  len(position);j++{
			if position[i] < position[j] {
				tPosition := position[i]
				tSpeed := speed[i]
				
				position[i] = position[j]
				speed[i] = speed[j]

				position[j] = tPosition
				speed[j] = tSpeed
			}
		}
	}

	targetPositionFirst := float32(target-position[0]) / float32(speed[0])
	stack =  append(stack,targetPositionFirst)
	for i:=0; i < len(position);i++{
		targetPosition := float32(target-position[i]) / float32(speed[i])
		if i > 0 && targetPosition > float32(stack[len(stack)-1]) {
			stack = append(stack,targetPosition)
		}

	}

	return len(stack)

	


}
