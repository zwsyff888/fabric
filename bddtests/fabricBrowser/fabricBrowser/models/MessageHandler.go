package models

import (
	"sync"
)

var SocketsProperty Property

type state struct {
	next *state
	done chan struct{}
}

func newState() *state {
	return &state{
		done: make(chan struct{}),
	}
}

func (s *state) update() *state {
	s.next = newState()
	close(s.done)
	return s.next
}

type Stream interface {
	//新内容到达时尝试去改变
	Changes() chan struct{}

	//寻找下一个订阅者
	Next()

	//判断是否还存在下一个订阅者
	HasNext() bool

	// 等待当前订阅被关闭
	WaitNext()

	// 复制当前订阅
	Clone() Stream
}

type stream struct {
	state *state
}

func (s *stream) Clone() Stream {
	return &stream{state: s.state}
}

func (s *stream) Changes() chan struct{} {
	return s.state.done
}

func (s *stream) Next() {
	s.state = s.state.next
}

func (s *stream) HasNext() bool {
	select {
	case <-s.state.done:
		return true
	default:
		return false
	}
}

func (s *stream) WaitNext() {
	<-s.state.done
	s.state = s.state.next
}

type Property interface {
	Update()
	Observe() Stream
}

func (p *property) Update() {
	p.Lock()
	defer p.Unlock()
	p.state = p.state.update()
}

func NewProperty() Property {
	return &property{state: newState()}
}

type property struct {
	sync.RWMutex
	state *state
}

func (p *property) Observe() Stream {
	p.RLock()
	defer p.RUnlock()
	return &stream{state: p.state}
}
