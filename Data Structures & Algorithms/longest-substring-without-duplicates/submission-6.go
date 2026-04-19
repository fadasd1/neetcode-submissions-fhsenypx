func lengthOfLongestSubstring(s string) int {
    max := 0

    l := 0
    m := make(map[rune]bool)
    for r, val := range s {
        
        for m[val] == true {
            m[rune(s[l])] = false
            l++
        }
        m[val] = true
        
        if r - l + 1 > max {
            max = r - l + 1
        }

        
    }

    return max
}
