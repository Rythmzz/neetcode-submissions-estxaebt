func checkInclusion(s1 string, s2 string) bool {
  if len(s1) > len(s2) {
	return false
  }

  result := make(map[[26]int][]string)

  var freq[26]int
  for _,ch:=range s1 {
	freq[ch-'a']++
  }
  result[freq] = append(result[freq],s1)

 var resultStr []string
  for i := 0 ;i < len(s2);i++{
	length := i + len(s1)
	if i + len(s1) <= len(s2) {
		resultStr = append(resultStr,s2[i:length])
	} else {
		break
	}
  }

  for _, str:= range resultStr {
	var freqq[26]int
	
	for _, ch:= range str {
		freqq[ch-'a']++
	}

	result[freqq] = append(result[freqq],str)
  }

  if len(result[freq]) > 1 {
	return true
  } 

  return false


}
