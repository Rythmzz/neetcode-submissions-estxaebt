func evalRPN(tokens []string) int {
	var stack []int
	for i:= 0 ;i < len(tokens);i++ {
		if isOperand(tokens[i]) {
			if len(stack) > 1 {
				stack[len(stack)-2] = operationNum(tokens[i],stack[len(stack)-2],stack[len(stack)-1])
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

func isOperand(c string) bool {
	if c == "+" || c == "-" || c == "*" || c == "/" {
		return true
	}
	return false
}

func operationNum(c string, a int, b int) int {
	switch c {
		case "+" : return a + b
		case "-" : return a - b
		case "*" : return a * b
		case "/" : return a / b
	}

	return 0
}
