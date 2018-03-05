## [Cobinhood](https://cobinhood.com) command-line interface

### Overview

Command-line interface that allows to check the buy, sell and get information of orders and trades.


### Configuration
- Create an [apikey](https://cobinhood.com/api) 

```bash
export apikey=xxxxxxxxxxxx
```

### Running
Available arguments:

```
Usage of go-cobinhood:
  buy         Place a buy order.
  buymarket   Buy price market.
  cancel      Cancel order.
  info        Get system information.
  market      Get market information,currencies,orderbooks,ticker,trades.
  orders      Get orders from user
  sell        Place a sell order.
  sellmarket  Sell market price.
```

## Examples

Following provides list of main usages of command line.

### Get System Information

```bash
go-cobinhood info 
```

### Get All Currencies

```bash
go-cobinhood market currencies 
```

### Get All Trading Pairs

```bash
go-cobinhood market trading_pairs
```
### Get Order Book

```bash
go-cobinhood market orderbooks COB-USDT
```
### Get Ticker

```bash
go-cobinhood market ticker COB-USDT
```
### Get Recent Trades

```bash
go-cobinhood market trades COB-USDT
```
### Get All Orders

```bash
go-cobinhood orders open
```

### Get Order

```bash
go-cobinhood orders info <orderid>
```
### Get Trades of An Order

```bash
go-cobinhood orders orderTrades <orderid>
```

### Buy Order 

```bash
go-cobinhood buy  COB-USDT  <price> <size>
```

### Buy Order Market Price

```bash
go-cobinhood buymarket  COB-USDT  <size>
```
### Sell Order 

```bash
go-cobinhood sell  COB-USDT  <price> <size>
```

### Buy Order Market Price

```bash
go-cobinhood sellmarket  COB-USDT  <size>
```

### Cancel Order

```bash
go-cobinhood cancel  COB-USDT  <size>
```


### Get Order History

```bash
go-cobinhood orders history
```