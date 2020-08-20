/*
 * @lc app=leetcode.cn id=36 lang=java
 *
 * [36] 有效的数独
 */

// @lc code=start
class Solution {
    public boolean isValidSudoku(char[][] board) {
        int[][] row = new int[9][10];
        int[][] col = new int[9][10];
        int[][] box = new int[9][10];
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] == '.') {
                    continue;
                }
                int num = board[i][j] - 48;
                if (row[i][num] == 1) return false;
                row[i][num] = 1;
                if (col[j][num] == 1) return false;
                col[j][num] = 1;
                if (box[i / 3 * 3 + j / 3][num] == 1) return false;
                box[i / 3 * 3 + j / 3][num] = 1;
            }
        }
        return true;
    }
}
// @lc code=end

