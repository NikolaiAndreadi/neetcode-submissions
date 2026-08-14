func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}

    stack := make([]string, 0)
	bracketMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for _, c := range s {
		cs := string(c)
		corr, ok := bracketMap[cs]
		if !ok {
			stack = append(stack, cs)
			continue
		}

		sl := len(stack)-1
		if sl < 0 {
			return false
		}

		last := stack[sl]
		if corr != last {
			return false
		}

		stack = stack[:sl]
	}
	return len(stack) == 0
}
