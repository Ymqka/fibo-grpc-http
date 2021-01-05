## HTTP REST
1) в корневой папке репозитория сбилдить и запустить контейнер
```
docker-compose build && docker-compose up
```

2) проверить, что все работает:
```
curl "http://localhost:10000/fibonacci?start=1&stop=3"
```

вывод:
```
[{"ID":1,"Number":1},{"ID":2,"Number":1},{"ID":3,"Number":2}]
```

ограничения:
выставлено ограничение на стоп в 111111, чтобы обойти ограничение нужно добавить флажок force=1
```
curl "http://localhost:10000/fibonacci?start=111112&stop=111112&force=1"
```

## gRPC client
1) Нужно поднять контейнер командой из 1 пункта HTTP REST
2) запустить клиент из корневой папки репозитория
```
go run cmd/grpc-client/grpc-client.go
```

вывод
```
2021/01/05 21:23:41 [{1 1} {2 1} {3 2} {4 3} {5 5} {6 8} {7 13} {8 21} {9 34} {10 55}]
```

## Perfomance
на "холодный" кэш ~18 секунд
```
$ time curl "http://localhost:10000/fibonacci?start=200000&stop=200000&force=1" > /dev/null 2>&1
curl "http://localhost:10000/fibonacci?start=200000&stop=200000&force=1" >  2  0.00s user 0.00s system 0% cpu 17.709 total
```

на "горячий" кэш около 20 миллисекунды
```
$ time curl "http://localhost:10000/fibonacci?start=200000&stop=200000&force=1" > /dev/null 2>&1
curl "http://localhost:10000/fibonacci?start=200000&stop=200000&force=1" >  2  0.00s user 0.01s system 72% cpu 0.017 total
```

100 000 первых чисел фибоначчи в redis'e съедают
```
> info memory
# Memory
used_memory:483677752
used_memory_human:461.27M
```