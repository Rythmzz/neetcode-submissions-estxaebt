type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var buildString strings.Builder

    for _, str := range strs {
        length := len(str)
        buildString.WriteString(fmt.Sprintf("%04d",length))
        buildString.WriteString(str)
    }

    return buildString.String()
}

func (s *Solution) Decode(encoded string) []string {
    var result []string

    i:= 0
    for i < len(encoded) {
        lengthStr := encoded[i:i+4]
        length, _ := strconv.Atoi(lengthStr)

        i += 4
        result = append(result,encoded[i:i+length])

        i += length
    }

    return result
}
