FROM golang:1.19 as builder

WORKDIR /workspace
#COPY go.mod go.mod
#COPY go.sum go.sum
#COPY main.go main.go
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN   go mod tidy

RUN CGO_ENABLED=0  go build -o server .

FROM alpine

EXPOSE 8080
COPY --from=builder /workspace/server /server
CMD ["/server"]