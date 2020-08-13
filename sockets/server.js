const net = require('net');

let clientList = [];
const heartbeat = 'HEARTBEAT'; // 定义心跳包内容确保和平时发送的数据不会冲突

const server = net.createServer();
server.on('connection', (client) => {
  console.log('客户端建立连接:', client.remoteAddress + ':' + client.remotePort);
  clientList.push(client);
  client.on('data', (chunk) => {
    let content = chunk.toString();
    if (content === heartbeat) {
      console.log('收到客户端发过来的一个心跳包');
    } else {
      console.log('收到客户端发过来的数据:', content);
      client.write('服务端的数据:' + content);
    }
  });
  client.on('end', () => {
    console.log('收到客户端end');
    clientList.splice(clientList.indexOf(client), 1);
  });
  client.on('error', () => {
    clientList.splice(clientList.indexOf(client), 1);
  })
});
server.listen(9000);