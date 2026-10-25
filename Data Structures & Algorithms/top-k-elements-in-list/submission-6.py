class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = {}

        buckets = []
        for n in range(len(nums)):
            buckets.append(set())
        
        for n in nums:
            freq[n] = freq.get(n, -1) + 1
        for key, v in freq.items():
            buckets[v].add(key)

        res = []
    
        i = len(buckets) - 1
        while True:
            
            for n in buckets[i]:
                res.append(n)
                
                if len(res) == k:
                    return res
                    
            i -= 1
        return res



        