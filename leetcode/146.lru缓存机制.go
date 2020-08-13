/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (49.47%)
 * Likes:    718
 * Dislikes: 0
 * Total Accepted:    75.1K
 * Total Submissions: 151.3K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 * 
 * 获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 * 
 * 
 * 
 * 进阶:
 * 
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 * 
 * 
 * 
 * 示例:
 * 
 * LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得关键字 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得关键字 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 * 
 * 
 */

// @lc code=start
type LRUCache struct {
  cap int
  cache map[int]*DLinkedNode
  head,tail *DLinkedNode
}

type DLinkedNode struct {
  key, val int
  prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
  l := LRUCache{
    cap: capacity,
    head: &DLinkedNode{},
    tail: &DLinkedNode{},
    cache: map[int]*DLinkedNode{},
  }
  l.head.next = l.tail
  l.tail.prev = l.head
  return l
}


func (this *LRUCache) Get(key int) int {
  if node, ok := this.cache[key]; ok {
    this.remove(node)
    this.putHead(node)
    return node.val
  } else {
    return -1
  }
}


func (this *LRUCache) Put(key int, value int)  {
  if node, ok := this.cache[key]; ok {
    node.val = value
    this.remove(node)
    this.putHead(node)
  } else {
    if len(this.cache) >= this.cap {
      delKey := this.tail.prev.key
      this.remove(this.tail.prev)
      delete(this.cache, delKey)
    }
    newNode := DLinkedNode{ key: key, val: value }
    this.putHead(&newNode)
    this.cache[key] = &newNode
  }
}

func (this *LRUCache) putHead(node *DLinkedNode) {
  originNext := this.head.next
  this.head.next = node
  node.prev = this.head
  node.next = originNext
  originNext.prev = node
}


func (this *LRUCache) remove(node *DLinkedNode)  {
  node.prev.next = node.next
  node.next.prev = node.prev
}



/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

