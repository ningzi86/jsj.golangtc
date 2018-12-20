package model

type GoodDto struct {
	GoodsNumber      string `json:"goodsNumber"`
	GoodsName        string `json:"goodsName"`
	GoodsDescription string `json:"goodsDescription"`
	GoodsImages []struct {
		Order   int    `json:"order"`
		IconURL string `json:"iconUrl"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
	} `json:"goodsImages"`
	GoodsNbdPrice float64     `json:"goodsNbdPrice"`
	GoodsStatus   int         `json:"goodsStatus"`
	Type          int         `json:"type"`
	Score         int         `json:"score"`
	PreSaleTime   int64       `json:"preSaleTime"`
	SaleTime      int64       `json:"saleTime"`
	Tag           interface{} `json:"tag"`
	AlwaysShow    int         `json:"alwaysShow"`
}
