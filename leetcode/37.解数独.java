/*
 * @lc app=leetcode.cn id=37 lang=java
 *
 * {37] 解数独
 */

// @lc code=start
class Solution {
    public void solveSudoku(char[][] board) {
        boolean[][] row = new boolean[9][10];
        boolean[][] col = new boolean[9][10];
        boolean[][] box = new boolean[9][10];
        for (int i=0; i < 9; i++) {
            for (int j=0; j < 9; j++) {
                if (board[i][j] != '.') {
                    int num = board[i][j] - 48;
                    row[i][num] = true;
                    col[j][num] = true;
                    box[j / 3 + i / 3 * 3][num] = true;
                }
            }
        }
        recursiveSudoku(board, row, col, box, 0, 0);
    }
    private boolean recursiveSudoku(char[][] board, boolean[][] row,boolean[][] col, boolean[][] box, int i, int j) {
        if (j == board[0].length) {
            j = 0;
            i++;
            if (i == board.length) {
                return true;
            }
        }
        if (board[i][j] == '.') {
            for (int n = 1; n < 10; n++) {
                if (row[i][n] || col[j][n] || box[j / 3 + i / 3 * 3][n]) continue;
                row[i][n] = true;
                col[j][n] = true;
                box[j / 3 + i / 3 * 3][n] = true;
                board[i][j] = (char)(n + 48);
                if (recursiveSudoku(board, row, col, box, i, j + 1)) {
                    return true;
                }
                row[i][n] = false;
                col[j][n] = false;
                box[j / 3 + i / 3 * 3][n] = false;
                board[i][j] = '.';
            }
            return false;
        } else {
            return recursiveSudoku(board, row, col, box, i, j + 1);
        }
    }
}
// @lc code=end

