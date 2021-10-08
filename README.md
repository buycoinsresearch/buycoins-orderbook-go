## Introduction

A library for the Buycoins [orderbook](https://help.buycoins.africa/article/k0r0ayjfh7-order-books-explained) API.

API documentation- https://developers.buycoins.africa/orderbook-trading/introduction

---

## Installation

```bash
go get github.com/buycoinsresearch/buycoins-orderbook-go
```
---

## Authentication
To access the Buycoins API, you will need to generate your API credentials(a public key and a secret key) in the API Settings on the Buycoins app. At the moment, this is only open to fully verified Buycoins users. To be granted API access, please send an email to [support@buycoins.africa](mailto:support@buycoins.africa)

You should pass in your API keys as strings to the `Buycoins` function as shown here:

```go
package main

import (
	"github.com/buycoinsresearch/buycoins-orderbook-go"
)

var authorize = orderbooks.BuyCoins("public_key", "secret_key")
...
```

>**NOTE**<br/>
>Please ensure you pass in both the public key and the secret key
---

## Usage

### Stream Orderbook
Buycoins API documentation reference- https://developers.buycoins.africa/orderbook-trading/stream-order-book

The order book is a live computerized list of buy and sell orders organized by price level for a particular asset (in this context, cryptocurrencies). By connecting to the order book WebSockets API, you can get real-time updates on the Buycoins order book.

### Get pairs
Buycoins API documentation reference- https://developers.buycoins.africa/orderbook-trading/glossary#1-currency-pair-market

This returns an array of the currency pairs currently supported on Buycoins Pro.

Example:
```go
...

func main () {
    getPairs, err := authorize.GetPairs()
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s\n", getPairs)
}
```


### Get orders
Buycoins API documentation reference- https://developers.buycoins.africa/orderbook-trading/get-orders

The GetOrders function takes in three parameters:
- coin_pair: The [currency pair](https://developers.buycoins.africa/orderbook-trading/glossary#1-currency-pair-market) of the orders you'd like to see.
- order _status: Either `pending`, `in_progress`, `cancelled`, `partially_filled`, `successful`, or `failed`.
- side: The order side either `buy` or `sell`.

Example:

```go
...

func main () {
    getOrders, err := authorize.GetOrders("btc_usdt", "successful", "buy")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", getOrders)
}
```

This returns an array of all orders based on the parameters provided.
Example:
```json
{
  "data": {
    "getProOrders": {
      "edges": [
        {
          "node": {
            "id": "UHJvT3JkZXItOTU4ODQwNmMtZWM0Ny00ZjQ3LWEwMTItMDA3NTgzMTA3MTBi",
            "pair": "btc_ngnt",
            "price": "2000000",
            "side": "buy",
            "status": "pending",
            "timeInForce": "good_til_cancelled",
            "orderType": "limit_order",
            "fees": "0",
            "filled": "0",
            "total": "0.005",
            "initialBaseQuantity": "0.005",
            "initialQuoteQuantity": null,
            "remainingBaseQuantity": "0.005",
            "remainingQuoteQuantity": null,
            "meanExecutionPrice": null,
            "engineMessage": null
          }
        }
	{
          "node": {
            "id": "UHUhgTrEZXItOTU4ODQwNmMtZWM0Ny00ZjQ3LWEwMTItMDA3NTgzMTA3MTBi",
            "pair": "btc_ngnt",
            "price": "2000000",
            "side": "buy",
            "status": "pending",
            "timeInForce": "good_til_cancelled",
            "orderType": "limit_order",
            "fees": "0",
            "filled": "0",
            "total": "0.005",
            "initialBaseQuantity": "0.005",
            "initialQuoteQuantity": null,
            "remainingBaseQuantity": "0.005",
            "remainingQuoteQuantity": null,
            "meanExecutionPrice": null,
            "engineMessage": null
          }
        }
      ]
    }
  }
}
```


### Cancel order
Buycoins API documentation reference- https://developers.buycoins.africa/orderbook-trading/cancel-order

The CancelOrder function takes in the order ID of the order you wish to cancel and should return "order status": `cancelled`.

Example:

```go
...

func main () {
    cancelOrder, err := authorize.CancelOrder("UHJvT3JkZXItYzExYTY2ZmEtMzdiYy00ZDcyLWJmODgtNzg2NzRkNzdhZTlj")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", cancelOrder)
}
```

### Estimate order fees
Buycoins API documentation reference- https://developers.buycoins.africa/orderbook-trading/post-market-order#estimating-a-market-orders-fees

To get an estimate of how much placing a market order would cost:
```go
...

func main () {
    getProOrderFee,err := authorize.GetProOrderFees("market_order", "btc_usdt", "sell", 0.001)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", getProOrderFee)
}
```

### Post Market Order 
Buycoins API documentation reference- https://developers.buycoins.africa/orderbook-trading/post-market-order#placing-a-market-order

The function PostProMarketOrder takes in three parameters:
- The [currency pair](https://developers.buycoins.africa/orderbook-trading/glossary#1-currency-pair-market).
- The quantity of the currency you wish to buy or sell. 
- The order side: `buy` or `sell`.

Example:

```go
...

func main () {
    postProMarketOrder, err := authorize.PostProMarketOrder("btc_usdt", 5.00, "buy")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", postProMarketOrder)
}
```
### Post Limit Order 
Buycoins API reference- https://developers.buycoins.africa/orderbook-trading/post-limit-order

It is important to really [understand what limit  orders are](https://help.buycoins.africa/article/hwdaszt3ew-how-to-place-a-sell-limit-order), how they differ from market orders and the intricacies of posting limit orders to be able to make well informed decisions.
The function PostProLimitOrder takes in five parameters:
- The [currency pair](https://developers.buycoins.africa/orderbook-trading/glossary#1-currency-pair-market)
- Quantity of the currency you wish to buy or sell 
- The exact amount at which you want the order to be executed
- The order side: `buy` or `sell` 
- The [Time in force](https://developers.buycoins.africa/orderbook-trading/glossary#2-time-in-force): `fill_or_kill` or `good_til_cancelled`.

Example:

```go
...

func main () {
    postProLimitOrder, err := authorize.PostProLimitOrder("btc_usdt", 5.00, 26000000.00, "buy", "good_til_cancelled")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", postProLimitOrder)
}
```

### Get Deposit Link 
The function GetDepositLink takes in only one parameter:
- The amount you wish to deposit.

Example:

```go
...

func main () {
  getDepositLink, err := authorize.GetDepositLink(5000.00)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", getDepositLink)
```



>**NOTE**<br/>
>Please check the `example` directory to see a sample implementation.