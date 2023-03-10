# Bitcoin Tracker Proxy

Bitcoin Tracker Proxy is a Gin (go) proxy to get the prices from Binance and Crypto Compare.

NodeJs version: [alejandrosnz/bitcoin-tracker-proxy](https://github.com/alejandrosnz/bitcoin-tracker-proxy)

## Installation

Install Go [here](https://go.dev/doc/install)

Install the dependencies and start the server

```bash
go install .
go run .\main.go
```

Alternatively, the following docker image can be used: [alejandrosnz/bitcoin-tracker-proxy-go](https://hub.docker.com/r/alejandrosnz/bitcoin-tracker-proxy-go)

```bash
docker pull alejandrosnz/bitcoin-tracker-proxy-go
docker run --publish 3000:3000 alejandrosnz/bitcoin-tracker-proxy-go
```

## Usage

```
GET /api/ticker/current_price/:symbol
GET /api/ticker/current_price/BTC

{
  "currentPrice": 19919.8
}


GET /api/ticker/closing_price/:symbol
GET /api/ticker/closing_price/BTC

{
  "closingPrice": 20369.37
}

```