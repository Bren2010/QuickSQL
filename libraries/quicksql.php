<?php
class QuickSQL {
	
	var $conn;
	var $packetDelim;
	
	public function __construct() {
		$this->packetDelim = chr(4);
		$this->unitSep = chr(31);
	}
	
	public function exec($sql, $bypassCache = 0) {
		$this->write($sql, $bypassCache);
	}
	
	public function fetch($sql, $bypassCache = 0) {
		$data = $this->write($sql, $bypassCache);
		
		return $data['Results'];
	}
	
	public function count($sql, $bypassCache = 0) {
		$data = $this->write($sql, $bypassCache);
		
		return $data['Count'];
	}
	
	public function queryAll($sql, $bypassCache = 0) {
		$data = $this->write($sql, $bypassCache);
		
		return $data;
	}
	
	private function connect() {
		$this->conn = socket_create(AF_UNIX, SOCK_STREAM, 0);
		$connected = socket_connect($this->conn, "/tmp/dbsock", -1);
		
		if ($connected == false) die("Could not connect.");
	}
	
	private function write($array, $bypassCache) {
		if (empty($conn)) $this->connect();
		
		$toWrite = $array . $this->unitSep . (string) $bypassCache . $this->packetDelim;
		socket_write($this->conn, $toWrite);
		
		$read = true;
		$data = "";
		
		while ($read) {
			$return = @socket_recv($this->conn, $buffer, 1024, MSG_DONTWAIT);
			$data .= $buffer;

			if ($buffer[strlen($buffer) - 1] == $this->packetDelim) break;
		}
		
		return json_decode(substr($data, 0, -1), true);
	}
}
