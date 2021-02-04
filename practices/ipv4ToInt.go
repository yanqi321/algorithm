package main

import (
	"fmt"
	"strings"
	"strconv"
)
// ipV4ToInt ip转为数字
func ipV4ToInt(ipStr string) int {
	s := strings.Split(ipStr, ".")
	ret := 0
	for i := 3; i >= 0; i-- {
			v, err := strconv.Atoi(s[i])
			if err != nil {
				return ret
			}
			ret += v << (8 * (3 - i))
			fmt.Println("ret", ret)
	}
	return ret
}

func intToIpV4(n int) string {
	ret := ""
	for i:= 0; i < 4; i++ {
		val := n >> ((3-i) * 8) & 255 // 清除高位数据和首位符号
		ret += strconv.Itoa(val) + "."
	}
	return ret[:len(ret) - 1]
}

func main() {
	ret := ipV4ToInt("192.168.1.2")
	fmt.Println("ipV4ToInt", ret)
	rawIp := intToIpV4(3232235778)
	fmt.Println("intToIpV4", rawIp)
}