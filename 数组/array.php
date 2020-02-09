<?php

$str = "0123456789";
$arr = str_split($str);
print_r($arr);
$first = 0;
$last = count($arr) - 1;
while (true) {
    if ($first < $last) {
        $first_tmp = $arr[$first];
        $arr[$first] = $arr[$last];
        $arr[$last] = $first_tmp;
        $first++;
        $last--;
    } else {
        break;
    }
}
print_r($arr);
