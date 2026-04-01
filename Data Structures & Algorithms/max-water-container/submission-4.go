func maxArea(heights []int) int {
    maxHeight := 0

    i,j:= 0, len(heights)-1

    for i < j {
        waterCal := (j-i) * min(heights[i],heights[j])
        maxHeight = max(maxHeight, waterCal)

        if heights[i] > heights[j] {
            j--
        } else {
            i++
        }
    }

    return maxHeight
}
