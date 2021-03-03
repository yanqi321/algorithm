package main

import (
	"bufio"
	"fmt"
    "sort"
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
    nums, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
    }
    for i:=0; i < nums;i++ {
        str, err := in.ReadString('\n')
        if err != nil {
            fmt.Println(err)
        }
        str, err = in.ReadString('\n')
        if err != nil {
            fmt.Println(err)
        }
        str = strings.Replace(str, "\n", "", -1)
        arr := strings.Split(str, " ")
        newArr := make([]int, 0)
        sum := 0
        for _, s := range arr {
            v, err := strconv.Atoi(s)
            if err != nil {
                fmt.Println(err)
                continue
            }
            v = v%3
            if v == 0 {
                sum++
                continue
            }
            newArr = append(newArr, v)
        }
        sort.Ints(newArr)
        ans := ArrSum(newArr) + sum
        fmt.Println(ans)
    }
    
}
// ArrSum 计算最多个数合并被3 整除
func ArrSum(arr []int) int {
    if len(arr) == 1 {
        return 0
    }
    ans :=0
    i,j := 0,len(arr) -1;

    for i <= j{
        if arr[i] + arr[j] == 3 {
            ans++
        } else {
           break
        }
        i++
        j--
    }
    if i < j {
        return ans + (j -i +1) / 3
    }
    return ans
}
