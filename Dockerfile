# https://docs.docker.com/language/golang/build-images/

FROM golang:1.20.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY controllers/*.go ./controllers/

RUN go build -o /bitcoin-tracker-proxy

EXPOSE 3000

CMD [ "/bitcoin-tracker-proxy" ]
