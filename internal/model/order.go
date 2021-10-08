package model


type CancelOrder struct {
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

type GetProOrderFees struct {
	Fee                string
	BaseCurrencyTotal  string
	QuoteCurrencyTotal string
	Price              string
}

type PostProMarketOrder struct {
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

type LimitOrder struct {
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

type GetProOrders struct {
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

type GetDepositLink struct {
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

