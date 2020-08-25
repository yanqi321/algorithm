import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/*
 * @lc app=leetcode.cn id=127 lang=java
 *
 * [127] 单词接龙
 */

// @lc code=start
class Solution {
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        if (!wordList.contains(endWord)) {
            return 0;
        }
        // HashSet<String> visited = new HashSet<String>();
        boolean[] visited = new boolean[wordList.size()];
        int idx = wordList.indexOf(beginWord);
        if ( idx != -1) {
            visited[idx] = true;
        }
        Queue<String> queue = new LinkedList<String>();
        queue.offer(beginWord);
        int count = 0;
        while(queue.size() > 0) {
            count++;
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                String start = queue.poll();
                for (int j = 0; j < wordList.size(); j++) {
                    if (visited[j]) {
                        continue;
                    }
                    String s = wordList.get(j);
                    if (!canConvert(start, s)) {
                        continue;
                    }
                    if (endWord.equals(s)) {
                        return count + 1;
                    }
                    queue.offer(s);
                    visited[j] = true;
                }
            }
        }
        return 0;
    }
    private boolean canConvert(String start, String c) {
        if (start.length() != c.length()) {
            return false;
        }
        int diffCount = 0;
        for (int i = 0; i < c.length(); i++) {
            if (start.charAt(i) != c.charAt(i)) {
                diffCount++;
                if (diffCount > 1) return false;
            } 
        }
        return diffCount == 1;
    }
}
// @lc code=end

