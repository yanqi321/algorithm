package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    for {
        fmt.Println("How to read all lines here?")
        in := bufio.NewReader(os.Stdin)   
				result, err := in.ReadString('\n')
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println("\nresult", result)
    }
}