func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    maxAreaWater := 0

    n := len(height)
    heightsLeft := make([]int,n)
    heightsRight := make([]int, n)
    heightsLeft[0] = height[0]
    heightsRight[n-1] = height[n-1]


    for i:= 1 ; i < n;i++{
        heightsLeft[i] = max(height[i],heightsLeft[i-1])
    }

    for i:= n-2;i >= 0;i-- {
        heightsRight[i] = max(height[i],heightsRight[i+1])
    }

    for i:= 0 ; i < n; i++ {
        maxAreaWater += min(heightsRight[i],heightsLeft[i]) - height[i]
    }

    return maxAreaWater
}
