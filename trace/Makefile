.PHONY:

trace:
	go run main.go 2> trace.out

test:
	go test -trace=test.out

view:
	go tool trace trace.out