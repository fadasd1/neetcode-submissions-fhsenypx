func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	result := reg.ReplaceAllString(s, "")
	result = strings.ToLower(result)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		if result[i] != result[j] {
			return false
		}
	}

	return true
}
