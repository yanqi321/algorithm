package main
/**
* 几种生成随机字符串方法 a-zA-Z0-9
*/
import (
	"log"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateCode (n int) string {
	b := make([]rune, n)
	letterLen := len(letterRunes)
	log.Println("len", letterLen)
	for i := range b {
		b[i] = letterRunes[rand.Intn(letterLen)]
	}
	return string(b)
}

func main() {
	ret := generateCode(8)
	log.Println(ret)
}