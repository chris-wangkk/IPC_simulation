package ossmgr

import (
	"errors"
	"fmt"
	"time"
)

var (
	SYN_REQ_MSG   = uint16(0x0001)
	ASYN_REQ_MSG  = uint16(0x0002)
	TIMER_REQ_MSG = uint16(0x0003)

	maxQueue = 1000
)

type gwSched struct {
	modules map[string]chan Message
	//sndQueue chan Message
	rcvQueue chan Message
}

var schedIns *gwSched

func (s *gwSched) synSend(event uint16, data interface{}, tmout time.Duration, receiver string) error {
	if _, bExist := s.modules[receiver]; false == bExist {
		return errors.New("no such receiver")
	}
	msg := Message{Event: event, MsgType: SYN_REQ_MSG, Data: data}
	select {
	case s.modules[receiver] <- msg:
		return nil
	case <-time.After(tmout):
		return errors.New("timeout")
	}
}

func (s *gwSched) asynSend(event uint16, data interface{}, receiver string) error {
	if _, bExist := s.modules[receiver]; false == bExist {
		return errors.New("no such receiver")
	}
	msg := Message{Event: event, MsgType: ASYN_REQ_MSG, Data: data}
	go func() {
		s.modules[receiver] <- msg
	}()
	return nil
}

func (s *gwSched) createTimer(mid string, event uint16, tCnt time.Duration) error {

	if _, bExist := s.modules[mid]; false == bExist {
		return errors.New("no such module")
	}
	timer := time.NewTimer(tCnt)
	go func() {
		<-timer.C
		msg := Message{Event: event, MsgType: TIMER_REQ_MSG}
		s.modules[mid] <- msg
	}()
	return nil
}

func (s *gwSched) createTicker(mid string, event uint16, tCnt time.Duration) error {
	if _, bExist := s.modules[mid]; false == bExist {
		return errors.New("no such module")
	}
	ticker := time.NewTicker(tCnt)
	go func() {
		for {
			<-ticker.C
			msg := Message{Event: event, MsgType: TIMER_REQ_MSG}
			s.modules[mid] <- msg
		}
	}()
	return nil
}

func (s *gwSched) dispatch() {
	for {
		msg := <-s.rcvQueue
		if _, bExist := s.modules[msg.receiver]; false == bExist {
			fmt.Println("no such receiver")
			continue
		}
		s.modules[msg.receiver] <- msg
	}
}

func (s *gwSched) run() {
	for k, v := range gModules {
		s.modules[k] = make(chan Message, 5)
		go v.Execute(s.modules[k])
	}
	go s.dispatch()
}

func (s *gwSched) debugPrint(paras ...interface{}) {
	/*
		paras := make([]interface{})
		for _, para := range paras{
			paras = paras
		}
	*/
	return
}

func SSEND(event uint16, data interface{}, tmout time.Duration, receiver string) error {
	return schedIns.synSend(event, data, tmout, receiver)
}

func ASEND(event uint16, data interface{}, receiver string) error {
	return schedIns.asynSend(event, data, receiver)
}

func SetTimer(mid string, event uint16, tCnt time.Duration) error {
	return schedIns.createTimer(mid, event, tCnt)
}

func SetTicker(mid string, event uint16, tCnt time.Duration) error {
	return schedIns.createTicker(mid, event, tCnt)
}

func newGwSched() *gwSched {
	return &gwSched{
		modules:  make(map[string]chan Message),
		rcvQueue: make(chan Message, maxQueue),
	}
}

func SchedInit() {
	fmt.Println("sched init")
	schedIns = newGwSched()
	schedIns.run()
}
