class Node {
  constructor (el) {
    this.el = el
    this.next = null
  }
}

class LinkedList {
  constructor (val) {
    this.head = null
    this.length  = 0
    if (val) {
      this.head = new Node(val)
      this.length = 1
    }
  }
  append(el) {
    const node = new Node(el)
    if (this.head === null) {
      this.head = node
    }  else {
      let current = this.head
      while(current.next) {
        current = current.next
      }
      current.next = node
    }
    this.length++
  }
  removeAt(pos) {
    if (pos >= this.length || pos < 0) {
      return null
    }
    let current = this.head
    if (pos === 0) {
      this.head = current.next
    } else {
      let index = 0;
      let prev = null;
      while (index < pos) {
        prev = current
        current = current.next
        index++
      }
      prev.next = current.next
    }
    this.length--
    return current.val
  }
  insert(pos, val) {
    if (pos >= this.length || pos <0) {
      return false
    }
    const node =new Node(val)
    if (pos === 0) {
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
      node.next = current
      prev.next = node
    }
    this.length++
    return true
  }
  indexOf (val, start = 0) {
    if (start >= this.length) {
      return -1
    }
    let index = 0
    let current = this.head
    while (index < this.length) {
      if (current.val = val && index >= start) {
        return index
      }
      current = current.next;
      index++
    }
    return -1 
  }
  remove (val, start = 0) {
    if (start >= this.length) {
      return false
    }
    let current = this.head
    let prev = null
    while(index < this.length) {
      if (current.val = val && index >= start) {
        if (index === 0) {
          this.head = current.next;
        } else {
          prev.next = current.next
        }
        this.length--
        return true
      }
      prev = current
      current = current.next;
      index++
    }
  }
  size() {
    return this.length
  }
  isEmpty() {
    return !!this.length
  }
}