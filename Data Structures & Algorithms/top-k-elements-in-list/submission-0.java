
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
		Map<Integer, Integer> map = new HashMap<>();
		int[] result = new int[k];
		for(int n : nums) {
			map.put(n, map.getOrDefault(n, 0) + 1);
		}

		List<Map.Entry<Integer, Integer>> list = new ArrayList<>(map.entrySet());

		list.sort((a,b) -> b.getValue().compareTo(a.getValue()));

		for(int i = 0; i < k; i++) {
			result[i] = list.get(i).getKey();
		}
		
		return result;
	}
}
