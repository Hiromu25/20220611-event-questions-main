<?php
// 処理実行
main();

function main() {
    // 昇順にソートされた配列
    $sortedArray = array(1, 2, 3, 5, 12, 7890, 12345);
    // 探索対象の番号
    $targetNumber = 7890;
    // 探索実行
    $targetIndex = serchIndex($sortedArray, $targetNumber);
    // 結果出力
    print($targetIndex);
}

function serchIndex($sortedArray, $targetNumber) {

    // ここから記述


    // ここまで記述

    // 探索対象が存在しない場合、-1を返却
    return -1;
}
?>
