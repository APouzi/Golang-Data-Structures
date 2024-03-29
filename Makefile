#Name needs to be diff from folder name
bldrun:
	go build -o algo .; ./algo -command "hello"

lldup:
	go build -o algo .; ./algo -command "duplication"

llrev:
	go build -o algo .; ./algo -command "reverse-recursive"

pattern:
	go build -o algo .; ./algo -command "pattern"

binary-search:
	go build -o algo .; ./algo -command "binary-search"

fast-slow:
	go build -o algo .; ./algo -command "fast-slow"

merge:
	go build -o algo .; ./algo -command "merge-interval"

two-pointer:
	go build -o algo .; ./algo -command "two-pointer"

sliding-window:
	go build -o algo .; ./algo -command "sliding-window"

sort:
	go build -o algo .; ./algo -command "sort"

dp:
	go build -o algo .; ./algo -command "dynamic-programming"

stacks:
	go build -o algo .; ./algo -command "stack"

graph:
	go build -o algo .; ./algo -command "graph"

ll:
	go build -o algo .; ./algo -command "linked-list"
	
random:
	go build -o algo .; ./algo -command "random-leet-code"

heaps:
	go build -o algo .; ./algo -command "heaps"

matrix:
	go build -o algo .; ./algo -command "matrix"

array:
	go build -o algo .; ./algo -command "array"

tree:
	go build -o algo .; ./algo -command "tree"

problem:
	go build -o algo .; ./algo -command "problem"

problems:
	go build -o algo .; ./algo -command "problems"

mono:
	go build -o algo .; ./algo -command "mono"
master:
	go build -o algo .; ./algo -command "master"
