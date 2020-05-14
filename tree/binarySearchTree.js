class Node {
  constructor(val) {
    this.val = val
    this.left = null
    this.right = null
  }
}

class BinarySearchTree {
  constructor() {
    this.root = null
  }
  insert(val) {
    const node = new Node(val)
    if (this.root === null) {
      this.root = node
    } else {
      BinarySearchTree.insertNode(this.root, node)
    }
  }
  static insertNode(node, newNode){
    if (node.val > newNode.val) {
      if (node.left === null) {
        node.left = newNode
      } else {
        BinarySearchTree.insertNode(node.left, newNode)
      }
    } else {
      if (node.right) {
        BinarySearchTree.insertNode(node.right, newNode)
      } else {
        node.right = newNode
      }
    }
  }
  remove(val) {
    this.root = BinarySearchTree.removeNode(this.root, val)
  }
  static removeNode(node, val) {
    if (node === null) {
      return false
    }
    if (node.val === val ) {
      if (node.left === null && node.right === null) {
        node = null // FIXME:
        return node
      } else if (node.left === null) {
        node = node.right
      } else if (node.right === null) {
        node = node.left
      } else if (node.left && node.right) {
        let rightMinNode = node.right
        while (rightMinNode.left) {
          rightMinNode = rightMinNode.left
        }
        node.val = rightMinNode.val
        node.right = BinarySearchTree.removeNode(node.right, rightMinNode.val)
      }
      return node
    } else if (node.val > val) {
      node.left = BinarySearchTree.removeNode(node.left, val)
      return node
    } else if (node.val < val) {
      node.right = BinarySearchTree.removeNode(node.right, val)
      return node
    }
  }
  search (val) {
    return BinarySearchTree.searchNode(this.root, val)
  }
  static searchNode(node, val) {
    if (node === null) {
      return false
    }
    if (node.val === val) {
      return true
    } else if (node.val > val) {
      BinarySearchTree.searchNode(node.left, val)
    } else if (node.val < val) {
      BinarySearchTree.searchNode(node.right, val)
    }
  }
  static printNode(val) {
    console.log(val)
  }
  inOrderTraverse(cb = BinarySearchTree.printNode) {
    BinarySearchTree.inOrderTraverseNode(this.root, cb)
  }
  preOrderTraverse(cb = BinarySearchTree.printNode) {
    BinarySearchTree.preOrderTraverseNode(this.root, cb)
  }
  postOrderTraverse(cb = BinarySearchTree.printNode) {
    BinarySearchTree.postOrderTraverseNode(this.root, cb)
  }
  levelOrderTraverse(cb = BinarySearchTree.printNode) {
    BinarySearchTree.levelOrderTraverseNode(this.root, cb);
  }
  separateByLevel(cb = BinarySearchTree.printNode) {
    BinarySearchTree.separateByLevelFn(this.root, cb);
  }
  static inOrderTraverseNode(node, cb) {
    if (node === null) {
      return
    }
    BinarySearchTree.inOrderTraverseNode(node.left, cb)
    cb(node.val)
    BinarySearchTree.inOrderTraverseNode(node.right, cb)
  }
  static preOrderTraverseNode(node, cb) {
    if (node === null) {
      return
    }
    cb(node.val)
    BinarySearchTree.preOrderTraverseNode(node.left, cb)
    BinarySearchTree.preOrderTraverseNode(node.right, cb)
  }
  static postOrderTraverseNode(node, cb) {
    if(node === null) {
      return
    }
    BinarySearchTree.postOrderTraverseNode(node.left, cb)
    BinarySearchTree.postOrderTraverseNode(node.right, cb)
    cb(node.val)
  }
  static levelOrderTraverseNode(node, cb) { // 层次遍历
    if (node === null) {
      return null
    }
    const list = [node]
    while (list.length) {
      node = list.shift()
      cb(node.val)
      if (node.left) {
       list.push(node.left)
      }
      if (node.right) {
        list.push(node.right)
      }
    }
  }
  static separateByLevelFn(node, cb, separator = '---*---') {
    const list = []
    const END_FLAG = 'END_FLAG'
    list.push(node)
    list.push(END_FLAG)
    while (list.length) {
      node = list.shift()
      if (node === END_FLAG && list.length > 0) {
        list.push(END_FLAG)
        cb(separator)
      } else {
        cb(node.val)
        if (node.left) {
          list.push(node.left)
        }
        if(node.right) {
          list.push(node.right)
        }
      }
    }
  }
  min() {
    let node = this.root
    if (node === null) {
      return null
    }
    while(node.left) {
      node = node.left
    }
    return node.val
  }
  max() {
    let node = this.root
    if (node === null) {
      return null
    }
    while(node.right) {
      node = node.right
    }
    return node.val
  }
}

const tree = new BinarySearchTree()

tree.insert(11);
tree.insert(10);
tree.insert(9);
tree.insert(8);
tree.insert(7);
tree.insert(6);
tree.insert(5);
tree.insert(4);
tree.insert(3);
tree.insert(2);
tree.insert(1);

// tree.inOrderTraverse();

// tree.levelOrderTraverse();

tree.separateByLevel();

// tree.remove(7)
// tree.separateByLevel();
// tree.inOrderTraverse();

// tree.preOrderTraverse();

// tree.postOrderTraverse(); 