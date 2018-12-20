package model

type GoodDetailDto struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data GoodDetail `json:"data"`
}

type GoodDetail struct {
	MerchantName string      `json:"merchantName"`
	MerchantLogo interface{} `json:"merchantLogo"`
	GoodsNumber  string      `json:"goodsNumber"`
	GoodsName    string      `json:"goodsName"`
	GoodsDesc    string      `json:"goodsDesc"`
	GoodsImages []struct {
		Order   int         `json:"order"`
		IconURL string      `json:"iconUrl"`
		Width   interface{} `json:"width"`
		Height  interface{} `json:"height"`
	} `json:"goodsImages"`
	NbdPrice    float64     `json:"nbdPrice"`
	PreSaleTime int64       `json:"preSaleTime"`
	SaleTime    int64       `json:"saleTime"`
	StockCount  int         `json:"stockCount"`
	ShopTimes   int         `json:"shopTimes"`
	PayTimes    int         `json:"payTimes"`
	GoodsStatus int         `json:"goodsStatus"`
	BuyToken    string      `json:"buyToken"`
	CurrentTime int64       `json:"currentTime"`
	SellerType  int         `json:"sellerType"`
	GoodsExLink interface{} `json:"goodsExLink"`
	Discount    int         `json:"discount"`
}
