function ipv4ToInt(str) {
  const ipArr = str.split('.');
  let ret = ''
  for (let i = 0; i < ipArr.length; i++) {
    let val = ipArr[i].toString(2);
    ret += val;
  }
  return ret
}

console.log(ipv4ToInt('192.168.1.2'))