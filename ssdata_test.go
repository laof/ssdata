package ssdata

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	data, err := os.ReadFile("test.txt")

	if err != nil {
		log.Println(err.Error())
		return
	}

	txt := string(data)

	list := strings.Split(txt, "\n")

	ss := FilterSlice[string](list, func(i int, val string) bool {
		return strings.HasPrefix(val, "ss://")
	})

	Test(strings.Join(ss, "\n"))

}
