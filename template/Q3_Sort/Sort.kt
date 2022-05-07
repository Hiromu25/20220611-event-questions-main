fun main() {
    // ランダムに並べられた重複のない整数の配列
    val array = arrayOf(5, 4, 6, 2, 1, 9, 8, 3, 7, 10)
    // ソート実行
    val sortedArray = sort(array)
    // 結果出力
    sortedArray.forEach{
        println(it)
    }
}

fun sort(array: Array<Int>): Array<Int> {
    // 要素が一つの場合はソートの必要がないので、そのまま返却
    if(array.size == 1) {
        return array
    }

    // 配列の先頭を基準値とする
    val pivot = array[0]

    // ここから記述


    // ここまで記述

}
