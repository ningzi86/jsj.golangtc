package netstar

import (
	"testing"
	"fmt"
)

func TestGetCookie(t *testing.T) {

	cookie, err := GetCookie()
	fmt.Println(cookie, err)

}
