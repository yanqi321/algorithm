class Node {
  constructor(val) {
    this.val = val
    this.left = null
    this.right = null
    this.height = 0
  }
}

const BalanceFactor = {
  UNBALANCED_RIGHT: -2,
  SLIGHTLY_UNBALANCED_RIGHT: -1,
  BALANCED: 0,
  SLIGHTLY_UNBALANCED_LEFT: 1,
  UNBALANCED_LEFT: 2
};
class AVLTree {
  constructor() {
    this.root = null
  }
  insert (val) {
    const newNode = new Node(val)
    if (this.root === null) {
      newNode.height = 0
      this.root = newNode
    } else {
      this.root = AVLTree.insertNode(this.root, newNode)
    }
  }
  static insertNode (node, newNode) {
    if (node.val > newNode.val) {
      if (node.left === null) {
        node.left = newNode
      } else {
        node.left = AVLTree.insertNode(node.left, newNode)
      }
    } else {
      if (node.right === null) {
        node.right = newNode
      } else {
        node.right = AVLTree.insertNode(node.right, newNode)
      }
    }
    const balanceFactor = AVLTree.getBalanceFactor(node)
    if (balanceFactor === BalanceFactor.UNBALANCED_LEFT) {
      if (node.left.val > newNode.val) {
        node = AVLTree.rotateLL(node)
      } else {
        node = AVLTree.rotateLR(node)
      }
    } else if (balanceFactor === BalanceFactor.UNBALANCED_RIGHT) {
      if (node.right.val < newNode.val) {
        node = AVLTree.rotateRR(node)
      } else {
        node = AVLTree.rotateRL(node)
      }
    }
    node.height = Math.max(AVLTree.getNodeHeight(node.left), AVLTree.getNodeHeight(node.right)) + 1
    return node
  }
  remove(val) {
    this.root = AVLTree.removeNode(this.root, val)
  }
  static removeNode(node, val) {
    if (node === null) {
      return false
    }
    if (node.val === val ) {
      if (node.left === null && node.right === null) {
        node = null 
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
        node.right = AVLTree.removeNode(node.right, rightMinNode.val)
      }
    } else if (node.val > val) {
      node.left = AVLTree.removeNode(node.left, val)
    } else if (node.val < val) {
      node.right = AVLTree.removeNode(node.right, val)
    }
    const balanceFactor = AVLTree.getBalanceFactor(node)
    if (balanceFactor === BalanceFactor.UNBALANCED_LEFT) {
      if (node.left.val > newNode.val) {
        node = AVLTree.rotateLL(node)
      } else {
        node = AVLTree.rotateLR(node)
      }
    } else if (balanceFactor === BalanceFactor.UNBALANCED_RIGHT) {
      if (node.right.val < newNode.val) {
        node = AVLTree.rotateRR(node)
      } else {
        node = AVLTree.rotateRL(node)
      }
    }
    if (node) {
      node.height = Math.max(AVLTree.getNodeHeight(node.left), AVLTree.getNodeHeight(node.right)) + 1
    }
    return node
  }
  static getNodeHeight (node) {
    if (node === null) {
      return -1
    }
    return node.height
  }
  static getBalanceFactor (node) {
    const heightDiff = AVLTree.getNodeHeight(node.left) - AVLTree.getNodeHeight(node.right);
    switch (heightDiff) {
      case -2:
        return BalanceFactor.UNBALANCED_RIGHT
      case -1:
        return BalanceFactor.SLIGHTLY_UNBALANCED_RIGHT
      case 0:
        return BalanceFactor.BALANCED
      case 1: 
        return BalanceFactor.SLIGHTLY_UNBALANCED_LEFT
      case 2:
        return BalanceFactor.UNBALANCED_LEFT
      default:
        return 'UNKNOWN'
    }
  }
  static rotateLL (node) {
    const tmp = node.left
    node.left = tmp.right
    tmp.right = node
    node.height = Math.max(AVLTree.getNodeHeight(node.left), AVLTree.getNodeHeight(node.right)) + 1
    tmp.height = Math.max(AVLTree.getNodeHeight(tmp.left), AVLTree.getNodeHeight(tmp.right)) + 1
    return tmp

  }
  static rotateRR (node) {
    const tmp = node.right
    node.right = tmp.left
    tmp.left = node
    node.height = Math.max(AVLTree.getNodeHeight(node.left), AVLTree.getNodeHeight(node.right)) + 1
    tmp.height = Math.max(AVLTree.getNodeHeight(tmp.left), AVLTree.getNodeHeight(tmp.right)) + 1
    return tmp
  }
  static rotateLR (node) {
    node.left = AVLTree.rotateRR(node.left)
    return AVLTree.rotateLL(node)
  }
  static rotateRL (node) {
    node.right = AVLTree.rotateLL(node.right)
    return AVLTree.rotateRR(node)
  }
  search (val) {
    return AVLTree.searchNode(this.root, val)
  }
  static searchNode(node, val) {
    if (node === null) {
      return false
    }
    if (node.val === val) {
      return true
    } else if (node.val > val) {
      return AVLTree.searchNode(node.left, val)
    } else if (node.val < val) {
      return AVLTree.searchNode(node.right, val)
    }
  }
  static printNode(val) {
    console.log(val)
  }
  inOrderTraverse(cb = AVLTree.printNode) {
    AVLTree.inOrderTraverseNode(this.root, cb)
  }
  preOrderTraverse(cb = AVLTree.printNode) {
    AVLTree.preOrderTraverseNode(this.root, cb)
  }
  postOrderTraverse(cb = AVLTree.printNode) {
    AVLTree.postOrderTraverseNode(this.root, cb)
  }
  levelOrderTraverse(cb = AVLTree.printNode) {
    AVLTree.levelOrderTraverseNode(this.root, cb);
  }
  separateByLevel(cb = AVLTree.printNode) {
    AVLTree.separateByLevelFn(this.root, cb);
  }
  static inOrderTraverseNode(node, cb) {
    if (node === null) {
      return
    }
    AVLTree.inOrderTraverseNode(node.left, cb)
    cb(node.val)
    AVLTree.inOrderTraverseNode(node.right, cb)
  }
  static preOrderTraverseNode(node, cb) {
    if (node === null) {
      return
    }
    cb(node.val)
    AVLTree.preOrderTraverseNode(node.left, cb)
    AVLTree.preOrderTraverseNode(node.right, cb)
  }
  static postOrderTraverseNode(node, cb) {
    if(node === null) {
      return
    }
    AVLTree.postOrderTraverseNode(node.left, cb)
    AVLTree.postOrderTraverseNode(node.right, cb)
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


const tree = new AVLTree()


tree.insert(11);
tree.insert(3);
tree.insert(5);
tree.insert(6);
tree.insert(7);
tree.insert(9);
tree.insert(8);
tree.insert(10);
tree.insert(13);
tree.insert(12);
tree.insert(14);
tree.insert(15);
tree.insert(20);
tree.insert(18);
tree.insert(25);

// tree.inOrderTraverse();

// tree.levelOrderTraverse();

// tree.separateByLevel();
tree.search(25)

console.log('\n')
// tree.remove(7)
// tree.separateByLevel();
// tree.inOrderTraverse();

console.log('\n')
// tree.preOrderTraverse();

console.log('\n')
// tree.postOrderTraverse(); 