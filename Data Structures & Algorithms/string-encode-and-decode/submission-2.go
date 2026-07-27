type Solution struct{}

const sep = '|'

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder

	size := len(strconv.Itoa(len(strs))) + 1
	for _, str := range strs {
		size += len(strconv.Itoa(len(str))) + 1 + len(str)
	}
	b.Grow(size)

	b.WriteString(strconv.Itoa(len(strs)))
	b.WriteByte(sep)

	for _, str := range strs {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteByte(sep)
		b.WriteString(str)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	i := 0

	count, next := readNumber(encoded, i)
	i = next

	result := make([]string, 0, count)

	for range count {
		length, next := readNumber(encoded, i)
		i = next

		end := i + length
		result = append(result, encoded[i:end])
		i = end
	}

	return result
}

func readNumber(encoded string, i int) (int, int) {
	number := 0

	for encoded[i] != sep {
		number = number*10 + int(encoded[i]-'0')
		i++
	}

	return number, i + 1
}
