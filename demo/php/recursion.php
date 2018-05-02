<?php
function test($arr, $search)
{
    $start = 0;
    $end   = count($arr) - 1;
    // 遍历次数
    $count = 0;

    // 开始位置不能大于结束位置
    while ($start <= $end) {
        $count++;
        $mid = floor(($start + $end) / 2);
        if ($search > $arr[$mid]) {
            // 搜索值大于中间值，开始位置为中间值+1
            $start = $mid + 1;
        } elseif ($search < $arr[$mid]) {
            // 搜索值小于中间值，结束位置为中间值-1
            $end = $mid - 1;
        } else {
            // 搜索值等于中间值，返回
            return [$mid, $count];
        }
    }
    // 没找到返回没找到
    return $count . "未找到\r\n";
}
$arr = [1, 3, 5, 7, 9];
print_r(test($arr, 1));
print_r(test($arr, 2));
print_r(test($arr, 9));
