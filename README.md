# go-websocket-testing
go websocket server testing

### 依赖项

- go
- nodejs
- pm2

### 开始

```bash
make build
pm2 start ./bin/*
cd frontend
npm install
pm2 start index.js --watch
sh bench.sh
```

测试程序退出后, 访问 http://localhost:8080

![bench](assets/bench.jpg)