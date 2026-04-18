func evalRPN(tokens []string) int {
	stack := []int{}

	for i:= 0;i< len(tokens);i++ {
		ch := tokens[i]
		if isOperator(ch) {
			if len(stack) > 1 {
				stack[len(stack)-2] = operands(ch,stack[len(stack)-2],stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		} else {
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack,num)
		}
	}

	return stack[0]
}

func isOperator(ch string) bool {
	if ch == "+" || ch == "-" || ch == "/" || ch == "*" {
		return true
	}

	return false
}

func operands(c string, a,b int) int {
	switch c {
		case "+": return a + b
		case "-": return a - b
		case "*": return a * b
		case "/": return a / b
	}

	return 0
}
