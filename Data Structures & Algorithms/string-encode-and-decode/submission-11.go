type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var plainText strings.Builder

	for _, str := range strs {
		length := fmt.Sprintf("%04d",len(str))
		plainText.WriteString(length)
		plainText.WriteString(str)
	}

	return plainText.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string

	i := 0
	for i < len(encoded) {
		lengthStr := encoded[i:i+4]
		length, _ := strconv.Atoi(lengthStr)

		i += 4
		result = append(result,encoded[i:i+length])

		i += length
	}

	return result
}
