/*
 * @lc app=leetcode.cn id=111 lang=java
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public int minDepth(TreeNode root) {
        int deep = dfs(root);
        return deep;
    }
    private int dfs(TreeNode node) {
        if (node == null) {
            return 0;
        }
        if (node.left == null && node.right == null) {
            return 1;
        }
        int deepL = dfs(node.left);
        int deepR = dfs(node.right);
        if (deepL == 0) {
            return deepR + 1;
        } else if (deepR == 0) {
            return deepL + 1;
        }
        return Math.min(deepL, deepR) + 1;
    }
}
// @lc code=end

