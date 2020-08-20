/*
 * @lc app=leetcode.cn id=529 lang=java
 *
 * [529] 扫雷游戏
 */

// @lc code=start
class Solution {
    int[] dx = { -1, 1, 0, 0, -1, 1, -1, 1 };
    int[] dy = { 0, 0, -1, 1, -1, -1, 1, 1 };
    public char[][] updateBoard(char[][] board, int[] click) {
        int x = click[0], y = click[1];
        if (board[x][y] == 'M') {
            board[x][y] = 'X';
        } else {
            dfs(board, x, y);
        }
        return board;
    }
    private void dfs(char[][] board, int x,int y) {
        int cnt = 0;
        for (int i = 0; i < 8; i++) {
            int xx = x + dx[i];
            int yy = y + dy[i];
            if (xx < 0 || xx >= board.length || yy < 0 || yy >= board[0].length) {
                continue;
            }
            if (board[xx][yy] == 'M') {
                cnt++;
            }
        }
        if (cnt > 0) {
            board[x][y] = (char)(cnt + 48);
            return;
        }
        board[x][y] = 'B';
        for (int i = 0; i < 8; i++) {
            int xx = x + dx[i];
            int yy = y + dy[i];
            if (xx < 0 || xx >= board.length || yy < 0 || yy >= board[0].length || board[xx][yy] != 'E') {
                continue;
            }
            dfs(board, xx, yy);
        }
    }
}
// @lc code=end

