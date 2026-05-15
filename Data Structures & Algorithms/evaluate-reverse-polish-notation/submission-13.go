func evalRPN(tokens []string) int {
	st := []int{}

	for i:= 0 ;i< len(tokens);i++ {
		c := tokens[i]
		if isOperators(c) {
			if len(st) > 1 {
				st[len(st)-2] = operands(st[len(st)-2],st[len(st)-1],c)
				st = st[:len(st)-1]
			} else {
				return 0
			}
		} else {
			num, _ := strconv.Atoi(c)
			st = append(st,num)
		}
	}

	return st[0]

}

func isOperators(c string) bool {
	if c == "+" || c == "-" || c == "*" || c == "/" {
		return true
	}

	return false
}

func operands(a,b int, c string) int {
	switch c {
		case "+" : return a + b
		case "-" : return a - b
		case "*" : return a * b
		case "/" : return a / b
	}

	return 0
}


