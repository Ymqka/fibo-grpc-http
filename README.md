## HTTP REST
1) в корневой папке репозитория сбилдить и запустить контейнер
```
$ docker-compose build && docker-compose up
```

2) проверить, что все работает:
```
$ curl "http://localhost:10000/fibonacci?start=1&stop=3"
```

вывод:
```
[0 1 1]
```

## gRPC client
1) Нужно поднять контейнер командой из 1 пункта HTTP REST
2) запустить клиент из корневой папки репозитория
```
$ go run cmd/grpc-client/grpc-client.go
```

вывод
```
2020/12/25 00:08:22 [0 1 1 2 3 5 8 13 21 34]
```