package solutions

func IsAnagram(s string, t string) bool {
	lenS := len(s)
	if lenS != len(t) {
		return false
	}

	mapS := make(map[byte]int)
	mapT := make(map[byte]int)

	for i := 0; i < lenS; i++ {
		mapS[s[i]]++
		mapT[t[i]]++
	}

	for k, v := range mapT {
		if mapS[k] != v {
			return false
		}
	}

	return true
}
