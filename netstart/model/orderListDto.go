package model

type OrderListDto struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data *OrderData `json:"data"`
}

type OrderData struct {
	CurPageNo  int         `json:"curPageNo"`
	TotalPage  int         `json:"totalPage"`
	RecordList []*OrderDto `json:"recordList"`
}

type OrderDto struct {
	OrderID         string  `json:"orderId"`
	OrderStatus     int     `json:"orderStatus"`
	OrderAmount     float64 `json:"orderAmount"`
	OrderCreateTime int64   `json:"orderCreateTime"`
	GoodsNumber     string  `json:"goodsNumber"`
	GoodsCount      int     `json:"goodsCount"`
	GoodsName       string  `json:"goodsName"`
	GoodsHomeImage struct {
		Order   int    `json:"order"`
		IconURL string `json:"iconUrl"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"goodsHomeImage"`
}
