package sevenKyu

func GetCount(str string) int {
	count := 0
	for _, ch := range str {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}