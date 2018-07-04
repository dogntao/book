### 出现频率最高的前k个元素
> 输入
```
    function solution($line)
    {
        $arrs       = explode(" ", $line);
        $nums       = explode(",", $arrs[0]);
        $cou        = $arrs[1];
        $num_counts = [];
        // 统计出现次数
        foreach ($nums as $key => $value) {
            if (empty($num_counts[$value])) {
                $num_counts[$value] = 1;
            } else {
                $num_counts[$value]++;
            }
        }
        // print_r($num_counts);die();
        // 将对应个数的key放到对应捅中
        $bucket_key_max = 0;
        $bucket_arrs    = [];
        foreach ($num_counts as $key => $value) {
            $bucket_arrs[$value][] = $key;
            if ($value > $bucket_key_max) {
                $bucket_key_max = $value;
            }
        }
        // print_r($bucket_arrs);die();
        // 取出现最多的前k个元素
        $ret = [];
        $j   = 1;
        for ($i = $bucket_key_max; $i >= 0; $i--) {
            if (!empty($bucket_arrs[$i])) {
                foreach ($bucket_arrs[$i] as $key => $value) {
                    if ($j <= $cou) {
                        $ret[] = $value;
                        $j++;
                    }
                }
            }
        }
        sort($ret);
        return implode(",", $ret);
    }
    echo solution("1,2,2,3,3 3");
```