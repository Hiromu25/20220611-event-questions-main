package main

import (
	"fmt"
	"strconv"
)
func main(){
    for i := 1; i <= 100; i++ {
        str := ""

        // ここから記述
        
        if (i % 15 == 0){
          str = "FizzBuzz"
        }else if (i % 5 == 0){
          str = "Buzz"
        }else if (i % 3 == 0){
          str = "Fizz"
        }else {
          str = strconv.Itoa(i)
        }

        // ここまで記述

    fmt.Println(str)
  }
}
