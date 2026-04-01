func maxArea(heights []int) int {
    maxWater := 0

    left, right := 0, len(heights)-1

    for left < right {
        calWater :=  (right-left) * min(heights[right],heights[left])

        maxWater = max(maxWater,calWater)

        if heights[right] > heights[left] {
            left++
        } else {
            right--
        }
    }

    return maxWater
}
