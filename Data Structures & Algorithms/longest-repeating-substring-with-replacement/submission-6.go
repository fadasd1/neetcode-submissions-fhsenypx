func characterReplacement(s string, k int) int {
    m := make(map[byte]int)

    
        maxLen := 0
        
        l := 0
        maxFrequ := 0
        

        for r, _ := range s {
            m[s[r]]++
            if m[s[r]] > maxFrequ {
                maxFrequ = m[s[r]]
            }

            for r - l + 1 - maxFrequ > k {
                m[s[l]]--
                l++
            } 
            if r - l + 1 > maxLen {
                    maxLen = r - l + 1
                }

        }
        return maxLen
    }
    

