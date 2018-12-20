package netstar

import (
	"testing"
	"fmt"
	"time"
)

func TestNgo_time(t *testing.T) {

	ngo := NewNgo()
	ngo.startSyncTime(100)

	for {
		fmt.Println(ngo.currentTime)
		time.Sleep(100)
	}
}

func TestNgo_time2(t *testing.T) {

	nano := time.Now().Unix()
	fmt.Println(nano)
}
