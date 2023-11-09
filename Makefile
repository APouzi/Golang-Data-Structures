bldrun:
	go build -o algo .; ./algo -command "hello"

lldup:
	go build -o algo .; ./algo -command "duplication"

llrev:
	go build -o algo .; ./algo -command "reverse-recursive"

pattern:
	go build -o algo .; ./algo -command "pattern"

sort:
	go build -o algo .; ./algo -command "sort"

dp:
	go build -o algo .; ./algo -command "dynamic-programming"

stack:
	go build -o algo .; ./algo -command "stack"

graph:
	go build -o algo .; ./algo -command "graph"

LL:
	go build -o algo .; ./algo -command "linked-list"
	
random:
	go build -o algo .; ./algo -command "random-leet-code"

heaps:
	go build -o algo .; ./algo -command "heaps"

matrix:
	go build -o algo .; ./algo -command "matrix"