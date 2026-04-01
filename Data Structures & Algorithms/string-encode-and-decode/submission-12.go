type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var text strings.Builder

	for _, str:= range strs {
		length := fmt.Sprintf("%04d",len(str))
		text.WriteString(length)
		text.WriteString(str)
	}

	return text.String()
}

func (s *Solution) Decode(encoded string) []string {
	i:=0
	var result []string
	for i < len(encoded) {
		lengthStr := encoded[i:i+4]
		length,_ := strconv.Atoi(lengthStr)

		i += 4
		
		result = append(result,encoded[i:length+i])

		i += length
	}

	return result
}

