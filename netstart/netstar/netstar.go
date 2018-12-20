package netstar

import (
	"fmt"
	"jsj.golangtc/netstart/core"
	"encoding/json"
	"errors"
	"jsj.golangtc/netstart/model"
)

func List() ([]string, error) {

	url := `https://star.8.163.com/api/goods/home/v2/list`
	headers := `
		Host: star.8.163.com
		Accept: application/json, text/plain, */*
		Origin: https://star.8.163.com
		X-Requested-With: XMLHttpRequest
		User-Agent: Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "ab09afc1788e2737ac8d7c464dcb7f2cf2f69a8e4b1c0dd49c11759348539bb9",channel: "e01170001"}star_client_info_end
		Referer: https://star.8.163.com/m
		Accept-Encoding: gzip, deflate
		Accept-Language: zh-CN,en-US;q=0.9
		Cookie: %s`

	headers = fmt.Sprintf(headers, NetCookie)

	body := ""
	//fmt.Printf("请求数据：%s\n", body)

	client := core.NewNetClient(url, "POST", headers, string(body), nil, nil, 0)
	err := client.Do()

	//fmt.Printf("响应结果：%s\n", client.ResponseBody)

	//
	//{
	//	"code": 200,
	//	"msg": "请求成功",
	//	"data": {
	//"list": [
	//"120001543482888051275234",
	//"120101544605530401789046",
	//"120001544436059877923708",
	//"120001541578607296840660",
	//
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(client.ResponseBody), &result)

	code := result["code"].(float64)
	msg := result["msg"].(string)

	if code != 200 {
		return nil, errors.New(msg)
	}

	data := result["data"]
	if data == nil {
		return nil, errors.New("未知错误")
	}

	resultMap := data.(map[string]interface{})
	list := resultMap["list"].([]interface{})

	var goodsNumbers []string
	for _, v := range list {
		goodsNumbers = append(goodsNumbers, v.(string))
	}

	return goodsNumbers, nil

}

func ListDetail(goodsNumbers []string) (*model.GoodListDto, error) {

	url := `https://star.8.163.com/api/goods/home/v2/detail`

	headers := `
			Host: star.8.163.com
			Accept: application/json, text/plain, */*
			Origin: https://star.8.163.com
			X-Requested-With: XMLHttpRequest
			User-Agent: Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "a1d823382b7f23ed568eedc68c556908b2f5fbe29a13313c1d052a66742fbd65",channel: "e01170001"}star_client_info_end
			Content-Type: application/json;charset=UTF-8
			Referer: https://star.8.163.com/m
			Accept-Encoding: gzip, deflate
			Accept-Language: zh-CN,en-US;q=0.9
			Cookie: %s`

	headers = fmt.Sprintf(headers, NetCookie)

	mp := map[string]interface{}{
		"goodsNumbers": goodsNumbers,
	}

	body, _ := json.Marshal(mp)

	//fmt.Println(headers)
	//fmt.Printf("请求数据：%s\n", body)

	client := core.NewNetClient(url, "POST", headers, string(body), nil, nil, 0)
	err := client.Do()

	//fmt.Printf("响应结果：%s\n", client.ResponseBody)

	if err != nil {
		return nil, err
	}

	goodListDto := &model.GoodListDto{}
	err = json.Unmarshal([]byte(client.ResponseBody), &goodListDto)

	if err != nil {
		return nil, err
	}

	if goodListDto.Code != 200 {
		return nil, errors.New(goodListDto.Msg)
	}

	return goodListDto, nil
	//{"goodsNumbers":["120001543482888051275234","120101544605530401789046","120001544436059877923708","120001541578607296840660","900101543544606481860051","300101544085767924292090","700101544518262243443720","130101544580129696267130","700101544587219758209409","700101544587887281610208","600101544660286442062312","600101544660434750309349","600101544660591556430172","130101544667910263598661","900101544597178570415219","300101544752663947662267","600101544496624396233074","130101544667653418868807","700101544668085492593626","700101544672581402383601"]}

}

