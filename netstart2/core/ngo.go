package core

import (
	"jsj.golangtc/netstart2/model"
	"github.com/astaxie/beego/logs"
	"fmt"
	"strconv"
	"time"
	"encoding/json"
)

type Ngo struct {
	Ready    chan *ReadyDto
	Accounts []*model.Account
}

func NewNgo() *Ngo {

	ngo := &Ngo{}
	ngo.Ready = make(chan *ReadyDto)

	go func(ngo *Ngo) {
		ngo.scanNet()
	}(ngo)

	go func(ngo *Ngo) {
		ngo.scanDb()
	}(ngo)

	return ngo
}

func (n *Ngo) scanNet() {
	for {

		if (time.Now().Hour() <= 7) {
			time.Sleep(30 * time.Second)
			continue
		}

		account := &model.Account{}
		accounts, err := account.GetAccounts()

		if err != nil {
			logs.Error("查询账户信息出错", err)
			panic(err)
		}

		if len(accounts) == 0 {
			panic("没有找到账户信息")
		}
		n.Accounts = accounts
		account = accounts[0]

		logs.Info(fmt.Sprintf("找账户[%s],数据扫描中………………", account.AccountName))

		act2 := Account{Cookies: account.Cookie, AddressId: strconv.FormatInt(int64(account.AddressId), 10)}
		goods, err := act2.List()

		if err != nil {
			logs.Error("查询网络商品信息出错", err)
			time.Sleep(20 * time.Second)
			continue
		}

		b, _ := json.Marshal(goods)
		logs.Info("扫描结果", string(b))

		err = act2.Save(goods)
		if err != nil {
			logs.Error("保存网络商品信息出错", err)
		}

		time.Sleep(20 * time.Second)
	}
}

func (n *Ngo) scanDb() {

	for {

		logs.Info("数据库扫描中…………")

		product := &model.Product{}
		products, err := product.GetProducts()

		if err != nil {
			logs.Error("查询数据库商品信息出错", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if (len(products) > 0) {
			logs.Info(fmt.Sprintf("找到%d个商品信息", len(products)))
		}
		for _, p := range products {

			for _, account := range n.Accounts {
				go func(p *model.Product, account *model.Account) {
					acc := &Account{
						AddressId:   strconv.FormatInt(int64(account.AddressId), 10),
						Cookies:     account.Cookie,
						AccountName: account.AccountName,
						AccountId:   account.AccountId,
					}
					ready := &ReadyDto{
						Account:      acc,
						BeginTime:    p.BeginTimeUnix,
						GoodsNumbers: p.ProductNumber,
						GoodsName:    p.ProductName,
						GoodsPrice:   p.ProductPrice,
					}
					n.Push(ready)
				}(p, account)
			}
		}
		time.Sleep(1 * time.Second)
	}

}

func (n *Ngo) Run() {
	for {
		ready := <-n.Ready
		Start(ready)
	}
}

func (n *Ngo) Push(ready *ReadyDto) {
	n.Ready <- ready
}
