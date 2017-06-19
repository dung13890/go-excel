<?php

error_reporting(E_ALL);
ini_set('memory_limit', '-1');

require_once __DIR__ . '/vendor/autoload.php';
use PHPExcel_IOFactory;

$start = microtime(true);
$total = 0;

$excel = PHPExcel_IOFactory::load('test2.xlsx');

$sheetData = $excel->getActiveSheet()->toArray();
foreach ($sheetData as $value) {
    $total += $value[1];
    echo $value[1] . "\n";
}

echo "Total:" . $total . " Time:" . (microtime(true) - $start) . "\n";
