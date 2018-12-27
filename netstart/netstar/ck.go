package netstar

import (
	"io/ioutil"
	"strconv"
	"os"
	"log"
)

var NetCookie string
var AddressId int32

func init() {

	c, err := GetCookie()
	if err != nil {
		log.Fatal(err)
	}
	NetCookie = c

	a, err := GetAddress()
	if err != nil {
		log.Fatal(err)
	}

	AddressId = a

	envName := "生产环境"
	if Env == "true" {
		envName = "测试环境"
	}

	log.Println("初始化成功", NetCookie, AddressId, envName)

}

func GetCookie() (string, error) {

	pwd, _ := os.Getwd()
	path := pwd + `/config/cookies.txt`
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func GetAddress() (int32, error) {

	pwd, _ := os.Getwd()
	path := pwd + `/config/address.txt`
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return -1, err
	}

	address, err := strconv.Atoi(string(bytes))
	if err != nil {
		return -1, err
	}
	return int32(address), nil
}
