package main

import (
	"fmt"
    "math/big"
)

func main() {
    for {
        var num int
        _, err := fmt.Scanf("%d", &num)
        if err != nil {
            fmt.Println(err)
        }
        if num == 0 {
            break
        }
        ans := Fibonacci(num)
        fmt.Println(ans)
    }
    
}
// Fibonacci 计算斐波那契数列
func Fibonacci(n int) string {
    if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	}
	a := big.NewInt(0)
	b := big.NewInt(1)

	
	for i := 0;i < n;i++ {
		// 计算下一个斐波那契数列
		a.Add(a, b)
		// 交换a和b，使a保留计算后的结果
		a, b = b, a
	}
	return a.String()
}
