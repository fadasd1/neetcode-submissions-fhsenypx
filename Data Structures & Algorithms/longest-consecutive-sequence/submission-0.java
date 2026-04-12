class Solution {
    public int longestConsecutive(int[] nums) {
		Set<Integer> set = new HashSet<>();

		for(int n : nums) {
			set.add(n);
		}

		int max = 0;
		int count = 1;
		for(int n : set) {
			while(set.contains(n+1)) {
				count++;
				n += 1;
			}
			if(count > max) max = count;
			count = 1;
		}

		return max;
	}
}
