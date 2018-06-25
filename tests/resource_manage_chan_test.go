package tests

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
	"time"

	"jsj.golangtc/models"
)

func Test_Resourcechan01(t *testing.T) {

	var queue = &models.QueueScheduler{Locker: new(sync.Mutex), Queue: list.New()}
	for index := 0; index <= 100; index++ {

		var message = new(models.Message)
		message.Msgid = index
		message.Msg = fmt.Sprintf("这是第%d个消息", index)

		queue.Push(message)
	}

	mc := models.NewResourceManageChan(10000)

	for {

		message := queue.Pull()

		if message == nil {

			if mc.Has() == 0 {
				fmt.Println("所有程序处理完成，程序退出")
				break
			}

			//等待一秒钟后继续执行
			time.Sleep(1 * time.Second)
			continue
		}

		mc.GetOne()

		go func() {

			fmt.Println("处理消息", message)

			time.Sleep(3 * time.Second)
			mc.FreeOne()

		}()
	}

	fmt.Println("所有程序执行完成")

}