func Detail(goodsNumber string) (*model.GoodDetailDto, error) {

	//currentTime := time.Now().Unix() * 1000
	//saleTime := int64(1545287313000 + 30*1000)
	//
	//m := &model.GoodDetailDto{}
	//m.Data = model.GoodDetail{
	//	CurrentTime: currentTime,
	//	SaleTime:    saleTime,
	//}
	//return m, nil

	url := `https://star.8.163.com/api/goods/detail`

	headers := `
			Host: star.8.163.com
			Accept: application/json, text/plain, */*
			Origin: https://star.8.163.com
			X-Requested-With: XMLHttpRequest
			User-Agent: Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "a1d823382b7f23ed568eedc68c556908b2f5fbe29a13313c1d052a66742fbd65",channel: "e01170001"}star_client_info_end
			Content-Type: application/json;charset=UTF-8
			Referer: https://star.8.163.com/m
			Accept-Encoding: gzip, deflate
			Accept-Language: zh-CN,en-US;q=0.9
			Cookie: %s`

	headers = fmt.Sprintf(headers, NetCookie)

	mp := map[string]interface{}{
		"goodsNumber": goodsNumber,
	}

	body, _ := json.Marshal(mp)

	//fmt.Println(headers)
	//fmt.Printf("请求数据：%s\n", body)

	client := core.NewNetClient(url, "POST", headers, string(body), nil, nil, 0)
	err := client.Do()

	//fmt.Printf("响应结果：%s\n", client.ResponseBody)

	if err != nil {
		return nil, err
	}

	dto := &model.GoodDetailDto{}
	err = json.Unmarshal([]byte(client.ResponseBody), &dto)

	if err != nil {
		return nil, err
	}

	if dto.Code != 200 {
		return nil, errors.New(dto.Msg)
	}

	return dto, nil

}

var NetError6 = errors.New("原力不足")
var NetError11 = errors.New("今日达到上限")
var NetError7 = errors.New("商品已售罄")

func Buy(goodsNumber, buyToken string) (string, error) {

	//Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "a1d823382b7f23ed568eedc68c556908b2f5fbe29a13313c1d052a66742fbd65",channel: "e01170001"}star_client_info_end
	//Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion
	url := `https://star.8.163.com/api/order/buy`
	//url := `http://localhost/JSJTickets.Web/_Logs/Default2.aspx`
	headers := `
			Host: star.8.163.com
			Accept: application/json, text/plain, */*
			Origin: https://star.8.163.com
			X-Requested-With: XMLHttpRequest
			User-Agent: Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "a1d823382b7f23ed568eedc68c556908b2f5fbe29a13313c1d052a66742fbd65",channel: "e01170001"}star_client_info_end
			Content-Type: application/json;charset=UTF-8
			Referer: https://star.8.163.com/m
			Accept-Encoding: gzip, deflate
			Accept-Language: zh-CN,en-US;q=0.9
			Cookie: %s`
	//
	headers = fmt.Sprintf(headers, NetCookie)

	//headers := `Cookie: STAREIG=c8d7e0816ea86e47861ca4eb39c809c2345d437e; NTES_YD_SESS=VbsoxtV31Dfpv5mbwI3kMLI4TKpSmmr.XG6DhldVqDKPzSAkzfcxNLfNW52f.Cvt3DHRrMYSBAmUH1BvUphe7e53qnU2uA0LoLYtTf50ZKIkGkcEgbWw.6UjqfCGFArSUk4sM3JJIhcURe1q4uF4Mp5XgKt_mhGJkkj8Fxy8_bKBxmZR812yE4NlKicjK5E9.1vrRUrJ0anUMSEtOR.Y8IT_mwj.B7.R6OBPPFinS9W.W; STAR_YD_SESS=VbsoxtV31Dfpv5mbwI3kMLI4TKpSmmr.XG6DhldVqDKPzSAkzfcxNLfNW52f.Cvt3DHRrMYSBAmUH1BvUphe7e53qnU2uA0LoLYtTf50ZKIkGkcEgbWw.6UjqfCGFArSUk4sM3JJIhcURe1q4uF4Mp5XgKt_mhGJkkj8Fxy8_bKBxmZR812yE4NlKicjK5E9.1vrRUrJ0anUMSEtOR.Y8IT_mwj.B7.R6OBPPFinS9W.W
	//Accept: application/json, text/plain, */*
	//Content-Type: application/json;charset=UTF-8
	//Host: star.8.163.com
	//Content-Length: 116`

	mp := map[string]interface{}{
		"goodsNumber": goodsNumber,
		"addressId":   AddressId,
		"buyToken":    buyToken,
	}

	body, _ := json.Marshal(mp)

	//fmt.Println(headers)
	fmt.Printf("请求数据：%s\n", body)

	client := core.NewNetClient(url, "POST", headers, string(body), nil, nil, 0)
	err := client.Do()

	fmt.Printf("响应结果：%s\n", client.ResponseBody)

	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(client.ResponseBody), &result)

	code := result["code"].(float64)
	msg := result["msg"].(string)

	if code != 200 {
		return "", errors.New(msg)
	}

	data := result["data"]
	if data == nil {
		return "", errors.New("未知错误")
	}

	resultMap := data.(map[string]interface{})
	resultCode := resultMap["result"].(float64)

	if resultCode == 6 {
		return "", NetError6
	}

	if resultCode == 11 {
		return "", NetError11
	}

	if resultCode == 7 {
		return "", NetError7
	}

	if resultCode != 1 {
		return "", fmt.Errorf("未知错误,错误码:%f", resultCode)
	}

	orderId := resultMap["orderId"]
	if orderId == nil {
		return "", fmt.Errorf("未知错误,没有获取到订单编号")
	}

	return orderId.(string), nil

	//{"code":-24,"msg":"订单重复提交","data":null}
	//{"code":200,"msg":"请求成功","data":{"result":7,"orderId":null}}
	//{"code":200,"msg":"请求成功","data":{"result":6,"orderId":null}} 原力不足
	//{"code":200,"msg":"请求成功","data":{"result":5,"orderId":null}} 还未开始
}

