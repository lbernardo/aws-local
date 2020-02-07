ARC=$(shell go env | grep GOOS | tr "=" " " | sed s/\"/""/g |  awk '{print $2}')
build:
	 env GOOS=darwin go build -o bin/darwin/awslocal
	 env GOOS=linux  go build -o bin/linux/awslocal
install: build
	cp bin/linux/awslocal /bin/awslocal



