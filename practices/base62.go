package main

import (
	"fmt"
	"math"
)

const base = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" // 这个顺序可以随机打乱用来混淆,这个应该是与62进制的区别

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

func main() {
	m := Base62encode(5202340)
	fmt.Println(m)
	fmt.Println(Base62decode("BAAAAAAA"))
}