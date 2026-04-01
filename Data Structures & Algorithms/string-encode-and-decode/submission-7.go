type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builderStr strings.Builder

	for _, str := range strs {
		length := len(str)
		builderStr.WriteString(fmt.Sprintf("%04d", length))
		builderStr.WriteString(str)
	}

	return builderStr.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	i := 0

	for i < len(encoded) {
		lengthStr := encoded[i : i+4]
		length, _ := strconv.Atoi(lengthStr)

		i += 4
		result = append(result, encoded[i:length+i])

		i += length
	}

	return result
}