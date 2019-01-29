package core

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"jsj.golangtc/netstart2/netstar"
	"jsj.golangtc/netstart2/model"
	"time"
)

func Start(ready *ReadyDto) {

	logs.Info(fmt.Sprintf("[%s]发现商品“%s_%s_%s”……", ready.Account.AccountName, ready.GoodsNumbers, ready.GoodsName, ready.GoodsPrice))

	go func(ready *ReadyDto) {
		for {

			if ready.RetryTimes >= 10 {
				logs.Warn("抢购失败，超出重试次数")
				break
			}

			ns := &netstar.NetStar{
				AddressId: ready.Account.AddressId,
				Cookie:    ready.Account.Cookies,
			}
			//获取商品详情
			dto, err := ns.Detail(ready.GoodsNumbers)
			//
			//fmt.Println(ready.GoodsNumbers, dto.Data.BuyToken)
			//break

			orderId, err := ns.Buy(ready.GoodsNumbers, dto.Data.BuyToken)
			if err != nil {

				if err == netstar.NetError6 || err == netstar.NetError11 || err == netstar.NetError7 {
					memo := fmt.Sprintf("抢购失败，停止抢购 %s %s %s", err.Error(), dto.Data.GoodsNumber, dto.Data.GoodsName)
					logs.Warn(memo)
					addProductLogs(ready.Account.AccountId, ready.GoodsNumbers, ready.GoodsName, memo)
					return
				}

				ready.RetryTimes ++
				memo := fmt.Sprintf("抢购失败，继续抢购 %s %s %s", err.Error(), dto.Data.GoodsNumber, dto.Data.GoodsName)
				logs.Warn(memo)
				addProductLogs(ready.Account.AccountId, ready.GoodsNumbers, ready.GoodsName, memo)
				continue
			}

			_, err = ns.Pay(orderId)
			if err != nil {
				ready.RetryTimes ++
				continue
			}

			memo := fmt.Sprintf("抢购成功 %s %s", dto.Data.GoodsNumber, dto.Data.GoodsName)
			logs.Info(memo)
			addProductLogs(ready.Account.AccountId, ready.GoodsNumbers, ready.GoodsName, memo)

			break
		}
	}(ready)

}

func addProductLogs(accountId int, goodsNumber string, goodsName string, memo string) {

	log := model.ProductLogs{
		AccountId:     accountId,
		ProductNumber: goodsNumber,
		ProductName:   goodsName,
		Memo:          memo,
		CreateTime:    int(time.Now().Unix()),
	}
	log.Insert()
}
