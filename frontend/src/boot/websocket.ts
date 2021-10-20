const url = 'ws://localhost:8888/ws';
const ws = new WebSocket(url);

ws.onmessage = (msg) => {
  console.log(JSON.parse(msg.data));
};
