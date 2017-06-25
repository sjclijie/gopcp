package main

import (
	"fmt"
	"time"
)

type RawReq struct {
	ID  int64
	Req []byte
}

type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

type RetCode int

type CallResult struct {
	ID     int64
	Req    RawReq
	Resp   RawResp
	Code   RetCode
	Msg    string
	Elapse time.Duration
}

type Caller interface {
	BuildReq() RawReq
	Call(req []byte, timeoutNs, duration time.Duration) ([]byte, error)
	CheckResp(rawReq RawReq, rawResp RawResp) *CallResult
}

type GoTickets interface {
	Take()
	Return()
	Active() bool
	Total() uint32
	Remainder() uint32
}

type myGoTickets struct {
	total    uint32
	ticketCh chan struct{}
	active   bool
}

type Generator interface {
	Start() bool
	Stop() bool
	Status() uint32
	CallCount() int64
}

type myGenerator struct {
	caller      Caller
	timeoutNs   time.Duration
	lps         uint32
	durationNs  time.Duration
	concurrency uint32
}

func (gen *myGenerator) Start() bool {
	return true
}

func (gen *myGenerator) Stop() bool {
	return true
}

func (gen *myGenerator) Status() uint32 {
	return 3333
}

func (gen *myGenerator) CallCount() uint64 {
	return 444
}

func NewGenerator() (Generator, error) {

	gen := &myGenerator{}

	return gen, nil
}

func main() {

	fmt.Println("hello")
}
