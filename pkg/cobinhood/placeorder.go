package cobinhood

import (
	"bytes"
	"encoding/json"
)

type PlaceOrder struct {
	TradingPairId string `json:"trading_pair_id"`
	Side            string `json:"side"`
	Type            string `json:"type"`
	Price           string `json:"price"`
	Size            string `json:"size"`
}

func (c *CobinhoodClient) PlaceOrder(datajson PlaceOrder) (Order, error) {

	var respond Respond
	data := new(bytes.Buffer)
	err := json.NewEncoder(data).Encode(datajson)
	err = c.PostBody("/v1/trading/orders", data, &respond)

	if err != nil {
		return respond.Result.Order, err
	}

	return respond.Result.Order, nil
}
