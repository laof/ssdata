package main

import (
	"log"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	data, err := Get("https://laof.github.io/get-nodes-test-app/json/data.json")

	if err != nil {
		return
	}
	mapdata := PingAll(data, true)

	if len(mapdata) > 0 {
		log.Println("ok ==" + strconv.Itoa(len(mapdata)))
	} else {
		t.Errorf("fdafda")
	}

}
