package binance

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type usdTestSuite struct {
	baseTestSuite
}

func TestUsdService(t *testing.T) {
	suite.Run(t, new(usdTestSuite))
}
func (s *usdTestSuite) TestGetAccount() {
	ctx := context.Background()
	param := &UsdAccountParam{BaseParam: BaseParam{Timestamp: time.Now().UnixMilli(), RecvWindow: 5000}}
	_, err := s.client.UsdAccount(ctx, param).Get()
	s.Equal(nil, err)
}

func (s *usdTestSuite) TestGetAllOrders() {
	ctx := context.Background()
	param := &UsdAllOrdersParam{BaseParam: BaseParam{Timestamp: time.Now().UnixMilli(), RecvWindow: 5000}}
	resp, err := s.client.UsdAllOrders(ctx, param).Get()
	s.Equal(nil, err)
	fmt.Println(resp.GetData())
}

func (s *usdTestSuite) TestGetOpenOrder() {
	ctx := context.Background()
	param := &UsdOpenOrderParam{
		BaseParam: BaseParam{Timestamp: time.Now().UnixMilli(), RecvWindow: 5000},
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
		BaseParam:    BaseParam{Timestamp: time.Now().UnixMilli(), RecvWindow: 5000},
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
