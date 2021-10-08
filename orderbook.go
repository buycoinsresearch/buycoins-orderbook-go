package orderbooks

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/buycoinsresearch/buycoins-orderbook-go/internal/model"
	"github.com/buycoinsresearch/buycoins-orderbook-go/pkg/utils"
	"github.com/machinebox/graphql"
	"github.com/pkg/errors"
	"log"
)


type configCredentials struct {
	basicAuth string
	client *graphql.Client
}

func BuyCoins(publicKey, secretKey string) *configCredentials {
	auth := "Basic " + b64.URLEncoding.EncodeToString([]byte(publicKey+":"+secretKey))
	return &configCredentials{
		basicAuth: auth,
		client: graphql.NewClient(utils.Endpoint),
	}
}

func (config *configCredentials) GetPairs() ([]byte, error) {
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

	if err := config.client.Run(ctx, req, &res); err != nil {
		log.Fatal(errors.Wrap(err,"Get Pairs"))
	}

	pairs, err := json.MarshalIndent(res.GetPairs, "", "  ")
	if err != nil {
		fmt.Println(errors.Wrap(err,"Get Pairs"))
	}

	return pairs, nil
}

func (config *configCredentials) GetOrders(coinPair, status, side string) (*model.GetProOrders, error) {
	req := graphql.NewRequest(`
		query ($pair_: Pair!, $status_: ProOrderStatus!, $side_: OrderSide!) {
			getProOrders (pair:$pair_, status:$status_, side:$side_) {
				edges {
				  node {
					id
					pair
					price
					side
					status
					timeInForce
					orderType
					fees
					filled
					total
					initialBaseQuantity
					initialQuoteQuantity
					remainingBaseQuantity
					remainingQuoteQuantity
					meanExecutionPrice
					engineMessage
				  }
    			}
			}
		}
	`)
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

	if err := config.client.Run(ctx, req, &res); err != nil {
		log.Println(errors.Wrap(err,"Get Orders"))
		return nil, err
	}
	fmt.Println("Successfully connected")

	return &model.GetProOrders{
		Edges: res.GetProOrders.Edges,
	}, nil
}

func (config *configCredentials) CancelOrder(id string) (*model.CancelOrder, error) {
	req := graphql.NewRequest(`
			mutation($id: ID!) {
				cancelOrder(proOrder: $id){
				id
				pair
				price
				side
				status
				timeInForce
				orderType
				fees
				filled
				total
				initialBaseQuantity
				initialQuoteQuantity
				remainingBaseQuantity
				remainingQuoteQuantity
				meanExecutionPrice
				engineMessage
				}
			}
	`)

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

	if err := config.client.Run(ctx, req, &res); err != nil {
		log.Println(errors.Wrap(err,"Cancel Order"))
		return nil, err
	}

	return &model.CancelOrder{
		Id:                     res.CancelOrder.Id,
		Pair:                   res.CancelOrder.Pair,
		Price:                  res.CancelOrder.Price,
		Side:                   res.CancelOrder.Side,
		Status:                 res.CancelOrder.Status,
		TimeInForce:            res.CancelOrder.TimeInForce,
		OrderType:              res.CancelOrder.OrderType,
		Fees:                   res.CancelOrder.Fees,
		Filled:                 res.CancelOrder.Filled,
		Total:                  res.CancelOrder.Total,
		InitialBaseQuantity:    res.CancelOrder.InitialBaseQuantity,
		InitialQuoteQuantity:   res.CancelOrder.InitialQuoteQuantity,
		RemainingBaseQuantity:  res.CancelOrder.RemainingBaseQuantity,
		RemainingQuoteQuantity: res.CancelOrder.RemainingQuoteQuantity,
		MeanExecutionPrice:     res.CancelOrder.MeanExecutionPrice,
		EngineMessage:          res.CancelOrder.EngineMessage,
	}, nil

}

func (config *configCredentials) GetProOrderFees(orderType string, pair string, side string, amount float64) (*model.GetProOrderFees, error) {
	req := graphql.NewRequest(`
		query($orderType_: OrderMatchingEngineOrder!, $pair_: Pair!, $side_: OrderSide!, $amount_: BigDecimal!) {
			getProOrderFees(orderType: $orderType_, pair: $pair_, side: $side_, amount: $amount_){
			fee
			baseCurrencyTotal
			quoteCurrencyTotal
			price
			}
		}
	`)

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

	if err := config.client.Run(ctx, req, &res); err != nil {
		log.Println(errors.Wrap(err,"Get Pro Order Fees"))
		return nil, err
	}

	return &model.GetProOrderFees{
		Fee:                res.GetProOrderFees.Fees,
		BaseCurrencyTotal:  res.GetProOrderFees.BaseCurrencyTotal,
		QuoteCurrencyTotal: res.GetProOrderFees.QuoteCurrencyTotal,
		Price:              res.GetProOrderFees.Price,
	}, nil
}

