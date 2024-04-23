package eastmoney

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type baseTestSuite struct {
	suite.Suite
	client *Client
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) SetupTest() {
	cli := NewClient()
	s.client = cli
}

type emTestSuite struct {
	baseTestSuite
}

func TestEmService(t *testing.T) {
	suite.Run(t, new(emTestSuite))
}
func (s *emTestSuite) TestGetStockInfo() {
	ctx := context.Background()
	param := DefaultStockInfoParam
	param.Secid = "0.000002"
	_, err := s.client.StockInfo(ctx, &param).Get()
	s.Equal(nil, err)
}
