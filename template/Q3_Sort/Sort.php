<?php
// 処理実行
main();

function main() {
    // ランダムに並べられた重複のない整数の配列
    $array = array(5, 4, 6, 2, 1, 9, 8, 3, 7, 10);
    // ソート実行
    $sortedArray = sortArray($array);
    // 結果出力
    foreach ($sortedArray as $i) {
        print_r($i."\n");
    }
}

function sortArray(array $array): array
{
    // 要素が一つの場合はソートの必要がないので、そのまま返却
    if (count($array) == 1) {
        return $array;
    }

    // 配列の先頭を基準値とする
    $pivot = $array[0];

    // ここから記述


    // ここまで記述

}
?>
