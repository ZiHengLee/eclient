package usd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/slowly-richer/eclient/binance"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type baseTestSuite struct {
	suite.Suite
	client    *Client
	apiKey    string
	secretKey string
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) SetupTest() {
	s.apiKey = "ABGIfT4VzEiqerwziXSNi38Jjm9xAU7fWkYxIYtG2LzcO1aSosap4IHCD3rbAilo"
	s.secretKey = "ZHVGW0yeEQfqm3m9ktgZGbjVa6szBQ4ECL8Te3Icd69e26mnoENKjJxEtS44fvaE"
	cli := NewClient(s.apiKey, s.secretKey)
	s.client = cli
}

type usdTestSuite struct {
	baseTestSuite
}

func TestUsdService(t *testing.T) {
	suite.Run(t, new(usdTestSuite))
}
func (s *usdTestSuite) TestGetAccount() {
	ctx := context.Background()
	param := &UsdAccountParam{BaseParam: binance.BaseParam{Timestamp: time.Now().UnixMilli(), RecvWindow: 5000}}
	_, err := s.client.UsdAccount(ctx, param).Get()
	s.Equal(nil, err)
}

func (s *usdTestSuite) TestGetAllOrders() {
	ctx := context.Background()
	param := &UsdAllOrdersParam{BaseParam: binance.BaseParam{Timestamp: time.Now().UnixMilli(), RecvWindow: 5000}}
	resp, err := s.client.UsdAllOrders(ctx, param).Get()
	s.Equal(nil, err)
	fmt.Println(resp.GetData())
}

func (s *usdTestSuite) TestGetOpenOrder() {
	ctx := context.Background()
	param := &UsdOpenOrderParam{
		BaseParam: binance.BaseParam{Timestamp: time.Now().UnixMilli(), RecvWindow: 5000},
		Symbol:    "BTCUSDT",
		OrderId:   279134299467,
	}
	resp, err := s.client.UsdOpenOrder(ctx, param).Get()
	s.Equal(nil, err)
	fmt.Println(resp.GetData())
}

func (s *usdTestSuite) TestOpenOrder() {
	ctx := context.Background()
	param := &UsdOrderParam{
		BaseParam:    binance.BaseParam{Timestamp: time.Now().UnixMilli(), RecvWindow: 5000},
		Symbol:       "BTCUSDT",
		Side:         "BUY",
		PositionSide: "LONG",
		Type:         "LIMIT",
		Price:        "67000",
		Quantity:     "0.01",
	}
	resp, err := s.client.UsdOrder(ctx, param).Get()
	s.Equal(nil, err)
	fmt.Println(resp.GetData())
}
