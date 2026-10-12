class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums: number[], target: number): number[] {
        const seen: Map<number, number> = new Map<number, number>

        for(let i = 0; i < nums.length; i++) {
            const complement = target - nums[i]
            if(seen.has(complement)) {
                return [seen.get(complement), i]
            }
            seen.set(nums[i], i)
        }

        return [0, 0]
    }
}
