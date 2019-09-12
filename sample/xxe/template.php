<?php
$filename = "{{ .FILENAME }}";

$doc = new DOMDocument();
$doc->substituteEntities = true;

$doc->load($filename);

$contents = $doc->getElementsByTagName('data')->item(0)->textContent;

echo $contents;
