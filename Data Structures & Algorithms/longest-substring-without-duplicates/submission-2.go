func lengthOfLongestSubstring(s string) int {
    max := 0
    trueMax := 0
    l := 0
    m := make(map[rune]bool)
    for _, val := range s {
        
        for m[val] == true {
            m[rune(s[l])] = false
            l++
            max--
        }
        m[val] = true
        max++
        if max > trueMax {
            trueMax = max
        }

        
    }

    return trueMax
}
