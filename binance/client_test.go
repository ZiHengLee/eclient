package binance

import (
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
