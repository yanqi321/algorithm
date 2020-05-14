class Node {
  constructor(val) {
    this.val = val
    this.next = null
  }
}

class DoublyNode extends Node {
  constructor(val) {
    super(val)
    this.prev = null
  }
}

class DoublyLinkedList {
  constructor(val) {
    this.head = null
    this.length  = 0
    if (val) {
      this.head = new DoublyNode(val)
      this.length = 1
    }
  }
  append(val) {
    const node = new DoublyNode(val)
    if (this.head === null) {
      this.head = node
    } else {
      let current  = this.head
      while(current.next) {
        current = current.next
      }
      current.next = node
      node.prev = current
    }
    this.length++
  }
  removeAt(pos) {
    if (pos >= this.length || pos < 0) {
      return false
    }
    let current = this.head
    if (pos === 0) {
      this.head = current.next
      this.head.prev = null
      return true
    }
    let index = 0
    let prev = null
    while (index < pos) {
      prev = current
      current = current.next
    }
    prev.next = current.next
    current.next.prev = prev
    this.length--
  }
  insert(pos, val) {
    if (pos >= this.length || pos < 0) {
      return false
    }
    const node =new Node(val)
    if (pos === 0) {
      this.head.prev = node
      node.next = this.head
      this.head = node
    } else {
      let index = 0
      let current = this.head
      let prev = null
      while (index < pos) {
        prev = current
        current = current.next
      }
      current.prev = node
      node.next = current
      node.prev = prev
      prev.next = node

    }
    this.length++
    return true
  }
}