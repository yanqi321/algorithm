/*
 * @lc app=leetcode.cn id=43 lang=java
 *
 * [43] 字符串相乘
 */
// package leetcode;
// @lc code=start
class Solution {
    public String multiply(String num1, String num2) {
        if (num1.equals("0") || num2.equals("0")) {
            return "0";
        }
        String res = "0";
        for (int i = num2.length() - 1; i >=0; i--) {
            StringBuilder output = new StringBuilder();
            for (int j = 0; j < num2.length() - 1 - i; j++) {
                output.append(0);
            }
            int n2 = num2.charAt(i) - '0';
            // 相乘
            int tmp = 0;
            for (int j = num1.length() - 1; j >=0 || tmp != 0; j--) {
                int n1 = j < 0 ? 0 : num1.charAt(j) - '0';
                int product = (n1 * n2 + tmp) % 10;
                output.append(product);
                tmp = (n1 * n2 + tmp) / 10;
            }
            res = addString(res, output.reverse().toString());
        }

        return res;
    }
    String addString(String num1, String num2) {
        StringBuilder sb = new StringBuilder();
        int tmp = 0;
        for (int i = num1.length() - 1, j = num2.length() - 1;
            i >=0 || j >=0 || tmp !=0; i --, j--
        ) {
            int n1 = i < 0 ? 0 : num1.charAt(i) - '0';
            int n2 = j < 0 ? 0 : num2.charAt(j) - '0';
            int sum = (n1 + n2 + tmp) % 10;
            tmp = (n1 + n2 + tmp) / 10;
            sb.append(sum);
        }
        return sb.reverse().toString();
    }
}
// @lc code=end

