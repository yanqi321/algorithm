function foo(c) {
  c = {ee: 22}
}

function test() {
  let obj = {c: {s: 1}, a: {d: 1}}
  foo(obj.c)
  console.log(obj)
}
test()