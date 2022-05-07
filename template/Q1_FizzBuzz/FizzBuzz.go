package main

import (
	"fmt"
	"strconv"
)
func main(){
    for i := 1; i <= 100; i++ {
        str := ""

        // ここから記述

        if i % 15 == 0 {
          // 15で割り切れる場合
          str = "FizzBuzz"
        } else if i % 5 == 0 {
          // 5で割り切れる場合
          str = "Buzz"
        } else if i % 3 == 0 {
          // 3で割り切れる場合
          str = "Fizz"
        } else {
          // それ以外の数の場合
          str = strconv.Itoa(i)
        }

        // ここまで記述

    fmt.Println(str)
  }
}
