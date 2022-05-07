package main

import (
	"fmt"
)
func main(){
    // 昇順にソートされた配列
    sortedArray := []int{1, 2, 3, 5, 12, 7890, 12345}
    // 探索対象の番号
    targetNumber := 7890
    // 探索実行
    targetIndex := serchIndex(sortedArray, targetNumber)
    // 結果出力
    fmt.Println(targetIndex)
}

func serchIndex(sortedArray []int, targetNumber int) int {

    // ここから記述

    // 探索対象配列のindexを変数で指定
    searchStart := 0
    searchEnd := len(sortedArray) - 1
    for {
        if searchStart == searchEnd {
            // 配列の長さが1のとき
            if sortedArray[searchStart] == targetNumber {
                return searchStart
            } else {
                break
            }
        }else {
            // 探索対象配列の中間のindexを取得
            // 配列の長さが偶数長の場合，前半の長さ = 後半の長さ+1になるようにする
            // ex) 配列の長さが8のときは中間のindexは4
            center := len(sortedArray[searchStart:searchEnd])/2 + searchStart

            if sortedArray[center] == targetNumber {
                // 中間値と一致した場合
                return center
            } else if sortedArray[center] < targetNumber {
                // 中間値よりも探索対象の方が大きい場合
                // 探索対象配列の先頭indexを更新
                searchStart = center + 1
            } else if sortedArray[center] > targetNumber {
                // 中間値よりも探索対象の方が小さい場合
                // 探索対象配列の末尾indexを更新
                searchEnd = center - 1
            }
        }

    }

    // ここまで記述

    // 探索対象が存在しない場合、-1を返却
    return -1
}
