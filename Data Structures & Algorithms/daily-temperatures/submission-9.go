func dailyTemperatures(temperatures []int) []int {
    var stack []int
    result := make([]int,len(temperatures))
    
    for i:= 0 ;i < len(temperatures);i++{
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            result[idx] = i - idx
        }
        stack = append(stack,i)
    }

    return result
}
