package netstar

import (
	"jsj.golangtc/netstart/model"
	lq "github.com/ahmetalpbalkan/go-linq"
	"github.com/getlantern/errors"
	"fmt"
	"log"
	"time"
	"jsj.golangtc/netstart/utils"
)

type Ngo struct {
	currentTime int64
	ready       chan *model.GoodDetailDto

	flag     chan bool
	stopSync chan bool

	timer *time.Ticker
}

func NewNgo() *Ngo {

	ngo := &Ngo{}
	ngo.ready = make(chan *model.GoodDetailDto)
	ngo.flag = make(chan bool)
	ngo.stopSync = make(chan bool)

	go func() {
		for {
			dto := <-ngo.ready
			log.Printf("开始抢购 %s %s %s", dto.Data.GoodsNumber, dto.Data.GoodsName, dto.Data.BuyToken)

			if Env == "true" {
				fmt.Println("测试环境，模拟抢购成功，停止抢购")
				ngo.flag <- true
				return
			}

			orderId, err := Buy(dto.Data.GoodsNumber, dto.Data.BuyToken)
			if err != nil {

				if err == NetError6 || err == NetError11 || err == NetError7 {
					log.Printf("抢购失败，停止抢购 %s %s %s", err.Error(), dto.Data.GoodsNumber, dto.Data.GoodsName)
					ngo.flag <- true
					return
				}

				ngo.flag <- false
				log.Printf("抢购失败，继续抢购 %s %s %s", err.Error(), dto.Data.GoodsNumber, dto.Data.GoodsName)
				continue
			}

			_, err = Pay(orderId)
			if err != nil {
				ngo.flag <- false
				continue
			}
			log.Printf("抢购成功 %s %s", dto.Data.GoodsNumber, dto.Data.GoodsName)
			ngo.flag <- true
		}
	}()

	return ngo
}

func (Ngo) Init() (*model.GoodListDto, error) {

	goodNumbers, err := List()
	if err != nil {
		return nil, err
	}

	goodListDto, err := ListDetail(goodNumbers)
	return goodListDto, err

}

func (n *Ngo) startSyncTime(currentTime int64) {

	if n.timer != nil {
		n.timer.Stop()
	}

	n.currentTime = currentTime
	n.timer = time.NewTicker(1 * time.Millisecond * 100)

	go func(nn *Ngo) {
		for {
			select {
			case <-nn.stopSync:
				if nn.timer != nil {
					nn.timer.Stop()
				}
				return
			case <-n.timer.C:
				nn.currentTime += 100
				//fmt.Println(nn.currentTime, currentTime, time.Now().UnixNano()/1000000)
			}
		}
	}(n)

}

func (n *Ngo) stopSyncTime() {
	n.stopSync <- true
}

func (Ngo) First(goodsNumber string, goods []model.GoodDto) (model.GoodDto, error) {

	if goods == nil {
		return model.GoodDto{}, errors.New("商品列表为空")
	}

	where := lq.From(goods).Where(func(i interface{}) bool {
		return i.(model.GoodDto).GoodsNumber == goodsNumber
	})

	if where.Count() > 0 {
		return where.First().(model.GoodDto), nil
	}
	return model.GoodDto{}, fmt.Errorf("没有找到商品编号：%s", goodsNumber)

}

func (n *Ngo) Start(dto model.GoodDto) (bool, error) {

	goodsNumber := dto.GoodsNumber
	goodsName := dto.GoodsName

	var secondArrays [][]int32

	for {

		//获取商品详情
		dto, err := Detail(goodsNumber)
		if err != nil {
			log.Fatal("获取商品信息出错", err)
			time.Sleep(50 * time.Millisecond)
			continue
		}
		currentTime := dto.Data.CurrentTime
		saleTime := dto.Data.SaleTime

		leftSeconds := int64(saleTime - currentTime)
		if len(secondArrays) == 0 {
			secondArrays = utils.TimeArrays(int32(leftSeconds / 1000))
		}

		if leftSeconds <= 0 {

			log.Printf("准备抢购商品：%s %s", goodsNumber, goodsName)
			n.ready <- dto
			success := <-n.flag
			if success {
				return true, nil
			}

			if Env == "true" {
				fmt.Println("测试环境，退出抢购")
				break
			}

			continue

		} else {

			arrays := utils.CalTimeArrays(int32(leftSeconds/1000), secondArrays)
			min := arrays[0] / 2

			fmt.Println("分割时间", arrays)
			fmt.Println("当前使用", secondArrays)
			fmt.Printf("剩余时间：%d  超时：%d\n", leftSeconds/1000, min)

			if leftSeconds <= 30*1000 {

				//n.startSyncTime(currentTime)
				//for {
				//	if (n.currentTime < saleTime) {
				//		time.Sleep(50 * time.Millisecond)
				//		log.Printf("当前时间：%s %s 开售时间：%s 准备抢购 %s %s 系统时间：%s\n",
				//			utils.TimeFormat(currentTime/1000),
				//			utils.TimeFormat(n.currentTime/1000),
				//			utils.TimeFormat(saleTime/1000), goodsNumber, goodsName,
				//			utils.TimeFormat(time.Now().Unix()))
				//		continue
				//	}
				//	break
				//}
				//n.stopSyncTime()

				log.Printf("准备抢购：%s %s", goodsNumber, goodsName)

				n.ready <- dto
				success := <-n.flag
				if success {
					return true, nil
				}

				continue
			} else {

				log.Printf("当前时间：%s %s 开售时间：%s 剩余：%d秒 未开始 %s %s\n",
					utils.TimeFormat(currentTime/1000),
					utils.TimeFormat(n.currentTime/1000),
					utils.TimeFormat(saleTime/1000),
					(saleTime-currentTime)/1000,
					goodsNumber,
					goodsName)

				log.Printf("【%d-%d】%d秒后重试……\n", arrays[0], arrays[1], min)
				<-time.After(time.Second * time.Duration(min))

				continue
			}
		}
	}

	return false, errors.New("未知错误")
}

func (Ngo) NOrders() (*model.OrderListDto, error) {
	dto, err := Orders(1, 20)
	return dto, err
}
