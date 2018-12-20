package model

type GoodListDto struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		GoodsDetails []GoodDto `json:"goodsDetails"`
		Datetime     int       `json:"datetime"`
	} `json:"data"`
}
