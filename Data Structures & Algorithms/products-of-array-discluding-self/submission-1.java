class Solution {
    public int[] productExceptSelf(int[] nums) {
		int product = 1;
		int count = 0;
		int latestZero = 0;
		for(int i = 0; i < nums.length; i++) {
			if(nums[i] != 0 ) product *= nums[i];
			else {count++; latestZero = i;}
		}
		if(count > 1) return new int[nums.length];
		int[] result = new int[nums.length];
		
		if(count == 1) {
			result[latestZero] = product;
		}
		
		if(count == 0) {
			for(int i = 0; i < nums.length; i++) {
				result[i] = product / nums[i];
			}
		}


		return result;
	}
}  
