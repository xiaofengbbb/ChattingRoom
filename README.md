## Go语言实现的基于TCP Socket的并发聊天室Demo

### Usage

#### Server
```bash
cd src/chat
go build
./chat
```

#### Clients
```bash
cd src/chat
telnet localhost 8888
/nick [id]        # 给当前client起个名字
/join [roomName]  # 加入roomName，若不存在则创建一个新room
/quit             # 退出当前房间
/msg [msg]        # 发送消息
/rooms            # 查询所有聊天室
```


