FROM golang:1.17
COPY ./go-app /app/go
WORKDIR /app/go
RUN go build -o main .
CMD ["/app/go/main"]