func characterReplacement(s string, k int) int {
    m := make(map[byte]struct{})

    for i, _ := range s {
        m[s[i]] = struct{}{}
    }
    max := 0
    for c := range m {
        list := make([]int, 0)
        index := 0
        l := 0
        count := 0
        

        for r, _ := range s {
            if s[r] != c {
                count++
                list = append(list, r)
            }
            if count > k {
             
                l = list[index] + 1
                index++
                count--
            }
            if r - l + 1 > max {
                    max = r - l + 1
                }

        }
    }
    return max
}
