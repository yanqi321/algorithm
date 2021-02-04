/*
 * @lc app=leetcode.cn id=1631 lang=golang
 *
 * [1631] 最小体力消耗路径
 * dp[0][0] 
 */

// @lc code=start
func minimumEffortPath(heights [][]int) int {
    rLen, cLen := len(heights), len(heights[0])
    edgeList := make([]edge, 0)
    for i:=0; i<rLen; i++ {
        for j := 0; j < cLen; j++ {
            start := i *cLen + j
            if i > 0 {
                edgeList = append(edgeList, edge{ 
                    start: start, 
                    end: start - cLen, 
                    diff: abs(heights[i][j] - heights[i - 1][j]),
                    })
            }
            if j > 0 {
                edgeList = append(edgeList, edge{ 
                    start: start, 
                    end: start - 1, 
                    diff: abs(heights[i][j] - heights[i][j - 1]),
                    })
            }
        }
    }
    sort.Slice(edgeList, func(i,j int) bool { return edgeList[i].diff < edgeList[j].diff })
    uf := New(rLen * cLen)
    for _, v := range edgeList {
        uf.Union(v.start, v.end)
        if uf.Find(0) == uf.Find(rLen * cLen - 1) { // 从最小高度差开始尝试
            return v.diff // 题目是求最值而不是路径和
        }
    }
    return 0
}

type edge struct {
    start int // 边起始
    end int // 边结束
    diff int // 高度差
}

type UnionFind struct {
    parent []int
}

func New(n int) *UnionFind {
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    return &UnionFind{
        parent,
    }
}

func (uf *UnionFind) Union(x, y int) {
    rootX := uf.Find(x)
    rootY := uf.Find(y)
    if rootX == rootY {
        return
    }
    if rootX < rootY {
        uf.parent[rootX] = uf.parent[rootY]
    } else {
        uf.parent[rootY] = uf.parent[rootX]
    }
} 

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] == x {
        return x
    }
    root := uf.Find(uf.parent[x])
    uf.parent[x] = root
    return root
}

func abs(x int) int {
    if x < 0 {
        x = -x
    }
    return x
}
// @lc code=end

