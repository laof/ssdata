package ssdata

func reverseString(str string) string {
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

func includes(arr []string, value string) bool {

	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false

}

var CodeMap = [][]string{
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
