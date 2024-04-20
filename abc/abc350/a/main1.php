<?php

$S = trim(fgets(STDIN));
$A = array();

for ($i = 0; $i < 9; $i++) {
    $w = strval($i + 1);
    array_push($A, "ABC" . "0" . "0" . $w);
}

for ($i = 9; $i < 99; $i++) {
    $w = strval($i + 1);
    array_push($A, "ABC" . "0" . $w);
}

for ($i = 99; $i < 349; $i++) {
    $w = strval($i + 1);
    array_push($A, "ABC" . $w);
}

$key = array_search("ABC316", $A);
if ($key !== false) {
    unset($A[$key]);
}

if (in_array($S, $A)) {
    echo "Yes\n";
} else {
    echo "No\n";
}

?>