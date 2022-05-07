package main
import "fmt"
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


    // ここまで記述

    // 探索対象が存在しない場合、-1を返却
    return -1
}
