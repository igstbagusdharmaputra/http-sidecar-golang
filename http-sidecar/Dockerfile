FROM golang:latest as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]