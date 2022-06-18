FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

RUN go build -o /go-app

EXPOSE 8080

CMD [ "/go-app" ]
