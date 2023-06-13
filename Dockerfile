FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go build -o simple-server .

FROM debian:12-slim
COPY --from=builder /app/simple-server .
EXPOSE 8080
CMD [ "./simple-server" ]

