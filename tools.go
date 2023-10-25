package ssdata

import (
	"strings"
	"time"
)

func ReverseString(str string) string {

	d := strings.Split(str, "")

	a := []string{}

	for i := range d {
		a = append(a, d[len(d)-1-i])
	}
	return strings.Join(a, "")
}

// 中文有问题。。。
func ReverseString11(str string) string {
	// 将字符串转换为字节切片
	byteSlice := []byte(str)
	length := len(byteSlice)

	// 使用双指针进行字节切片的反转
	for i := 0; i < length/2; i++ {
		byteSlice[i], byteSlice[length-i-1] = byteSlice[length-i-1], byteSlice[i]
	}

	// 将字节切片转换为字符串并返回
	return string(byteSlice)
}

func Includes(arr []string, value string) bool {

	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false

}

func now() string {
	tn := time.Now()
	location, _ := time.LoadLocation("Asia/Shanghai")
	t := tn.In(location).Format("2006-01-02 15:04:05")
	return t
}

var codemap = [][]string{
	{"A", "龘"},
	// {"B", "鬻"},
	{"C", "鱻"},
	// {"D", "龖"},
	// {"E", "鱲"},
	// {"F", "鑴"},
	// {"G", "龕"},
	// {"H", "龗"},
	// {"I", "爨"},
	// {"J", "爚"},
	// {"K", "蠃"},
	// {"L", "鑾"},
	// {"M", "韙"},
	// {"N", "齾"},
	{"O", "鰙"},
	// {"P", "鸺"},
	// {"Q", "蠣"},
	// {"R", "囃"},
	// {"S", "鏻"},
	// {"T", "鸄"},
	{"U", "韡"},
	// {"V", "鑹"},
	// {"W", "鱽"},
	// {"X", "鑷"},
	{"Y", "鏃"},
	// {"Z", "齨"},
	// --
	{"s", "蜑"},
	// {"d", "硐"},
	// {"t", "肽"},
	{"w", "魍"},
}
