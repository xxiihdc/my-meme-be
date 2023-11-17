# docker build -t golang-my-meme .
# docker run -p 8080:8080 -v /home/tran.van.hoai.duc/go/my-meme/my-meme-be:/app golang-my-meme


# Sử dụng hình ảnh golang:alpine làm base image
FROM golang:alpine AS builder

# Thiết lập biến môi trường
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Di chuyển vào thư mục làm việc của ứng dụng
WORKDIR /app

ADD . /app

# Sao chép go.mod và go.sum vào /app
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Sao chép mã nguồn của ứng dụng vào /app
COPY . .

# Build ứng dụng
RUN go build -o main .

# Tạo hình ảnh nhỏ gọn không chứa compiler và dependencies
FROM alpine:latest

# Di chuyển vào thư mục làm việc của ứng dụng
WORKDIR /app

# Sao chép binary từ builder stage
COPY --from=builder /app/main .

# Mở cổng 8080
EXPOSE 8080

RUN chmod +x main
# Chạy ứng dụng
CMD ["sh", "-c", "./main"]


