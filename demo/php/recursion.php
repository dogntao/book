<?php

function test1(int $n)
{
    if ($n >0) {
        echo 'a'.$n . "\r\n";
        test1($n - 1);
    }
    echo 'c'.$n . "\r\n";
    return $n;
}
echo test1(2)."\r\n";
echo "......\r\n";

function test2(int $n)
{
    if ($n > 0) {
        echo 'a' . $n . "\r\n";
        $n = test2($n - 1);
    } else {
        echo 'b' . $n . "\r\n";
        return $n;
    }
    echo 'c' . $n . "\r\n";
    return $n;
}
echo test2(2)."\r\n";
echo "......\r\n";

function test3(int $n)
{
    if ($n > 0) {
        echo 'a' . $n . "\r\n";
        test3($n - 1);
    } else {
        echo 'b' . $n . "\r\n";
        return $n;
    }
    echo 'c' . $n . "\r\n";
    return $n;
}
echo test3(2)."\r\n";
echo "......\r\n";

// 查找出现位置
function FunctionName()
{
    $arr = [1,3,5,7,9];
    $search = 7;
    // 取中间数
    $len = floor(count($arr)/2);
    $mid = 0;
    
    
}
