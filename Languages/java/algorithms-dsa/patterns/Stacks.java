import java.util.*;

/**
 * Stack Patterns
 * 
 * Key insight: Stacks excel at problems involving matching (parentheses),
 * maintaining monotonic order, and evaluating expressions.
 * 
 * Common patterns:
 * 1. Monotonic Stack: Find next greater/smaller element in O(n)
 * 2. Daily Temperatures: Monotonic decreasing stack variant
 * 3. Valid Parentheses: Classic stack matching problem
 * 4. Largest Rectangle in Histogram: Monotonic increasing stack
 * 5. Evaluate RPN: Operand stack with operator application
 * 
 * Monotonic Stack Trick:
 * - Maintain a stack where elements are in sorted order
 * - When a new element violates the order, pop and process
 * - Each element is pushed/popped at most once → O(n) total
 */
public class Stacks {

    // ==================== Next Greater Element ====================
    /**
     * For each element, find the next greater element to its right.
     * Strategy: Traverse right to left, maintain a monotonic decreasing stack.
     * The stack holds candidates for "next greater". Pop smaller elements.
     * 
     * Time: O(n), Space: O(n)
     */
    public static int[] nextGreaterElement(int[] nums) {
        int n = nums.length;
        int[] result = new int[n];
        Deque<Integer> stack = new ArrayDeque<>(); // Stores values

        // Process from right to left
        for (int i = n - 1; i >= 0; i--) {
            // Pop elements that are not greater than current
            while (!stack.isEmpty() && stack.peek() <= nums[i]) {
                stack.pop();
            }
            result[i] = stack.isEmpty() ? -1 : stack.peek();
            stack.push(nums[i]);
        }

        return result;
    }

    // ==================== Daily Temperatures ====================
    /**
     * Given daily temperatures, find how many days until a warmer temperature.
     * Strategy: Monotonic decreasing stack storing indices.
     * When we find a warmer day, pop and calculate the distance.
     * 
     * Time: O(n), Space: O(n)
     */
    public static int[] dailyTemperatures(int[] temperatures) {
        int n = temperatures.length;
        int[] result = new int[n];
        Deque<Integer> stack = new ArrayDeque<>(); // Stores indices

        for (int i = 0; i < n; i++) {
            // Pop all days that are colder than today
            while (!stack.isEmpty() && temperatures[stack.peek()] < temperatures[i]) {
                int prevDay = stack.pop();
                result[prevDay] = i - prevDay;
            }
            stack.push(i);
        }

        return result; // Remaining indices in stack have answer 0 (no warmer day)
    }

    // ==================== Valid Parentheses ====================
    /**
     * Check if a string of brackets is valid (properly nested and matched).
     * Strategy: Push opening brackets, pop and match for closing brackets.
     * 
     * Time: O(n), Space: O(n)
     */
    public static boolean isValidParentheses(String s) {
        Deque<Character> stack = new ArrayDeque<>();

        for (char c : s.toCharArray()) {
            if (c == '(' || c == '{' || c == '[') {
                stack.push(c);
            } else {
                if (stack.isEmpty()) return false;
                char top = stack.pop();
                if (c == ')' && top != '(') return false;
                if (c == '}' && top != '{') return false;
                if (c == ']' && top != '[') return false;
            }
        }

        return stack.isEmpty();
    }

    // ==================== Largest Rectangle in Histogram ====================
    /**
     * Find the largest rectangular area in a histogram.
     * Strategy: Monotonic increasing stack of indices.
     * When a bar is shorter than the stack's top, pop and calculate area.
     * The popped bar's width extends from the new stack top to current index.
     * 
     * Time: O(n), Space: O(n)
     */
    public static int largestRectangleInHistogram(int[] heights) {
        int n = heights.length;
        int maxArea = 0;
        Deque<Integer> stack = new ArrayDeque<>(); // Indices of increasing heights

        for (int i = 0; i <= n; i++) {
            int currHeight = (i == n) ? 0 : heights[i]; // Sentinel to flush stack

            while (!stack.isEmpty() && currHeight < heights[stack.peek()]) {
                int height = heights[stack.pop()];
                // Width: from current stack top + 1 to i - 1
                int width = stack.isEmpty() ? i : (i - stack.peek() - 1);
                maxArea = Math.max(maxArea, height * width);
            }
            stack.push(i);
        }

        return maxArea;
    }

    // ==================== Evaluate Reverse Polish Notation ====================
    /**
     * Evaluate an arithmetic expression in Reverse Polish Notation.
     * Strategy: Push numbers onto stack. For operators, pop two operands,
     * compute, and push the result back.
     * 
     * Time: O(n), Space: O(n)
     */
    public static int evalRPN(String[] tokens) {
        Deque<Integer> stack = new ArrayDeque<>();

        for (String token : tokens) {
            switch (token) {
                case "+":
                    stack.push(stack.pop() + stack.pop());
                    break;
                case "-":
                    int b = stack.pop(), a = stack.pop();
                    stack.push(a - b);
                    break;
                case "*":
                    stack.push(stack.pop() * stack.pop());
                    break;
                case "/":
                    int divisor = stack.pop(), dividend = stack.pop();
                    stack.push(dividend / divisor);
                    break;
                default:
                    stack.push(Integer.parseInt(token));
            }
        }

        return stack.pop();
    }

    // ==================== Tests ====================
    public static void main(String[] args) {
        // Test Next Greater Element
        assert Arrays.equals(
            nextGreaterElement(new int[]{4, 5, 2, 25}),
            new int[]{5, 25, 25, -1}
        ) : "Next greater test 1 failed";
        assert Arrays.equals(
            nextGreaterElement(new int[]{13, 7, 6, 12}),
            new int[]{-1, 12, 12, -1}
        ) : "Next greater test 2 failed";

        // Test Daily Temperatures
        assert Arrays.equals(
            dailyTemperatures(new int[]{73, 74, 75, 71, 69, 72, 76, 73}),
            new int[]{1, 1, 4, 2, 1, 1, 0, 0}
        ) : "Daily temps test failed";
        assert Arrays.equals(
            dailyTemperatures(new int[]{30, 40, 50, 60}),
            new int[]{1, 1, 1, 0}
        ) : "Daily temps test 2 failed";

        // Test Valid Parentheses
        assert isValidParentheses("(){}[]") : "Parentheses test 1 failed";
        assert isValidParentheses("({[]})") : "Parentheses test 2 failed";
        assert !isValidParentheses("(]") : "Parentheses test 3 failed";
        assert !isValidParentheses("([)]") : "Parentheses test 4 failed";
        assert isValidParentheses("") : "Parentheses test 5 failed";

        // Test Largest Rectangle in Histogram
        assert largestRectangleInHistogram(new int[]{2, 1, 5, 6, 2, 3}) == 10
            : "Histogram test 1 failed";
        assert largestRectangleInHistogram(new int[]{2, 4}) == 4
            : "Histogram test 2 failed";
        assert largestRectangleInHistogram(new int[]{1, 1, 1, 1}) == 4
            : "Histogram test 3 failed";

        // Test Evaluate RPN
        assert evalRPN(new String[]{"2", "1", "+", "3", "*"}) == 9
            : "RPN test 1 failed";
        assert evalRPN(new String[]{"4", "13", "5", "/", "+"}) == 6
            : "RPN test 2 failed";
        assert evalRPN(new String[]{"10", "6", "9", "3", "+", "-11", "*",
            "/", "*", "17", "+", "5", "+"}) == 22
            : "RPN test 3 failed";

        System.out.println("All stack pattern tests passed!");
    }
}
