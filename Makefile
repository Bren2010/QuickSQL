include $(GOROOT)/src/Make.inc

TARG=quicksql
GOFILES=quicksql.go \
		cache.go 	\
		convenience.go \
		sql.go

include $(GOROOT)/src/Make.pkg 
