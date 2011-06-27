QuickSQL
========

Overview
--------

QuickSQL is a small MySQL wrapper coded in [Go](http://golang.org/ "GoLang Website") that 
was designed to do all the work associated with using [Memcached](http://memcached.org/ "Memcached website") 
at more than x3 the speed.  Graciously using Philio's wonderful 
[GoMySQL](https://github.com/Philio/GoMySQL "GitHub for GoMySQL") library.


Compiling
---------

*(Go compilers must be installed!)*

1. Run the command:  `goinstall github.com/Philio/GoMySQL`  (See installation guide 
[here](https://github.com/Philio/GoMySQL "GoMySQL") 
for alternatives.)
2. Run the following command in the base directory of the repository: `./compile` 
(This is only if you're on a 64 bit system.  On 32 bit systems, you have to compile 
manually using 8g.)


Changelog
---------


**Version 1.1**

1. Removed encoding in place of EOT.
2. Removed the signal handler.
3. Errors, number of affected rows, and last insert id are now returned with the row count and query results.
4. Cache lifespan and the frequency of cache updates is now configureable in quicksql.go.
5. Clients have the option to bypass the cache.


**Version 1.2**

1. Multiple connections at once.
