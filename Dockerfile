FROM golang:1.21 as builder
RUN mkdir movie-crud
COPY . /movie-crud
WORKDIR /movie-crud
RUN go mod tidy
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 8080