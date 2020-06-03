/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (68.71%)
 * Likes:    405
 * Dislikes: 0
 * Total Accepted:    41.9K
 * Total Submissions: 60.1K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 
 * 
 * 
 * 上图为 8 皇后问题的一种解法。
 * 
 * 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
 * 
 * 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 * 
 * 示例:
 * 
 * 输入: 4
 * 输出: [
 * ⁠[".Q..",  // 解法 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 * 
 * ⁠["..Q.",  // 解法 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * 解释: 4 皇后问题存在两个不同的解法。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 
 * 皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步，可进可退。（引用自
 * 百度百科 - 皇后 ）
 * 
 * 
 */

// @lc code=start
type queen struct {
	n int
	stack map[int]int
	col []int
	pie map[int]int
	na map[int]int
	ret [][]string
}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	if n == 0 {
		return res
	}
	q := NewQueen(n)
	q.backtrack(0)
	return q.ret
}

func NewQueen(n int) *queen {
	stack := make(map[int]int, n)
	col := make([]int, n)
	pie := make(map[int]int, n)
	na := make(map[int]int, n)
	return &queen{ n, stack, col ,pie, na, [][]string{}}
}

func (q *queen) backtrack(row int) {
	if q.n < 1 {
		return
	}
	for col :=0; col < q.n; col++ {
		if q.isNotUnderAttack(row, col) {
			q.addPos(row, col)
			if row + 1 == q.n {
				q.addSolution()
			} else {
				q.backtrack(row + 1)
			}
			q.rmPos(row, col)
		}
	}
}

func (q *queen) isNotUnderAttack(row, col int) bool {
	res := q.col[col] + q.pie[row + col] + q.na[col - row + 2 * q.n ]
	// fmt.Println(row, res)
	if res == 0 {
		return true
	}
	return false
}

func (q *queen) addPos(row, col int) {
	q.stack[row] = col
	q.col[col] = 1
	q.pie[row + col] = 1
	q.na[col - row + 2 * q.n] = 1
}
func (q *queen) rmPos(row, col int) {
	delete(q.stack, row)
	q.col[col] = 0
	q.pie[row + col] = 0
	q.na[col - row + 2 * q.n] = 0
}
func (q *queen) addSolution(){
	solution := make([]string, q.n)
	for row := 0; row < q.n; row++ {
		rowCol := q.stack[row]
		output := ""
		for col := 0; col < q.n; col++ {
			if rowCol == col {
				output += "Q"
			} else {
				output += "."
			}
		}
		solution[row] = output
	}
	q.ret = append(q.ret, solution)
}

// @lc code=end

