func evalRPN(tokens []string) int {
	st := []int{}

	for i:= 0 ;i< len(tokens);i++ {
		c := tokens[i]
		if isOp(c) {
			if len(st) > 1 {
				st[len(st)-2] = operands(c, st[len(st)-2],st[len(st)-1])
				st = st[:len(st)-1]
			}
		} else {
			num, _ := strconv.Atoi(c)
			st = append(st, num)
		}
	}

	return st[0]
}

func isOp(c string) bool {
	if c == "+" || c == "-" || c == "*" || c == "/" {
		return true
	}

	return false
}

func operands(c string, a, b int) int {
	switch c {
		case "+" : return a + b
		case "*" : return a * b
		case "-" : return a - b
		case "/" : return a / b
	}

	return 0
}
