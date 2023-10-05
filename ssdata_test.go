package ssdata

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	data, err := Get("https://laof.github.io/get-nodes-test-app/json/data.json")

	if err != nil {
		return
	}
	mapdata := PingAll(data, 3)

	fmt.Println(mapdata)

}

func TestReverseString(t *testing.T) {
	aaa := ReverseString("你")
	fmt.Println(aaa)
}
