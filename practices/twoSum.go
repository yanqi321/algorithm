package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    in := bufio.NewReader(os.Stdin)   
    str, err := in.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    str = strings.Replace(str, "\n", "", -1)
    arr := strings.Split(str, " ")
    ans := twoSum(arr[0], arr[1])
    fmt.Println(ans)
}
func twoSum(a, b string) string {
    prev := 0
    ans := ""
    aLen, bLen := len(a), len(b)
    for i := 0;i < len(a) || i <len(b) || prev > 0;i++ {
        curr := prev
        if  i < len(a) {
            val, e := strconv.Atoi(string(a[aLen -1 - i]))
            if e != nil {
                fmt.Println(e)
            }
            curr += val
        }
        if  i < len(b) {
            val, e := strconv.Atoi(string(b[bLen -1 - i]))
            if e != nil {
                fmt.Println(e)
            }
            curr += val
        }
        prev = curr / 10
        curr = curr % 10
        ans = strconv.Itoa(curr) + ans
    }
    return ans
}
