# Step 1: ใช้ Go version 1.24 หรือสูงกว่า
FROM golang:1.24-alpine

# Step 2: ตั้ง working directory เป็น /app
WORKDIR /app

# Step 3: ติดตั้ง Air สำหรับ live reload
RUN go install github.com/air-verse/air@latest

# Step 4: คัดลอก go.mod และ go.sum
COPY go.mod go.sum ./

# Step 5: ดาวน์โหลด dependencies
RUN go mod download

# Step 6: คัดลอก source code
COPY . .

# Step 7: เปิดพอร์ต 8080
EXPOSE 8080