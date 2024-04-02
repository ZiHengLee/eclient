package binance

import (
	"testing"
	"time"

	"github.com/slowly-richer/eclient/utils/logger"
	"github.com/stretchr/testify/suite"
)

type websocketServiceTestSuite struct {
	baseTestSuite
	s *WebSocket
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(websocketServiceTestSuite))
}

func (s *websocketServiceTestSuite) SetupTest() {
	ws, err := NewWebSocket()
	if err != nil {
		logger.Error("create web socket err:%v", err)
		return
	}
	s.s = ws
	s.s.Run()
}

func (s *websocketServiceTestSuite) TearDownTest() {
	if s.s != nil {
		s.s.closeConn()
	}
}

func (s *websocketServiceTestSuite) TestGetKlines() {
	// klines, err := NewSubscriber(s.s)
	// s.Suite.Equal(err, nil)
	//ws.Subscribe([]string{"btcusdt@depth"})
	//ws.Subscribe([]string{"ethusdt@depth"})
	// ws.Subscribe([]string{"Ct7wPIBn1wAKvYcbC2nimpZCn83KhNvYJH2FAezulBwXQ0u0VMijmNI47lm5"})
	s.s.Subscribe([]string{"btcusdt@kline_1s"})
	time.Sleep(time.Second * 5)
	s.s.Subscribe([]string{"ethusdt@kline_1s"})
	time.Sleep(time.Second * 5)
	// for msg := range klines.Msgs() {
	// 	event := new(WsKlineEvent)
	// 	json.Unmarshal(msg, event)
	// 	fmt.Println(event)
	// }
	// for msg := range klines.Msgs() {
	// 	event := new(WsKlineEvent)
	// 	json.Unmarshal(msg, event)
	// 	fmt.Println(event)
	// }
}
