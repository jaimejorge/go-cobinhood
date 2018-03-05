package cobinhood

import (
	"strconv"
	"strings"
	"time"
)


type Respond struct {
	success bool   `json:"success"`
	Result  Result  `json:"result"`
}
type Result struct {
	Info          Info  `json:"info"`
	Currencies    Currencies `json:"currencies"`
	TradingPairs   TradingPairs `json:"trading_pairs"`
	Orderbook     Orderbook `json:"orderbook"`
	Ticker        Ticker `json:"ticker"`
	Trades        Trades `json:"trades"`
	Order         Order `json:"order"`
	Orders        []Order `json:"orders"`
	Error         Error `json:"error"`
}

type jsonTime time.Time

func (t jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *jsonTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q/1000, 0)
	return
}

func (t jsonTime) StringHour() string { return time.Time(t).Format("15:04") }
func (t jsonTime) String() string     { return time.Time(t).String() }
