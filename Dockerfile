# ใช้ Golang image สำหรับ build
FROM golang:1.25-alpine AS builder

WORKDIR /app

# install gcc + musl for cgo
RUN apk add --no-cache gcc musl-dev

# copy go mod
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# เปิด cgo
ENV CGO_ENABLED=1

# build binary
RUN go build -o main ./cmd/server

# ----------------------------

# ใช้ image เล็กสำหรับ run
FROM alpine:latest

WORKDIR /root/

# copy binary จาก builder
COPY --from=builder /app/main .

# expose port
EXPOSE 8080

# run app
CMD ["./main"]