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
    cfg := strings.Split(str, " ")
    n, err1 := strconv.Atoi(cfg[0])
    m, err2 := strconv.Atoi(cfg[1])
    if err1 != nil || err2 != nil {
        fmt.Println(err1, err2)
    }
    str2, err := in.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    str2 = strings.Replace(str2, "\n", "", -1)
    strArr := strings.Split(str2, " ")
    arr := make([]int, len(strArr))
    for i, v := range strArr {
        arr[i],err = strconv.Atoi(v)
        if err != nil {
            fmt.Println(err)
        }
    }
    ans := getWater(n, m , arr) + 1
    fmt.Println(ans)
}
func getWater(n, m int, arr []int) int {
    if n < m {
        return 0
    }
    idx := getMin(arr[:m])
    delVal := arr[idx]
    newItem := make([]int,0)
    for i:=0; i < m;i++ {
        arr[i] -= delVal
        if arr[i] == 0 {
            n--
            continue
        }
        newItem = append(newItem, arr[i])
    }
    arr = append(newItem, arr[m:]...)
    return delVal + getWater(n, m, arr)
}
func getMin(arr []int) int {
    idx, min := 0, arr[0]
    for i, v := range arr {
        if v < min {
            min = v
            idx = i
        }
    }
    return idx
}
