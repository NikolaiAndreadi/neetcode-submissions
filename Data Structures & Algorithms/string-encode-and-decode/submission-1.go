type Solution struct{}

const sep = '|'

func (s *Solution) Encode(strs []string) string {
	var b strings.Builder
	for _, str := range strs {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteByte(sep)
		b.WriteString(str)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := make([]string, 0, 0)
	if len(encoded) == 0 {
		return result
	}

	i := 0
	for j := 0; j < len(encoded); j++ {
		if encoded[j] != sep {
			continue
		}
		substrLen, _ := strconv.Atoi(encoded[i:j])
		result = append(result, encoded[j+1:j+substrLen+1])
		i = j+substrLen+1
		j = i
	}
	return result
}
