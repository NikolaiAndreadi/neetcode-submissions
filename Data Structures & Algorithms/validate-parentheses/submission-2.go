func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s))
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		openBracket, isClosingBracket := bracketMap[c]

		if !isClosingBracket {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastIndex := len(stack) - 1
		if stack[lastIndex] != openBracket {
			return false
		}

		stack = stack[:lastIndex]
	}

	return len(stack) == 0
}