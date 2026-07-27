type Solution struct{}

const sep = '|'

func (s *Solution) Encode(strs []string) string {
	buf := make([]byte, 0, len(strs)*4)

	buf = strconv.AppendInt(buf, int64(len(strs)), 10)
	buf = append(buf, sep)

	for _, str := range strs {
		buf = strconv.AppendInt(buf, int64(len(str)), 10)
		buf = append(buf, sep)
		buf = append(buf, str...)
	}

	return string(buf)
}

func (s *Solution) Decode(encoded string) []string {
	i := 0

	count, i := readNumber(encoded, i)
	result := make([]string, count)

	for index := 0; index < count; index++ {
		length, next := readNumber(encoded, i)
		i = next

		end := i + length
		result[index] = encoded[i:end]
		i = end
	}
	return result
}

func readNumber(encoded string, i int) (number, next int) {
	for encoded[i] != sep {
		number = number*10 + int(encoded[i]-'0')
		i++
	}
	return number, i + 1
}