func Pay(orderId string) (float64, error) {

	url := `https://star.8.163.com/api/order/pay`

	headers := `
			Host: star.8.163.com
			Accept: application/json, text/plain, */*
			Origin: https://star.8.163.com
			X-Requested-With: XMLHttpRequest
			User-Agent: Mozilla/5.0 (Linux; Android 9; EML-AL00 Build/HUAWEIEML-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hybrid/1.0.0 star_client_info_begin {hybridVersion: "1.0.0",clientVersion: "1.8.1",accountId: "a1d823382b7f23ed568eedc68c556908b2f5fbe29a13313c1d052a66742fbd65",channel: "e01170001"}star_client_info_end
			Content-Type: application/json;charset=UTF-8
			Referer: https://star.8.163.com/m
			Accept-Encoding: gzip, deflate
			Accept-Language: zh-CN,en-US;q=0.9
			Cookie: %s`
	//
	headers = fmt.Sprintf(headers, NetCookie)

	mp := map[string]interface{}{
		"orderId": orderId,
	}

	body, _ := json.Marshal(mp)

	//fmt.Println(headers)
	//fmt.Printf("请求数据：%s\n", body)

	client := core.NewNetClient(url, "POST", headers, string(body), nil, nil, 0)
	err := client.Do()

	//fmt.Printf("响应结果：%s\n", client.ResponseBody)

	if err != nil {
		return -1, err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(client.ResponseBody), &result)

	code := result["code"].(float64)
	msg := result["msg"].(string)

	//{"code":200,"msg":"请求成功","data":{"result":1,"payAmount":0.0100}}
	if code != 200 {
		return -1, errors.New(msg)
	}

	data := result["data"]
	if data == nil {
		return -1, errors.New("未知错误")
	}

	resultMap := data.(map[string]interface{})
	resultCode := resultMap["result"].(float64)

	if resultCode != 1 {
		return -1, fmt.Errorf("未知错误,错误码:%f", resultCode)
	}

	return resultMap["payAmount"].(float64), nil

}
