func largestRectangleArea(heights []int) int {
    r := make([]int,len(heights))
    l := make([]int,len(heights))
    maxRectangle := 0

    for i:= 0 ;i< len(heights);i++ {
        l[i] = -1
        r[i] = len(heights)
    }

    stack := []int{}
    for i:= 0 ; i< len(heights);i++{
        for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            r[idx] = i
        }
        stack = append(stack,i)
    }

    stack = []int{}
    for i:= len(heights) - 1; i >= 0; i-- {
        for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            l[idx] = i
        }
        stack = append(stack,i)
    }

    for i:= 0; i < len(heights);i++{
        rectangle := (r[i] - l[i] - 1) * heights[i]
        maxRectangle = max(rectangle,maxRectangle)
    }

    return maxRectangle
}
