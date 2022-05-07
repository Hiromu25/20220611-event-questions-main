fun main() {
    // 昇順にソートされた配列
    val sortedArray = arrayOf(1, 2, 3, 5, 12, 7890, 12345)
    // 探索対象の番号
    val targetNumber = 7890
    // 探索実行
    val targetIndex = serchIndex(sortedArray, targetNumber)
    // 結果出力
    println(targetIndex)
}

fun serchIndex(sortedArray: Array<Int>, targetNumber: Int): Int {

    // ここから記述


    // ここまで記述

    // 探索対象がない場合は -1 を返す
    return -1
}
