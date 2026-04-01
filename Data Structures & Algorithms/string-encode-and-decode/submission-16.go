type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encode strings.Builder
	
	for i:= 0 ; i< len(strs); i++ {
		length := fmt.Sprintf("%04d",len(strs[i]))
		encode.WriteString(length)
		encode.WriteString(strs[i])
	}

	return encode.String()
}

func (s *Solution) Decode(encoded string) []string {
	i := 0
	result := []string{}

	for i < len(encoded) {
		lengthStr := encoded[i:i+4]
		length, _ := strconv.Atoi(lengthStr)

		i += 4
		result = append(result,encoded[i:i+length])

		i += length
	}

	return result
}
