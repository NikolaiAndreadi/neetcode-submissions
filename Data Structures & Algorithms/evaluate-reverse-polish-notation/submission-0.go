func evalRPN(tokens []string) int {
	type operation func(l int, r int) int
	operations := map[string]operation {
		"+": func(l int, r int) int { return l + r},
		"-": func(l int, r int) int { return l - r},
		"*": func(l int, r int) int { return l * r},
		"/": func(l int, r int) int { return l / r},
	}
	stack := make([]int, 0)
	for _, token := range tokens {
		op, isOp := operations[token]
		if !isOp {
			intToken, _ := strconv.Atoi(token)
			stack = append(stack, intToken)
			continue
		}
		r := stack[len(stack)-1]
		l := stack[len(stack)-2]
		result := op(l, r)
		stack[len(stack)-2] = result
		stack = stack[:len(stack)-1]
	}
	return stack[0]
}
