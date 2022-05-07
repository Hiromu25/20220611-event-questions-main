// 処理実行
main()

func main() {
    // 昇順にソートされた配列
    let sortedArray = [1, 2, 3, 5, 12, 7890, 12345]
    // 探索対象の番号
    let targetNumber = 7890
    // 探索実行
    let targetIndex = serchIndex(sortedArray: sortedArray, targetNumber: targetNumber)
    // 結果出力
    print(targetIndex)
}

func serchIndex(sortedArray: [Int], targetNumber: Int) -> Int {

    // ここから記述


    // ここまで記述

    // 探索対象が存在しない場合は -1を返す
    return -1
}
