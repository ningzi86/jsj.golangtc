
package models

import (
	"sync"
    "container/list"
)

type Message struct {
	Msgid int
	Msg   string
}

type QueueScheduler struct {
	Locker *sync.Mutex
	Queue  *list.List
}

func (this *QueueScheduler) Push(message *Message) {
	this.Locker.Lock()
	this.Queue.PushBack(message)
	this.Locker.Unlock()
}

func (this *QueueScheduler) Pull() *Message {

	this.Locker.Lock()
	if this.Queue.Len() <= 0 {
		this.Locker.Unlock()
		return nil
	}

	e := this.Queue.Front()
	message := e.Value.(*Message)

	this.Queue.Remove(e)
	this.Locker.Unlock()

	return message

}

func (this *QueueScheduler) Count() int {

	this.Locker.Lock()
	len := this.Queue.Len()
	this.Locker.Unlock()

	return len
}