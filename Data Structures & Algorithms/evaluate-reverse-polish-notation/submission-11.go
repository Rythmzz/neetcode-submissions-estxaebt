func evalRPN(tokens []string) int {
	var stack []int

	for i:=0;i<len(tokens);i++ {
		c := tokens[i]
		if isOperators(c) {
			if len(stack) <= 1 {
				return 0
			}

			stack[len(stack)-2] = operands(c,stack[len(stack)-2],stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else {
			num, _:= strconv.Atoi(c)
			stack = append(stack,num)
		}
	}

	return stack[0]
}

func isOperators(c string) bool {
	if c == "+" || c == "-" || c == "*" || c == "/" {
		return true
	}

	return false
}

func operands(s string, a,b int) int {
	switch s {
		case "+": return a+b
		case "-": return a-b
		case "*": return a*b
		case "/": return a/b
	}

	return 0
}
