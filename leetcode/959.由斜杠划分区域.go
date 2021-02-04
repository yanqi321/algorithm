/*
 * @lc app=leetcode.cn id=959 lang=golang
 *
 * [959] 由斜杠划分区域 
 *	--------
 *	|
 */

// @lc code=start
func regionsBySlashes(grid []string) int {
	gLen := len(grid)
	uSet := NewSet(gLen * gLen * 4)
	for i := 0; i < gLen; i++ {
		for j := 0; j < gLen; j++ {
			idx := (i * gLen + j) * 4
			switch grid[i][j] {
			case '/':
				uSet.Union(idx, idx + 3)
				uSet.Union(idx+ 1,idx + 2)
			case '\\':
				uSet.Union(idx,idx + 1)
				uSet.Union(idx + 2, idx + 3)
			case ' ':
				uSet.Union(idx,idx + 1)
				uSet.Union(idx + 1,idx + 2)
				uSet.Union(idx + 2,idx + 3)
			}
			if j + 1< gLen {
				uSet.Union(idx + 1, idx + 7)
			}
			if i + 1 < gLen {
				uSet.Union(idx + 2, idx + 4 * gLen)
			}
		}
	}
	return uSet.Count
}

type UnionFindSet struct {
	Parents []int
	Count int
}

func NewSet(size int) *UnionFindSet {
	s := make([]int, size)
	for i :=0; i < len(s); i++ {
		s[i] = i
	}
	return &UnionFindSet {
		Parents: s,
		Count: size,
	}
}

func (u *UnionFindSet) Union(x, y int) {
	rootX := u.Find(x)
	rootY := u.Find(y)
	if rootX == rootY {
		return
	} else if rootX <= rootY {
		u.Parents[rootY] = rootX
	} else {
		u.Parents[rootX] = rootY
	}

	u.Count--
}

func (u *UnionFindSet) Find(node int) int {
		if u.Parents[node] == node {
			return node
		}
		root := u.Find(u.Parents[node])
		u.Parents[node] = root
		return root
}
// @lc code=end

