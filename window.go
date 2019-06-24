package STWindow

import (
	"sync"
	"time"
)

type Root struct {
	head *node
	tail *node
	mutex sync.RWMutex
	period time.Duration
}

func NewArray(t time.Duration) *Root {
	return &Root{
		head: nil,
		tail: nil,
		period: t,
	}
}

type node struct {
	Next *node
	Value interface{}
	ts int64
}

func (a *Root)Append(i interface{}) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	nowTs := time.Now().UnixNano()
	deadline := nowTs - a.period.Nanoseconds()

	n := &node {
		ts: nowTs,
		Next: nil,
		Value: i,
	}

	if a.head != nil {
		a.moveOutExpired(deadline)
	}

	if a.tail == nil {
		a.head = n
		a.tail = n
	}else {
		a.tail.Next = n
		a.tail = n
	}
}

func (a *Root)List()[]interface{} {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	var array []interface{}
	var p = a.head
	for p != nil {
		array = append(array, p.Value)
		p = p.Next
	}

	return array
}

func (a *Root)moveOutExpired(deadline int64) {
	if a.head == nil {
		a.tail = nil
	}else if a.head.ts > deadline {
		return
	}else {
		a.head = a.head.Next
		a.moveOutExpired(deadline)
	}
}