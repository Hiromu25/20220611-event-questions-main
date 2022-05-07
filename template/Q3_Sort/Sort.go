package main

import (
	"fmt"
)
func main(){
    // ランダムに並べられた重複のない整数の配列
    array := []int{5, 4, 6, 2, 1, 9, 8, 3, 7, 10}
    // ソート実行
    sortedArray := sort(array)
    // 結果出力
    for _, v := range sortedArray {
        fmt.Println(v)
    }
}

func sort(array []int) []int {
    // 要素が一つの場合はソートの必要がないので、そのまま返却
    if len(array) == 1 {
        return array
    }

    // 配列の先頭を基準値とする
    pivot := array[0]

    // ここから記述

    // 検索するindexを定義
    headIndex := 0
    tailIndex := len(array) - 1

    for {
        if headIndex >= tailIndex {
            // 先頭と末端の検索がぶつかった場合
            var firstHarf, latterHarf []int
            if len(array) == 2 {
                // 配列の長さが2の場合は，ソートする必要がない
                return array
            } else if headIndex == 0 {
                // 配列内がpivot以上の値のみの場合
                // 先頭とその他に分ける
                // この処理を入れないと[5,7,6]のような配列がソートできない
                firstHarf = array[0:headIndex+1]
                latterHarf = array[headIndex+1:]
            } else if array[headIndex] >= pivot {
                // 検索がぶつかったところ or headがtailを追い越したところの値がpivot以上のとき
                firstHarf = array[0:headIndex]
                latterHarf = array[headIndex:]
            } else if array[headIndex] < pivot {
                 // 検索がぶつかったところ or headがtailを追い越したところの値がpivot未満のとき
                firstHarf = array[0:headIndex+1]
                latterHarf = array[headIndex+1:]
            }

            // sort関数を再帰的に呼び出す
            return append(sort(firstHarf),sort(latterHarf)...)

        } else if array[headIndex] >= pivot && array[tailIndex] < pivot {
            // pivot以上の値とpivot未満の値が見つかった場合

            // 配列の要素を入れ替え
            tmpValue := array[headIndex]
            array[headIndex] = array[tailIndex]
            array[tailIndex] = tmpValue
            
            // index更新
            headIndex++
            tailIndex--
        } else if array[headIndex] >= pivot && array[tailIndex] >= pivot {
            // pivot以上の値は見つかっているが，pivot未満の値が見つかっていない場合
            // 末尾の検索を進める
            tailIndex--
        } else if array[headIndex] < pivot && array[tailIndex] < pivot {
            // pivot以上の値は見つかっていないが，pivot未満の値が見つかっている場合
            // 先頭の検索を進める
            headIndex++
        } else if array[headIndex] < pivot && array[tailIndex] >= pivot {
            // pivot以上の値も，pivot未満の値も見つかっていない場合
            // 先頭と末尾の検索を進める
            headIndex++
            tailIndex--
        }
    }

    // ここまで記述

}
