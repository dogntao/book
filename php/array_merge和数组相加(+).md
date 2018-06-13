## array_merge和数组相加(+)的区别
> 1、当键名为字符串时候
```
    $arr1 = ['a' => 'php'];
    $arr2 = ['a' => 'go', 'b' => 'java'];

    array_merge 后者替换前者
    Array ( [a] => go [b] => java )
    相加(+)      后者不替换前者
    Array ( [a] => php [b] => java )
    
```
> 2、当键名为数字时候
```
    $arr1 = [0 => 'php'];
    $arr2 = [0 => 'go', 3 => 'java'];

    array_merge 合并所有值,key从0开始
    Array ( [0] => php [1] => go [2] => java )
    相加(+)      保留首先出现的
    Array ( [0] => php [3] => java )

```