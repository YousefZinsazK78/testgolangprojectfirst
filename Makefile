BINARYNAME = main.out

build: 
	go build -o ${BINARYNAME} main.go

run:
	go build -o ${BINARYNAME} main.go
	./${BINARYNAME}

clean:
	go clean
	rm ${BINARYNAME}