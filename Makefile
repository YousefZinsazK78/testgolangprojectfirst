BINARYNAME = main.out

lint:
	golangci-lint run

build: 
	go build -o ${BINARYNAME} main.go

run:
	go build -o ${BINARYNAME} main.go
	./${BINARYNAME}



clean:
	go clean
	del ${BINARYNAME}