func isOperator(t string) bool {
	if t == "+" || t == "-" || t == "*" || t == "/" {
		return true
	}

	return false
}

func operatorNum(t string, a int, b int) int {
	switch t {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func evalRPN(tokens []string) int {
	var stack []int

	for _, t := range tokens {
		if isOperator(t) {
			if len(stack) > 1 {
				result := operatorNum(t, stack[len(stack)-2], stack[len(stack)-1])
				stack[len(stack)-2] = result
				stack = stack[:len(stack)-1]
			}
		} else {
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		}
	}

	return stack[0]

}
