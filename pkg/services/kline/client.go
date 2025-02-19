package kline

import (
	"app/pkg/models/poloniex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const poloniexAPIBase = "https://api.poloniex.com/v3/market/candles"

type PoloniexClient struct {
	key              string
	signature        string
	signatureMethod  string
	signatureVersion string
	signTimestamp    string
}

func NewPoloniexClient(key, signature, signatureMethod, signatureVersion, signTimestamp string) *PoloniexClient {
	return &PoloniexClient{
		key:              key,
		signature:        signature,
		signatureMethod:  signatureMethod,
		signatureVersion: signatureVersion,
		signTimestamp:    signTimestamp,
	}
}

func (c PoloniexClient) FetchKlineData(request poloniex.KlineRequest) ([]poloniex.KlineData, error) {
	url := fmt.Sprintf(
		//"%s?symbol=%s&interval=%s&limit=%d&sTime=%d&eTime=%d",
		"%s?symbol=%s&interval=%s&limit=%d&sTime=%d",
		poloniexAPIBase,
		request.Symbol,
		request.Interval,
		request.Limit,
		request.StartTime,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("key", c.key)
	req.Header.Add("signatureMethod", c.signatureMethod)
	req.Header.Add("signatureVersion", c.signatureVersion)
	req.Header.Add("signTimestamp", c.signTimestamp)
	req.Header.Add("signature", c.signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error response from API: %s", body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var klineResponse poloniex.KlineResponse
	if err := json.Unmarshal(body, &klineResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %w", err)
	}

	klineData := make([]poloniex.KlineData, len(klineResponse.Data))
	for i, item := range klineResponse.Data {
		// todo log error and skip
		lowestPrice, _ := strconv.ParseFloat(item[0], 64)
		highestPrice, _ := strconv.ParseFloat(item[1], 64)
		openPrice, _ := strconv.ParseFloat(item[2], 64)
		closePrice, _ := strconv.ParseFloat(item[3], 64)
		quoteVolume, _ := strconv.ParseFloat(item[4], 64)
		numberOfTrades, _ := strconv.Atoi(item[5])
		trades, _ := strconv.Atoi(item[6])
		startTime, _ := strconv.Atoi(item[7])
		endTime, _ := strconv.Atoi(item[8])

		dataPoint := poloniex.KlineData{
			LowestPrice:    lowestPrice,
			HighestPrice:   highestPrice,
			OpenPrice:      openPrice,
			ClosePrice:     closePrice,
			QuoteVolume:    quoteVolume,
			NumberOfTrades: numberOfTrades,
			Trades:         trades,
			StartTime:      startTime,
			EndTime:        endTime,
		}
		klineData[i] = dataPoint
	}

	return klineData, nil
	// 1733403600000
	// 1739999345238

	// 1739999400000
	// 1739999345238
}
