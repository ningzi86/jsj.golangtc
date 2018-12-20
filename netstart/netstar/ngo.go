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

	n.currentTime = currentTime
	go func(nn *Ngo) {
		for {
			select {
			case <-nn.stopSync:
				nn.currentTime = time.Now().Unix()
				return
			case <-time.After(1 * time.Second):
				nn.currentTime ++
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

	for {

		//获取商品详情
		dto, err := Detail(goodsNumber)
		if err != nil {
			log.Fatal("获取商品信息出错", err)
			time.Sleep(50 * time.Millisecond)
			continue
		}
		currentTime := dto.Data.CurrentTime / 1000
		n.startSyncTime(currentTime)

		saleTime := dto.Data.SaleTime / 1000

		if currentTime >= saleTime {
			//if currentTime <= (saleTime + 30) {

			log.Printf("准备抢购商品：%s %s", goodsNumber, goodsName)

			n.ready <- dto
			success := <-n.flag
			if success {
				return true, nil
			}
			n.stopSyncTime()
			continue
			//}
			//return false, fmt.Errorf("当前时间：%s %s 开售时间：%s 超时购买失败：%s %s",
			//	utils.TimeFormat(currentTime),
			//	utils.TimeFormat(n.currentTime),
			//	utils.TimeFormat(saleTime),
			//	goodsNumber,
			//	goodsName)
		}

		if currentTime < saleTime {
			if currentTime >= saleTime-10 {
				for {
					if (n.currentTime < saleTime) {
						time.Sleep(50 * time.Millisecond)
						log.Printf("当前时间：%s %s 开售时间：%s 准备抢购 %s %s\n",
							utils.TimeFormat(currentTime),
							utils.TimeFormat(n.currentTime),
							utils.TimeFormat(saleTime), goodsNumber, goodsName)
						continue
					}
					break
				}
				log.Printf("准备抢购：%s %s", goodsNumber, goodsName)
				n.ready <- dto
				success := <-n.flag
				if success {
					return true, nil
				}
				continue
			}
			log.Printf("当前时间：%s %s 开售时间：%s 未开始 %s %s\n",
				utils.TimeFormat(currentTime),
				utils.TimeFormat(n.currentTime),
				utils.TimeFormat(saleTime), goodsNumber, goodsName)
			time.Sleep(1 * time.Second)
			n.stopSyncTime()
			continue
		}

	}

	return false, errors.New("未知错误")
}
