package tests

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
	"time"
	"jsj.golangtc/models"
)




func Test_Queue01(t *testing.T) {

	 var queue  = &models.QueueScheduler{Locker:new(sync.Mutex), Queue:list.New()}
    // var c = make(chan bool, 1)

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

		var message = new(models.Message)
		message.Msgid = index
		message.Msg = fmt.Sprintf("这是第%d个消息", index)

		queue.Push(message)

		time.Sleep(5 * time.Second)
	}
    
	// go func ()  {
	// 	for{
	// 		fmt.Println("newline")
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

    // <-c
}
