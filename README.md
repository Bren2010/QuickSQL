QuickSQL
========

Overview
--------

QuickSQL is a MySQL wrapper coded in [Go](http://golang.org/ "GoLang Website") that 
uses [Philio's GoMySQL](https://github.com/Philio/GoMySQL "GitHub for GoMySQL") library.
QuickSQL also has an automatic caching mechanism that is equal to or less than the speed 
of [Memcached](http://memcached.org/ "Memcached website") (I hope to have benchmarks soon).


Compiling
---------

*(Go compilers must be installed!)*

1. Run the command:  `goinstall github.com/Philio/GoMySQL`  (See installation guide [here](https://github.com/Philio/GoMySQL "") for alternatives.)
2. Run the following command in the base directory of the repository: `./compile` (This is only if you're on a 64 bit system.  On 32 bit systems, you have to compile manually using 8g.)


