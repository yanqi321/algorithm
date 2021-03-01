package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

const base = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/"

func main()  {
	for {
		in := bufio.NewReader(os.Stdin)   
		str, err := in.ReadString('\n')
		if err != nil {
				fmt.Println(err)
		}
		decodeStr := base64Decode(str)

		ans := encodeBase36(decodeStr)
		fmt.Println(ans)
	}
}

func encodeBase36(base64str []byte) string {
	N := len(base64str)
	prev := 0
	ans := ""
	for i:=N-1; i >=0; i-- {
		val := int(base64str[i]) 
		if prev > 0 {
			val = (prev + int(base64str[i]) * 256) % 36
			prev = (prev + int(base64str[i]) * 256 - val) / 256
		} else {
			val = int(base64str[i]) % 36
			prev = (int(base64str[i]) * 256 - val) / 256
		}
		for prev > 36 {
			val := prev % 36
			ans = strconv.Itoa(val) + ans
			prev = prev / 36
		}
		fmt.Println(ans, prev, val)
		ans = string(base[val]) + ans
		// num += int64(base64str[i]) * int64(math.Pow(256, float64(N-i-1)))
	}
	// fmt.Println(base64str, num)
	// ans := ""
	// for num > 0 {
	// 	i := num % 36
	// 	ans = string(base[i]) + ans
	// 	num = (num - i) / 36
	// }
	return ans
}

func base64Decode(str string) []byte {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(err)
	}
	return decodeBytes
}