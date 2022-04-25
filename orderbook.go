package orderbooks

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

const (
	prodEndpoint    = "https://backend.buycoins.tech/api/graphql"
	stagingEndpoint = "https://bitkoin-dev.herokuapp.com/api/graphql"
)

var endpoint string

type ConfigCredentials struct {
	basicAuth string
}

func Buycoins(publicKey, secretKey string) ConfigCredentials {
	env := os.Getenv("APP_ENV")
	fmt.Println("Env: ", env)
	switch env {
	case "STAGING":
		endpoint = stagingEndpoint
	case  "test":
		endpoint = stagingEndpoint
	case "PRODUCTION":
		endpoint = prodEndpoint
	}
	log.Printf("End point %v", endpoint)
	auth := "Basic " + b64.URLEncoding.EncodeToString([]byte(publicKey+":"+secretKey))
	return ConfigCredentials{
		basicAuth: auth,
	}
}

func (config ConfigCredentials) GetPairs() ([]byte, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(`
		query {
			getPairs
		}
	`)

	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()
	res := struct {
		GetPairs []string
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}

	pairs, err := json.MarshalIndent(res.GetPairs, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	return pairs, nil
}

func (config ConfigCredentials) GetOrders(coinPair, status, side string) (getProOrders, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(getProOrdersQuery)
	req.Var("pair_", coinPair)
	req.Var("status_", status)
	req.Var("side_", side)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()
	res := struct {
		GetProOrders struct {
			Edges []struct {
				Node struct {
					Id                     string
					Pair                   string
					Price                  string
					Side                   string
					Status                 string
					TimeInForce            string
					OrderType              string
					Fees                   string
					Filled                 string
					Total                  string
					InitialBaseQuantity    string
					InitialQuoteQuantity   string
					RemainingBaseQuantity  string
					RemainingQuoteQuantity string
					MeanExecutionPrice     string
					EngineMessage          string
				}
			}
		}
	}{}
	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Println(err)
		return getProOrders{}, err
	}
	fmt.Println("Successfully connected")

	return getProOrders{
		Edges: res.GetProOrders.Edges,
	}, nil
}

func (config ConfigCredentials) CancelOrder(id string) (cancelOrder, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(cancelOrderQuery)

	req.Var("id", id)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()

	res := struct {
		CancelOrder struct {
			Id                     string
			Pair                   string
			Price                  string
			Side                   string
			Status                 string
			TimeInForce            string
			OrderType              string
			Fees                   string
			Filled                 string
			Total                  string
			InitialBaseQuantity    string
			InitialQuoteQuantity   string
			RemainingBaseQuantity  string
			RemainingQuoteQuantity string
			MeanExecutionPrice     string
			EngineMessage          string
		}
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Println(err)
		return cancelOrder{}, err
	}

	return cancelOrder{
		Id:                     res.CancelOrder.Id,
		pair:                   res.CancelOrder.Pair,
		price:                  res.CancelOrder.Price,
		side:                   res.CancelOrder.Side,
		status:                 res.CancelOrder.Status,
		timeInForce:            res.CancelOrder.TimeInForce,
		orderType:              res.CancelOrder.OrderType,
		fees:                   res.CancelOrder.Fees,
		filled:                 res.CancelOrder.Filled,
		total:                  res.CancelOrder.Total,
		initialBaseQuantity:    res.CancelOrder.InitialBaseQuantity,
		initialQuoteQuantity:   res.CancelOrder.InitialQuoteQuantity,
		remainingBaseQuantity:  res.CancelOrder.RemainingBaseQuantity,
		remainingQuoteQuantity: res.CancelOrder.RemainingQuoteQuantity,
		meanExecutionPrice:     res.CancelOrder.MeanExecutionPrice,
		engineMessage:          res.CancelOrder.EngineMessage,
	}, nil

}

func (config ConfigCredentials) GetProOrderFees(orderType string, pair string, side string, amount float64) (getProOrderFees, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(getProOrderFeesQuery)

	req.Var("orderType_", orderType)
	req.Var("pair_", pair)
	req.Var("side_", side)
	req.Var("amount_", amount)
	ctx := context.Background()

	res := struct {
		GetProOrderFees struct {
			Fees               string
			BaseCurrencyTotal  string
			QuoteCurrencyTotal string
			Price              string
		}
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Println(err)
		return getProOrderFees{}, err
	}

	return getProOrderFees{
		fee:                res.GetProOrderFees.Fees,
		baseCurrencyTotal:  res.GetProOrderFees.BaseCurrencyTotal,
		quoteCurrencyTotal: res.GetProOrderFees.QuoteCurrencyTotal,
		price:              res.GetProOrderFees.Price,
	}, nil
}

