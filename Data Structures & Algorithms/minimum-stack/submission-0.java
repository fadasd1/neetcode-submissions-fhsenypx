class MinStack {
    private List<Integer> stack;
    private Stack<Integer> minima;
    private int min;
    
    public MinStack() {
        stack = new ArrayList<>();
        minima = new Stack<>();
    }
    
    public void push(int val) {
        if(stack.isEmpty()) {min = val; minima.push(val);}

        stack.add(val);
        if(val <= min) {min = val; minima.push(val);}
    }
    
    public void pop() {
        if(top() == min) {minima.pop(); min = minima.peek();}
        stack.removeLast();
    }
    
    public int top() {
        return stack.getLast();
    }
    
    public int getMin() {
        return min;
    }
}
