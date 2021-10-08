package orderbook

import (
	"github.com/buycoinsresearch/buycoins-orderbook-go/internal/model"
)

type OrderBook interface {
	GetPairs() ([]byte, error)
	GetOrders(coinPair, status, side string) (*model.GetProOrders, error)
	CancelOrder(id string) (*model.CancelOrder, error)
	GetProOrderFees(orderType string, pair string, side string, amount float64) (*model.GetProOrderFees, error)
	PostProMarketOrder(pair string, quantity float64, side string) (*model.PostProMarketOrder, error)
	PostProLimitOrder(pair string, quantity float64, price float64, side string, timeInForce string) (*model.LimitOrder, error)
	GetDepositLink(amount float64) (*model.GetDepositLink, error)
}
