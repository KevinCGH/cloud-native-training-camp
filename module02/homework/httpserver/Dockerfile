FROM golang:1.17 AS build

RUN mkdir /app
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
RUN go mod download && GOOS=linux go build -installsuffix cgo -o httpserver ./cmd/main.go

###
FROM scratch as final
COPY --from=build /app/httpserver .
EXPOSE 80
ENTRYPOINT ["/httpserver"]