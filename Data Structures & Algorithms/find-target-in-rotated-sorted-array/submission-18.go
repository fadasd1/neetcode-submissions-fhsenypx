func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    
    for l <= r {
        m := l + (r-l) / 2
        if nums[m] == target {
            return m
        }

        
        if nums[l] <= nums[m] {
            if target > nums[m] || target < nums[l] {
                l = m + 1
            } else {
                r = m 
            }
                
    } else {
        if target < nums[m] || target > nums[r] {
            r = m 
        } else {
            l = m + 1
        }
    }
    }

    return -1
}


