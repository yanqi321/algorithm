import java.util.Deque;
import java.util.HashMap;
import java.util.LinkedList;

/*
 * @lc app=leetcode.cn id=20 lang=java
 *
 * [20] 有效的括号
 */

// @lc code=start
class Solution {
    private final HashMap<Character, Character> pairs = new HashMap<Character, Character>() {
        private static final long serialVersionUID = 1L;

        {
        put(')', '(');
        put(']', '[');
        put('}', '{');
    }};
    public boolean isValid(final String s) {
        final int len = s.length();
        if (len % 2 != 0) {
            return false;
        }
        final Deque<Character> stack = new LinkedList<Character>();
        for (int i = 0; i < len; i++) {
            final char c = s.charAt(i);
            if (pairs.containsKey(c)) {
                if (stack.isEmpty() || stack.peek() != pairs.get(c)) {
                    return false;
                }
                stack.pop();
            } else {
                stack.push(c);
            }
        }
        return stack.isEmpty();
    }
}
// @lc code=end

