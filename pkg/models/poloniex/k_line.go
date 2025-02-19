package poloniex

// KlineRequest represents the query parameters for the API request
type KlineRequest struct {
	Symbol    string // Market symbol (e.g., BTC_USDT)
	Interval  string // Kline interval (e.g., 1m, 15m)
	Limit     int    // Number of klines (optional)
	StartTime int    // Start time in milliseconds (optional)
	EndTime   int    // End time in milliseconds (optional)
}

// KlineResponse represents the response structure of the Kline API
// type KlineResponse []KlineData

type KlineResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

// KlineData represents each candlestick's (Kline) data
type KlineData struct {
	LowestPrice    float64 // Lowest price
	HighestPrice   float64 // Highest price
	OpenPrice      float64 // Open price
	ClosePrice     float64 // Close price
	QuoteVolume    float64 // Quote asset volume
	NumberOfTrades int     // Number of trades
	Trades         int     // Volume of the trades
	StartTime      int     // Start of the kline
	EndTime        int     // End of the kline
}

/**
l	String	Lowest price
h	String	Highest price
o	String	Opening price
c	String	Closing price
amt	String	Trading unit, the quantity of the quote currency.
qty	String	Trading unit, the quantity of the base currency, or Cont for the number of contracts.
tC	String	Trades
sT	String	Start time
cT	String	End time

0 = {interface{} | string} "96067.89"
1 = {interface{} | string} "96097.17"
2 = {interface{} | string} "96087.14"
3 = {interface{} | string} "96081.2"
4 = {interface{} | string} "2113.78388"
5 = {interface{} | string} "22"
6 = {interface{} | string} "9"
7 = {interface{} | string} "1739989500000"
8 = {interface{} | string} "1739989559999"
*/
