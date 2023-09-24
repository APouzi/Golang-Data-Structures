bldrun:
	go build -o algo .; ./algo -command "hello"

lldup:
	go build -o algo .; ./algo -command "duplication"

llrev:
	go build -o algo .; ./algo -command "reverse-recursive"
