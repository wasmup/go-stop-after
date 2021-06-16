all:

test:
# go test main_test.go -v
	go test -benchmem  -bench . notmain
	go test -bench . notmain
	go test -bench .