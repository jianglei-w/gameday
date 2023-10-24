FROM golang:1.21.1-alpine as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o gameday main.go

FROM busybox as runner
COPY --from=builder /app/gameday /app
COPY --from=builder /app/config.yaml /config.yaml
ENTRYPOINT ["/app"]