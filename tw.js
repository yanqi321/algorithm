'use strict';
/**
 * 开闭原则
 * 最少知识原则
 */

const HotelMap = new Map().set('JW', {
  brand: 'Golden',
  level: 5,
  totalRoom: 400,
  remainRoom: 400,
}).set('Golden', {
  brand: 'Golden',
  level: 4,
  totalRoom: 400,
  restRoom: 400,
}).set(
  'Deltas', {
  brand: 'Deltas',
  level: 3,
  totalRoom: 400,
  remainRoom: 400,
}).set(
  'PeaceExpress', {
  brand: 'PeaceExpress',
  level: 3,
  totalRoom: 400,
  remainRoom: 400,
})
// eg: CNN JW 2018-4-1 2018-5-1 10 TODO: 可变参数处理
function booking(str) {
  if (!str) return 'Invalid'
  const cmds = str.split('\n', str);
  let err = '';
  for (let cmd of cmds) {
    bookArr = cmds.split(' ');
    const hotelType = bookArr[1]
    const checkinDate = bookArr[2]
    const checkoutDate = bookArr[3]
    const roomCount = bookArr[4]
    const info =  HotelMap.get(hotelType);
    if (info.restRoom < roomCount) {
      err = 'Invalid';
      break;
    }
    
  }
  return err
}

function validDate(checkinDate, checkoutDate) {
  if (checkinDate > checkoutDate || new Date(checkinDate + '00:00:00').getTime() < Date.now()) { // FIXME:
    return false
  }
}

// TODO: 日期 和 房间数量耦合 如何处理
function valideRoom() {

}

function handleBook() {

}