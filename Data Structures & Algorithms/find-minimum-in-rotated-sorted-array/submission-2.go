func findMin(nums []int) int {
    l, r := 0, len(nums)-1
    min := nums[0]
    for l < r {
        m := l + (r-l)/2
        if nums[m] < min {
            min = nums[m]
        }
        if nums[m] > nums[r] {
            l = m + 1
            if l == r {
                return nums[l]
            }
        } else {
            r = m
            if l == r {
                return nums[l]
            }
        }
    }
    return min
}