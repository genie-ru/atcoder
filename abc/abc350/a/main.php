<?php

$input = trim(fgets(STDIN));

$num = substr($input, 3, 3);

// print_r($num);

if($num == 316 || $num >= 350 || $num == 000){
    echo "No\n";
    exit;
}else{
    echo "Yes\n";
    exit;
}