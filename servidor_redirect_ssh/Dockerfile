#go mod init ssh-redirect
# Build
#docker build -t ssh-redirect .
#docker run -d --name ssh-redirect  --restart=unless-stopped -p 8081:8080 ssh-redirect

# Dockerfile
FROM golang:1.25.2-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o ssh-redirect .

# Fase final: imagem leve
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/ssh-redirect .
EXPOSE 8080
CMD ["./ssh-redirect"]