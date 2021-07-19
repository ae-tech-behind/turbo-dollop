FROM golang:1.16-alpine AS build

ADD . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

RUN go install

RUN go build -o /turbo-dollop

CMD [ "/turbo-dollop" ]