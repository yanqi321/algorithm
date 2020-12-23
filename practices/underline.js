
function underline(data) {
  for (let key of Object.keys(data)) {
      let val = data[key];
      if (typeof(val) === 'object' || Array.isArray(val)) {
        val = underline(val)
      }
      let newKey = key.replace(/([A-Z]){1}/g, '_$1').toLowerCase()
      delete data[key];
      data[newKey] = val;
  }
  return data;
}
const data = { myAge: 1, mySex: 2, hostName: [{myP: 1}, {mySl: 2}]}
const ret = underline(data)
console.log(ret)