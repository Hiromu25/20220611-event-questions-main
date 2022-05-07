using System.Linq;

public class Sort{
    public static void Main(){
        // ランダムに並べられた重複のない整数の配列
        var array = new int[] { 5, 4, 6, 2, 1, 9, 8, 3, 7, 10 };
        // ソート実行
        var sortedArray = new Sort().SortArray(array);
        // 結果出力
        foreach (int i in sortedArray)
        {
            System.Console.WriteLine(i);
        }
    }

    public int[] SortArray(int[] array){
        // 要素が一つの場合はソートの必要がないので、そのまま返却
        if(array.Length == 1) {
            return array;
        }

        // 配列の先頭を基準値とする
        var pivot = array[0];

        // ここから記述


        // ここまで記述

    }
}
