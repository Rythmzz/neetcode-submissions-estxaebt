func evalRPN(tokens []string) int {
	var stack []int
	
	for i:= 0 ;i< len(tokens);i++{
		ch:= tokens[i]

		if isOperand(ch) {
			if len(stack) > 1 {
				stack[len(stack)-2] = operators(ch,stack[len(stack)-2],stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		} else {
			num,_ := strconv.Atoi(ch)
			stack = append(stack,num)
		}
	}

	return stack[0]
}

func isOperand(s string) bool{
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}

	return false
}

func operators(t string, a,b int) int {
	switch t {
		case "+": return a + b
		case "*": return a * b
		case "/": return a / b
		case "-": return a - b
	}

	return 0
}