func (config *configCredentials) PostProMarketOrder(pair string, quantity float64, side string) (*model.PostProMarketOrder, error) {
	req := graphql.NewRequest(`
		mutation($pair_: Pair!, $quantity_: BigDecimal!, $side_: OrderSide!) {
			postProMarketOrder(pair: $pair_, quantity: $quantity_, side: $side_){
			id
			pair
			price
			side
			status
			timeInForce
			orderType
			fees
			filled
			total
			initialBaseQuantity
			initialQuoteQuantity
			remainingBaseQuantity
			remainingQuoteQuantity
			meanExecutionPrice
			engineMessage
			}
		}
	`)

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


	if err := config.client.Run(ctx, req, &res); err != nil {
		log.Println(errors.Wrap(err,"Post Pro Market Order"))
		return nil, err
	}

	return &model.PostProMarketOrder{
		Id:                     res.PostProMarketOrder.Id,
		Pair:                   res.PostProMarketOrder.Pair,
		Price:                  res.PostProMarketOrder.Price,
		Side:                   res.PostProMarketOrder.Side,
		Status:                 res.PostProMarketOrder.Status,
		TimeInForce:            res.PostProMarketOrder.TimeInForce,
		OrderType:              res.PostProMarketOrder.OrderType,
		Fees:                   res.PostProMarketOrder.Fees,
		Filled:                 res.PostProMarketOrder.Filled,
		Total:                  res.PostProMarketOrder.Total,
		InitialBaseQuantity:    res.PostProMarketOrder.InitialBaseQuantity,
		InitialQuoteQuantity:   res.PostProMarketOrder.InitialQuoteQuantity,
		RemainingBaseQuantity:  res.PostProMarketOrder.RemainingBaseQuantity,
		RemainingQuoteQuantity: res.PostProMarketOrder.RemainingQuoteQuantity,
		MeanExecutionPrice:     res.PostProMarketOrder.MeanExecutionPrice,
		EngineMessage:          res.PostProMarketOrder.EngineMessage,
	}, nil
}

func (config *configCredentials) PostProLimitOrder(pair string, quantity float64, price float64, side string, timeInForce string) (*model.LimitOrder, error) {
	req := graphql.NewRequest(`
		mutation($pair_: Pair!, $quantity_: BigDecimal!, $price_: BigDecimal! $side_: OrderSide!, $timeInForce_: TimeInForce!) {
			postProLimitOrder(pair: $pair_, quantity: $quantity_, price: $price_ side: $side_, timeInForce: $timeInForce_){
			id
			pair
			price
			side
			status
			timeInForce
			orderType
			fees
			filled
			total
			initialBaseQuantity
			initialQuoteQuantity
			remainingBaseQuantity
			remainingQuoteQuantity
			meanExecutionPrice
			engineMessage
			}
		}
	`)

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
	if err = config.client.Run(ctx, req, &res); err != nil {
		log.Println(errors.Wrap(err,"Post Pro Limit Order"))
		return nil, err
	}

	return &model.LimitOrder{
		Id:                     res.PostProLimitOrder.Id,
		Pair:                   res.PostProLimitOrder.Pair,
		Price:                  res.PostProLimitOrder.Price,
		Side:                   res.PostProLimitOrder.Side,
		Status:                 res.PostProLimitOrder.Status,
		TimeInForce:            res.PostProLimitOrder.TimeInForce,
		OrderType:              res.PostProLimitOrder.OrderType,
		Fees:                   res.PostProLimitOrder.Fees,
		Filled:                 res.PostProLimitOrder.Filled,
		Total:                  res.PostProLimitOrder.Total,
		InitialBaseQuantity:    res.PostProLimitOrder.InitialBaseQuantity,
		InitialQuoteQuantity:   res.PostProLimitOrder.InitialQuoteQuantity,
		RemainingBaseQuantity:  res.PostProLimitOrder.RemainingBaseQuantity,
		RemainingQuoteQuantity: res.PostProLimitOrder.RemainingQuoteQuantity,
		MeanExecutionPrice:     res.PostProLimitOrder.MeanExecutionPrice,
		EngineMessage:          res.PostProLimitOrder.EngineMessage,
	}, nil
}

func (config *configCredentials) GetDepositLink(amount float64) (*model.GetDepositLink, error) {
	req := graphql.NewRequest(`
		mutation ($amount: BigDecimal!) {
			createSendCashPayDeposit(amount: $amount){
				amount
				createdAt
				fee
				id
				link
				reference
				status
				totalAmount
				type
			}
		}
	`)
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

	if err := config.client.Run(ctx, req, &res); err != nil {
		log.Println(errors.Wrap(err,"Get Deposit Link"))
		return nil, err
	}
	log.Println(res)

	return &model.GetDepositLink{
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