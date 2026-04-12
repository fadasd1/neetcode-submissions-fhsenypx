class Solution {
      	public boolean isValid(String s) {
		Stack<Character> stack = new Stack<>();

		for(char c : s.toCharArray()) {
			if(c == '{' || c == '(' || c == '[') stack.push(c);
			else {
				if(stack.isEmpty()) return false;

				char current = stack.pop();
				if(
					c == '}' && current != '{' ||
					c == ']' && current != '[' ||
					c == ')' && current != '('		
				) return false;
			}

		}

		return stack.isEmpty();
	}
                    }

