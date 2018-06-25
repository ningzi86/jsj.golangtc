package tests

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
	"time"
)

type QueueScheduler struct {
	locker *sync.Mutex
	queue  *list.List
}

func (this *QueueScheduler) Push(message *Message) {
	this.locker.Lock()
	this.queue.PushBack(message)
	this.locker.Unlock()
}

func (this *QueueScheduler) Pull() *Message {

	this.locker.Lock()
	if this.queue.Len() <= 0 {
		this.locker.Unlock()
		return nil
	}

	e := this.queue.Front()
	message := e.Value.(*Message)

	this.queue.Remove(e)
	this.locker.Unlock()

	return message

}

func (this *QueueScheduler) Count() int {

	this.locker.Lock()
	len := this.queue.Len()
	this.locker.Unlock()

	return len
}

type Message struct {
	msgid int
	msg   string
}

func Test_Queue01(t *testing.T) {

	var queue  = &QueueScheduler{locker:new(sync.Mutex), queue:list.New()}
    var c = make(chan bool, 1)

	go func() {

		for {
			
			message := queue.Pull()
			if message != nil {
				fmt.Println(message)
			}

			time.Sleep(5 * time.Second)
		}

	}()

	for index := 0; index <= 100; index++ {

		var message = new(Message)
		message.msgid = index
		message.msg = fmt.Sprintf("这是第%d个消息", index)

		queue.Push(message)

		time.Sleep(5 * time.Second)
	}
    
    <-c
}
