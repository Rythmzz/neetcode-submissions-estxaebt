type Solution struct {
	key map[string]string
}

func (s *Solution) Encode(strs []string) string {
	s.key = map[string]string{}

	textEncode := ""
	for i, str := range strs {

		plainText := fmt.Sprintf("%d", i)
		s.key[plainText] = str
		textEncode += plainText
		if i != len(strs)-1 {
			textEncode += ","
		}

	}

	return textEncode
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	textKey := ""
	for i, num := range encoded {
		if num == ',' {
			result = append(result, s.key[textKey])
			textKey = ""
			continue
		}

		textKey += string(num)

		if i == len(encoded)-1 {
			result = append(result, s.key[textKey])
		}

	}

	return result
}