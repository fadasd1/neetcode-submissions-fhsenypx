class Solution {
    public int[] productExceptSelf(int[] nums) {
		int[] result = new int[nums.length];
		
		for(int i = 0; i < nums.length; i++) {
			int product = 1;
			for(int j = 0; j < nums.length; j++) {
				if(i == j) continue;
				if(nums[j] == 0) {product = 0; break;}
				product *= nums[j];
			}
			result[i] = product;
		}
		
		return result;
	}
}  
