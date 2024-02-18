package binance

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"

	"github.com/ZiHengLee/eclient/utils/logger"
)

type Subscriber struct {
	ws *WebSocket

	msgs chan []byte
}

func NewSubscriber(ws *WebSocket) (s *Subscriber, err error) {
	s = &Subscriber{
		ws:   ws,
		msgs: make(chan []byte, 1024),
	}
	ws.addSubscriber(s)
	return
}

func (s *Subscriber) Close() {
	s.ws.removeSubscriber(s)
	close(s.msgs)
}

func (s *Subscriber) Msgs() <-chan []byte {
	return s.msgs
}

type wsEvent struct {
	stream string
}

type WebSocket struct {
	mtx         sync.Mutex
	c           *websocket.Conn
	sendChannel chan []byte
	msgId       int64

	sbrs   []*Subscriber
	events map[string]*wsEvent

	sbrTimer *time.Timer
}

var sbrWaitTime = 5 * time.Millisecond

func NewWebSocket() (s *WebSocket, err error) {
	s = &WebSocket{
		sendChannel: make(chan []byte, 1024),
		events:      make(map[string]*wsEvent),
		sbrTimer:    time.NewTimer(sbrWaitTime),
	}
	return
}

func (s *WebSocket) Run() {
	go s.read()
	go s.send()
}

func (s *WebSocket) getConn() *websocket.Conn {
	s.mtx.Lock()
	c := s.c
	if c != nil {
		s.mtx.Unlock()
		return c
	}
	for c == nil {
		var err error
		c, _, err = websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws", nil)
		if err != nil {
			logger.Warn("dial websocket err:%v", err)
			time.Sleep(time.Second)
			continue
		}
		s.c = c
	}
	var ss []string
	if len(s.events) > 0 {
		ss = make([]string, 0, len(s.events))
		for _, s := range s.events {
			ss = append(ss, s.stream)
		}
	}
	s.mtx.Unlock()
	if len(ss) > 0 {
		s.subscribe(ss)
	}
	return c
}

func (s *WebSocket) closeConn() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.c != nil {
		s.c.Close()
		s.c = nil
	}
}

func (s *WebSocket) read() {
	for {
		c := s.getConn()
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				logger.Warn("read msg err:%v", err)
				s.closeConn()
				break
			}
			logger.Trace("recv msg:%v", string(msg))
			s.deliver(msg)
		}
	}
}

func (s *WebSocket) deliver(msg []byte) {
	s.mtx.Lock()
	sbrs := make([]*Subscriber, len(s.sbrs))
	copy(sbrs, s.sbrs)
	s.mtx.Unlock()
	for _, sbr := range sbrs {
		if !s.sbrTimer.Stop() {
			<-s.sbrTimer.C
		}
		s.sbrTimer.Reset(sbrWaitTime)
		select {
		case sbr.msgs <- msg:
		case <-s.sbrTimer.C:
			logger.Warn("deliver msg ignore")
		}
	}
}

func (s *WebSocket) send() {
	for dat := range s.sendChannel {
		c := s.getConn()
		logger.Info("write msg:%v", string(dat))
		err := c.WriteMessage(websocket.TextMessage, dat)
		if err != nil {
			s.closeConn()
			logger.Warn("write msg err:%v", err)
		}
	}
}

func (s *WebSocket) Write(dat []byte) {
	s.sendChannel <- dat
}

func (s *WebSocket) incrMsgId() int64 {
	return atomic.AddInt64(&s.msgId, 1)
}

func (s *WebSocket) addSubscriber(sbr *Subscriber) (err error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.sbrs = append(s.sbrs, sbr)
	return
}

func (s *WebSocket) removeSubscriber(sbr *Subscriber) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	j := -1
	for i, sb := range s.sbrs {
		if sb == sbr {
			j = i
			break
		}
	}
	if j >= 0 {
		for i := j + 1; i < len(s.sbrs); i++ {
			s.sbrs[i-1] = s.sbrs[i]
		}
		s.sbrs = s.sbrs[:len(s.sbrs)-1]
	}
}

func (s *WebSocket) Subscribe(streams []string) {
	s.mtx.Lock()
	ss := []string{}
	for _, name := range streams {
		if len(name) == 0 {
			continue
		}
		if _, ok := s.events[name]; ok {
			continue
		}
		ss = append(ss, name)
		s.events[name] = &wsEvent{
			stream: name,
		}
	}
	s.mtx.Unlock()

	s.subscribe(ss)
}

func (s *WebSocket) subscribe(ss []string) {
	if len(ss) > 0 {
		id := s.incrMsgId()
		args := map[string]interface{}{
			"method": "SUBSCRIBE",
			"params": ss,
			"id":     id,
		}
		logger.Info("subscribe streams:%v", ss)
		dat, _ := json.Marshal(args)
		s.Write(dat)
	}
}
