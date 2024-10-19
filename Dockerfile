# Build stage
# golang:<version>-alpine は、Alpine Linux プロジェクトをベースにしている。
# イメージサイズを最小にするため、git、gcc、bash などは、Alpine-based のイメージには含まれていない。
FROM golang:1.23.2-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main /app/main.go

# Run stage
# Goで作成したバイナリは Alpine Linux 上で動く。
# alpineLinux とは軽量でセキュアな Linux であり、とにかく軽量。
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD [ "/app/main" ]