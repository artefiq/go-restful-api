# Menggunakan image Go resmi
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy semua file proyek ke dalam container
COPY . .

# Download dependency
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["/app/main"]
