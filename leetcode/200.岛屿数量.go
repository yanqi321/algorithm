import "fmt"

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	u := &UnionFindSet{}
	row := len(grid)
	col := len(grid[0])
	u.Init(grid)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				if i-1 >= 0 && grid[i-1][j] == '1' {
					u.Union(i*col+j, (i-1)*col+j)
				}
				if i+1 < row && grid[i+1][j] == '1' {
					u.Union(i*col+j, (i+1)*col+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					u.Union(i*col+j, i*col+j-1)
				}
				if j+1 < col && grid[i][j+1] == '1' {
					u.Union(i*col+j, i*col+j+1)
				}
				grid[i][j] = '0' // why ?
			}
		}
	}
	return u.SetCount
}

type UnionFindSet struct {
	Parents  []int
	SetCount int
}

func (u *UnionFindSet) Init(grid [][]byte) {
	row := len(grid)
	col := len(grid[0])
	u.Parents = make([]int, row*col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			key := i* col + j
			u.Parents[key] = key
			if grid[i][j] == '1' {
				u.SetCount++
			}
		}
  }
}

func (u *UnionFindSet) Union(x, y int) {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return
	}
	if rootX < rootY {
		u.Parents[rootX] = rootY
	} else {
		u.Parents[rootY] = rootX
	}
	u.SetCount--
}

func (u *UnionFindSet) Find(node int) int {
	if u.Parents[node] == node {
		return node
	}
	root := u.Find(u.Parents[node])
	u.Parents[node] = root // 优化层级
	return root
}

// @lc code=end

