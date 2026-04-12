class Solution {
        public int evalRPN(String[] tokens) {
		Stack<Integer> stack = new Stack<>();
		for(String s : tokens) {
			if(isInteger(s)) {
				stack.push(Integer.parseInt(s));
			} else if (s.equals("+")) {
				int y = stack.pop();
				int x = stack.pop();
				stack.push(x + y);
			} else if (s.equals("*")) {
				int y = stack.pop();
				int x = stack.pop();
				stack.push(x * y);
			} else if (s.equals("-")) {
				int y = stack.pop();
				int x = stack.pop();
				stack.push(x - y);
			} else if (s.equals("/")) {
				int y = stack.pop();
				int x = stack.pop();
				stack.push(x / y);
			}
		}
		return stack.pop();
	}
	
	public static boolean isInteger(String s) {
		if (s == null) return false;
		try {
			Integer.parseInt(s);
			return true;
		} catch (NumberFormatException e) {
			return false;
		}
	}

}