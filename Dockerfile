FROM golang:1.14
RUN mkdir /app
RUN go get -u github.com/gorilla/mux
ADD main.go /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]