<?php

if (!isset($_GET['file'])) {
    die("require file param");
}

$file = $_GET['file'];

echo "file: ". $file ."\n";

$data = file_get_contents($file);
echo $data;
