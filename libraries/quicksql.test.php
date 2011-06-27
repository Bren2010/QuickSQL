<?php
include("./quicksql.php");

$start = microtime(true);

$db = new QuickSQL;
$response = $db->queryAll("SHOW TABLES;");

$end = microtime(true);

echo "Response:\n";
print_r($response);
echo "Took ", $end - $start, " seconds to query.\n";
?>
