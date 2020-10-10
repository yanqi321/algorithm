class Stack {
  constructor () {
    this.dataStore = []
    this.top = 0
  }
  push(item) {
    this.dataStore.push(item)
    this.top++
  }
  pop(item) {
    this.top--
    return this.dataStore.pop(item)
  }
  peek(item) {
    return this.dataStore[this.top - 1]
  }
  clear() {
    this.dataStore = []
  }
  size() {
    return this.top
  }
}