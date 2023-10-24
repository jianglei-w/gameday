# GameDay
GameDay是一个....
## 快速开始

```shell
go mod download # 或者 go mod tidy
go build && ./main.go        # 或者直接go run main.go
```


## Docker
```shell
docker build -t gameday:v1 .
docker run -d --name gameday -v ./config.yaml:/config.yaml gameday:v1
```