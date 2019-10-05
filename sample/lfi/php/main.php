<?php


$url = "{{ .VARIABLE }}";
echo "URL: ". $url ."\n";
// check if argument is a valid URL
if(filter_var($url, FILTER_VALIDATE_URL)) {
   // parse URL
   $r = parse_url($url);
   print_r($r);
   // check if host ends with google.com
   if(preg_match('/google\.com$/', $r['host'])) {
      // get page from URL
      $file = file_get_contents($url);
      echo($file);
   } else {
      echo "Error: Host not allowed";
   }
} else {
   echo "Error: Invalid URL";
}
