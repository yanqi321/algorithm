package main

import (
	"fmt"
	"math"
)

const base = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" // 这个顺序可以随机打乱用来混淆,这个应该是base62与62进制的区别

func Base62encode(num int) string {
	ans := ""
	for num > 0 {
		i := num % 62
		ans = string(base[i]) + ans
		num = (num - i) / 62
	}
	// for len(ans) < 8 {
	// 	ans = string(base[0]) + ans
	// }
	return ans
}

func Base62decode(base62 string) int {
	ans := 0
	N := len(base62)
	f := flip(base)
	for i:=N-1; i >=0; i-- {
		ans += f[base62[i]] * int(math.Pow(62, float64(N-i-1)))
	}
	return ans
}

func flip(s string) map[byte]int {
	f := make(map[byte]int)
	for i := range s {
		f[s[i]] = i
	}
	return f
}
func getID(num int) string {
	num += 1 << 47 // 增加一个起始增量，这样能够保证为八位(62进制第一位非0),
	m := Base62encode(num)
	return m
}
func main() {
	id := getID(1) // 每次获取手动传一个或者redis自增ID
	id2 := getID(1e12) // 10000亿个多ID可以用, 缺点是相邻ID 相似，容易被猜出来，有没有算法可以混淆一下～
	fmt.Println(id, id2) // n7vQqkSF oPWycp2I
}