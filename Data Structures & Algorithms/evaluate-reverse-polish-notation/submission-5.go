func evalRPN(tokens []string) int {
	var stack []int

	for i:= 0 ;i<len(tokens);i++{
		if isOperator(tokens[i]) {
			if len(stack) > 1 {
				stack[len(stack)-2] = operationNum(tokens[i],stack[len(stack)-2],stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		} else {
			num,_ := strconv.Atoi(tokens[i])
			stack = append(stack,num)
		}
	}

	return stack[0]
}

func isOperator(t string) bool {
	if t == "+" ||  t ==  "-" || t == "*" || t == "/" {
		return true
	}
	
	return false
}

func operationNum(t string, a, b int) int {
	switch t {
		case "+" : return  a + b
		case "-" : return  a - b
		case "*" : return a * b
		case "/" :  return a / b
	}

	return 0
}
