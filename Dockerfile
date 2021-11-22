FROM golang:1.17.3 as builder
ENV GOPROXY="https://goproxy.io"
ENV GOPATH=$HOME/go
ENV GOBIN=$HOME/go/bin
ENV PATH=$PATH:$GOPATH/bin
COPY . /app/
# 下载指定的包，go.mod已经记录，可以直接使用
RUN cd /app && go build -o gitbook-generator .

FROM debian:10.11
COPY --from=builder /app/gitbook-generator /app/gitbook-generator
ENTRYPOINT ["/app/gitbook-generator"]
