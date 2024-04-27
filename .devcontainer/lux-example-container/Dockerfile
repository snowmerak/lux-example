FROM golang:1.22-bookworm

RUN go install github.com/snowmerak/lux/v3@latest

CMD ["/bin/sh", "-c", "while true; do sleep 30; done;"]
