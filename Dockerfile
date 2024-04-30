FROM debian:stable-slim

WORKDIR /app

COPY ./server /app

CMD ["./server"]
