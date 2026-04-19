func characterReplacement(s string, k int) int {
    counts := make([]int, 26)
    l, maxFrequ, maxLen := 0, 0, 0

    for r := 0; r < len(s); r++ {
        counts[s[r] - 65]++

        maxFrequ = max(counts[s[r] - 65], maxFrequ)

    
        for (r - l + 1) - maxFrequ > k {
            counts[s[l] - 65]--
            l++
        }

        maxLen = max(maxLen, r - l + 1)
    }
    return maxLen
}