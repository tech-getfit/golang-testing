FROM golang:1.17-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY ./ ./

RUN go build -o /product-go-micro

CMD [ "/product-go-micro" ]