func (config ConfigCredentials) PostProMarketOrder(pair string, quantity float64, side string) (postProMarketOrder, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(postProMarketOrderQuery)

	req.Var("pair_", pair)
	req.Var("quantity_", quantity)
	req.Var("side_", side)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()

	res := struct {
		PostProMarketOrder struct {
			Id                     string
			Pair                   string
			Price                  string
			Side                   string
			Status                 string
			TimeInForce            string
			OrderType              string
			Fees                   string
			Filled                 string
			Total                  string
			InitialBaseQuantity    string
			InitialQuoteQuantity   string
			RemainingBaseQuantity  string
			RemainingQuoteQuantity string
			MeanExecutionPrice     string
			EngineMessage          string
		}
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Println(err)
		return postProMarketOrder{}, err
	}

	return postProMarketOrder{
		Id:                     res.PostProMarketOrder.Id,
		pair:                   res.PostProMarketOrder.Pair,
		price:                  res.PostProMarketOrder.Price,
		side:                   res.PostProMarketOrder.Side,
		status:                 res.PostProMarketOrder.Status,
		timeInForce:            res.PostProMarketOrder.TimeInForce,
		orderType:              res.PostProMarketOrder.OrderType,
		fees:                   res.PostProMarketOrder.Fees,
		filled:                 res.PostProMarketOrder.Filled,
		total:                  res.PostProMarketOrder.Total,
		initialBaseQuantity:    res.PostProMarketOrder.InitialBaseQuantity,
		initialQuoteQuantity:   res.PostProMarketOrder.InitialQuoteQuantity,
		remainingBaseQuantity:  res.PostProMarketOrder.RemainingBaseQuantity,
		remainingQuoteQuantity: res.PostProMarketOrder.RemainingQuoteQuantity,
		meanExecutionPrice:     res.PostProMarketOrder.MeanExecutionPrice,
		engineMessage:          res.PostProMarketOrder.EngineMessage,
	}, nil
}

func (config ConfigCredentials) PostProLimitOrder(pair string, quantity float64, price float64, side string, timeInForce string) (LimitOrder, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(postProLimitOrderQuery)

	req.Var("pair_", pair)
	req.Var("quantity_", quantity)
	req.Var("price_", price)
	req.Var("side_", side)
	req.Var("timeInForce_", timeInForce)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()

	res := struct {
		PostProLimitOrder struct {
			Id                     string
			Pair                   string
			Price                  string
			Side                   string
			Status                 string
			TimeInForce            string
			OrderType              string
			Fees                   string
			Filled                 string
			Total                  string
			InitialBaseQuantity    string
			InitialQuoteQuantity   string
			RemainingBaseQuantity  string
			RemainingQuoteQuantity string
			MeanExecutionPrice     string
			EngineMessage          string
		}
	}{}

	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Println(err)
		return LimitOrder{}, err
	}

	return LimitOrder{
		Id:                     res.PostProLimitOrder.Id,
		pair:                   res.PostProLimitOrder.Pair,
		price:                  res.PostProLimitOrder.Price,
		side:                   res.PostProLimitOrder.Side,
		status:                 res.PostProLimitOrder.Status,
		timeInForce:            res.PostProLimitOrder.TimeInForce,
		orderType:              res.PostProLimitOrder.OrderType,
		fees:                   res.PostProLimitOrder.Fees,
		filled:                 res.PostProLimitOrder.Filled,
		total:                  res.PostProLimitOrder.Total,
		initialBaseQuantity:    res.PostProLimitOrder.InitialBaseQuantity,
		initialQuoteQuantity:   res.PostProLimitOrder.InitialQuoteQuantity,
		remainingBaseQuantity:  res.PostProLimitOrder.RemainingBaseQuantity,
		remainingQuoteQuantity: res.PostProLimitOrder.RemainingQuoteQuantity,
		meanExecutionPrice:     res.PostProLimitOrder.MeanExecutionPrice,
		engineMessage:          res.PostProLimitOrder.EngineMessage,
	}, nil
}

func (config ConfigCredentials) GetDepositLink(amount float64) (getDepositLink, error) {
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(getDepositLinkQuery)
	req.Var("amount", amount)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()
	res := struct {
		CreateSendcashPayDeposit struct {
			Amount      string
			CreatedAt   int64
			Fee         string
			Id          string
			Link        string
			Reference   string
			Status      string
			TotalAmount string
			Type        string
		}
	}{}
	var err error
	if err = client.Run(ctx, req, &res); err != nil {
		log.Println(err)
		return getDepositLink{}, err
	}
	log.Println(res)

	return getDepositLink{
		Amount:      res.CreateSendcashPayDeposit.Amount,
		CreatedAt:   res.CreateSendcashPayDeposit.CreatedAt,
		Fee:         res.CreateSendcashPayDeposit.Fee,
		Id:          res.CreateSendcashPayDeposit.Id,
		Link:        res.CreateSendcashPayDeposit.Link,
		Reference:   res.CreateSendcashPayDeposit.Reference,
		Status:      res.CreateSendcashPayDeposit.Status,
		TotalAmount: res.CreateSendcashPayDeposit.TotalAmount,
		Type:        res.CreateSendcashPayDeposit.Type,
	}, nil
}

// Get balance should be used to fetch the balance for one cryptocurrency
func (config ConfigCredentials) GetBalance(crypto string) (getBalances, error) {
	var err error
	client := graphql.NewClient(endpoint)
	req := graphql.NewRequest(getBalancesQuery)
	req.Var("crypto", crypto)
	req.Header.Set("Authorization", config.basicAuth)
	ctx := context.Background()
	res := struct {
		GetBalances []struct {
			Id               string
			Cryptocurrency   string
			ConfirmedBalance string
		}
	}{}
	if err = client.Run(ctx, req, &res); err != nil {
		log.Println(err)
		return getBalances{}, err
	}
	log.Println(res)
	return getBalances{
		Id:               res.GetBalances[0].Id,
		Cryptocurrency:   res.GetBalances[0].Cryptocurrency,
		ConfirmedBalance: res.GetBalances[0].ConfirmedBalance,
	}, nil
}
