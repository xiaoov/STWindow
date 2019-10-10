package STWindow

/*
#include <stdint.h>
#include <time.h>
#include <stdlib.h>
#include <stdio.h>

typedef struct Node{
    struct Node *Next;
    int Value;
    uint32_t ts;
}Node;

Node *Append(int i){
	time_t t;
	t = time(NULL);

	int now_ts = time(&t);

	Node *n;
	n = (Node*)malloc(sizeof(Node));
	n->ts = now_ts;
	n->Value = i;
	n->Next = 0;

	return n;
}

Node *Remove(Node *n){
	Node *temp;
	temp = n->Next;
	free(n);
	return temp;
}
*/
import "C"

import (
	"log"
	"sync"
	"time"
)

//func  main () {
//	a := NewArray(2)
//	a.AppendInt(1)
//	a.AppendInt(2)
//
//	a.Print()
//
//	time.Sleep(3 * time.Second)
//
//	a.Print()
//}

type Root struct {
	head   *C.struct_Node
	tail   *C.struct_Node
	mutex  sync.RWMutex
	period time.Duration
	ticker *time.Ticker
}

func NewArray(t time.Duration) *Root {
	r := &Root{
		head:   nil,
		tail:   nil,
		period: t,
		ticker: time.NewTicker(1 * time.Second),
	}

	go r.periodicClear()

	return r
}

//
//type node struct {
//	Next *node
//	Value interface{}
//	ts int64
//}
//
func (a *Root) AppendInt(i int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	n := C.Append(C.int(i))

	if a.tail == nil {
		a.head = n
		a.tail = n
	} else {
		a.tail.Next = n
		a.tail = n
	}
}

//
func (a *Root) List() []interface{} {
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

//
func (a *Root) periodicClear() {
	for {
		nowTs := time.Now().Unix()
		deadline := uint64(nowTs - int64(a.period.Seconds()))

		select {
		case <-a.ticker.C:
			a.moveOutExpired(deadline)
		}
	}
}

//
func (a *Root) moveOutExpired(deadline uint64) {
	if a.head == nil {
		a.tail = nil
	} else if uint64(a.head.ts) > deadline {
		return
	} else {
		next := C.Remove(a.head)
		a.head = next
		a.moveOutExpired(deadline)
	}
}

func (a *Root) Print() {
	next := a.head
	for next != nil {
		log.Println("---")
		//value := *(*int)(next.Value)
		log.Println("value :", int(uintptr(next.Value)))
		log.Println("ts :", int(uintptr(next.ts)))
		log.Println("---")
		next = next.Next
	}
}
