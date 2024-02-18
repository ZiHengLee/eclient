package binance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type accountServiceTestSuite struct {
	baseTestSuite
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountServiceTestSuite))
}

func (s *accountServiceTestSuite) TestGetAccount() {
	ctx := context.Background()
	resp, err := s.client.AccountInfo(ctx, nil).Get()
	output(resp, err)
}
