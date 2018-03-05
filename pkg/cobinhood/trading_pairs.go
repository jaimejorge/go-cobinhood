package cobinhood

type TradingPair struct {
	Id                string `json:"Id"`
	BaseCurrencyId  string `json:"Base_currency_id"`
	QuoteCurrencyId string `json:"Quote_currency_id"`
	BaseMinSize     string `json:"Base_min_size"`
	BaseMaxSize     string `json:"Base_max_size"`
	QuoteIncrement   string `json:"Quote_increment"`
}

type TradingPairs []TradingPair

func (c *CobinhoodClient) GetTradingPairs() (TradingPairs, error) {
	var Respond Respond

	err := c.Get("/v1/market/trading_pairs", &Respond)

	if err != nil {
		return Respond.Result.TradingPairs, err
	}

	return Respond.Result.TradingPairs, nil
}
