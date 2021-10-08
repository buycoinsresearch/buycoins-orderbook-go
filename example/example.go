package main

import (
	"fmt"
	"github.com/buycoinsresearch/buycoins-orderbook-go"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//alternatively, you could pass in your API credentials into a global variable like so;
//var authorize = orderbooks.Buycoins(blahburibdblahbeubblah, blahblahblah)
//this eliminates the need to call godotenv.Load in the main function and  
//would permit a global use of the orderBooks.configCredentials



func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	authorize := orderbooks.BuyCoins(os.Getenv("PUBLIC_KEY"), os.Getenv("SECRET_KEY"))

	getPairs, err := authorize.GetPairs()
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s\n", getPairs)

	getOrders, err := authorize.GetOrders("btc_ngnt", "pending", "buy")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", getOrders)

	cancelOrder, err := authorize.CancelOrder("UHJvT3JkZXItZDRmYjMyZDYtOGZjMy00ZTJlLWEzNGYtNTQ3YmEwMzcxMDQ0")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", cancelOrder)

	getProOrderFee,err := authorize.GetProOrderFees("market_order", "btc_usdt", "sell", 0.001)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", getProOrderFee)

	postProMarketOrder, err := authorize.PostProMarketOrder("btc_usdt", 5.00, "buy")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", postProMarketOrder)

	postProLimitOrder, err := authorize.PostProLimitOrder("btc_ngnt", 0.005, 2000000.00, "buy", "good_til_cancelled")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", postProLimitOrder)

	getDepositLink, err := authorize.GetDepositLink(5000.00)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", getDepositLink)
}