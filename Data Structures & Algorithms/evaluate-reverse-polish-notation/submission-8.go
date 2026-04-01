func evalRPN(tokens []string) int {
	var stack []int

	for i:= 0 ; i< len(tokens);i++ {
		if isOperators(tokens[i]) {
			if len(stack) > 1 {
				stack[len(stack)-2] = operandsNum(tokens[i],stack[len(stack)-2],stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		} else {
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack,num)
		}
	}

	return stack[0]
}

func isOperators(ch string) bool {
	if ch == "+" || ch == "-" || ch == "*" || ch == "/" {
		return true
	}

	return false
}

func operandsNum(c string, a int, b int) int{
	switch c {
		case "+": return a + b
		case "-": return a - b
		case "*": return a * b
		case "/": return a / b
	}

	return 0
}
