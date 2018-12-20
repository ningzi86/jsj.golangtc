package netstar

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestJson(t *testing.T) {

	//j := `{"code":200,"msg":"请求成功","data":{"result":7,"orderId":null}}`

	j := `{"code":200,"msg":"请求成功","data":{"result":1,"payAmount":0.0100}}`

	var result map[string]interface{}
	json.Unmarshal([]byte(j), &result)

	fmt.Println(result)
	fmt.Println(result["data"].(map[string]interface{})["result"])
	fmt.Println(result["data"].(map[string]interface{})["payAmount"].(float64))

}

func TestBuy(t *testing.T) {
	//buyToken=72cabb1b-4311-48c7-ae89-d3e0041e4917903103
	//{"goodsNumber":"120001541578607296840660","addressId":34743,"buyToken":"081b1a69-5e27-4e07-b624-e6c2a0099d60275537"}
	//buyToken=28d1cc3c-ed4b-4179-9b8e-28de46995bdc66776

	orderId, err := Buy("700101544587219758209409", "28d1cc3c-ed4b-4179-9b8e-28de46995bdc66776")
	fmt.Println(orderId, err)

}

func TestList(t *testing.T) {
	list, err := List()
	fmt.Println(list, err)

	dto, err := ListDetail(list)
	bytes, err := json.Marshal(dto)

	fmt.Println(err)
	fmt.Printf("%s", bytes)
}
