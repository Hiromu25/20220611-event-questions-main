// 処理実行
main()

func main() {
    // ランダムに並べられた重複のない整数の配列
    let array = [5, 4, 6, 2, 1, 9, 8, 3, 7, 10]
    // ソート実行
    let sortedArray = sort(array: array)
    // 結果出力
    sortedArray.forEach{
        print($0)
    }
}

func sort(array: [Int]) -> [Int] {
    // 要素が一つの場合はソートの必要がないので、そのまま返却
    if(array.count == 1) {
        return array
    }

    // 配列の先頭を基準値とする
    let pivot = array[0]

    // ここから記述


    // ここまで記述

}
