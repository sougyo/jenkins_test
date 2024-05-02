FROM debian:stable-slim

WORKDIR /app

COPY ./server index.html /app

CMD ["./server"]
