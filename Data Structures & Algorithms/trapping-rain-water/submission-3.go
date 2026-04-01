func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    maxAreaWater := 0
    arrMaxLeft :=  make([]int,len(height))
    arrMaxRight :=  make([]int,len(height))

    arrMaxLeft[0] = height[0]
    for i:= 1; i < len(height); i++ {
        arrMaxLeft[i] = max(arrMaxLeft[i-1],height[i])
    }

    arrMaxRight[len(height)-1] = height[len(height)-1]
    for i:= len(height)-2; i >= 0; i-- {
        arrMaxRight[i] = max(arrMaxRight[i+1],height[i])
    }


    for i:= 0 ; i < len(height); i++ {
        maxAreaWater += min(arrMaxLeft[i],arrMaxRight[i]) - height[i]
    }

    return maxAreaWater
}
