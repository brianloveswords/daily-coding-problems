time : build/problem70
	time build/problem70

build/problem70 : build solution.go
	go build && mv $(notdir ${@}) build

test :
	go test

build :
	mkdir build


.PHONY : time test
